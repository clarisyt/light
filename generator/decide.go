package generator

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	mstrings "github.com/claris/light/generator/strings"
	"github.com/claris/light/generator/template"
	lg "github.com/claris/light/logger"
	"github.com/vetcher/go-astra/types"
)

const (
	TagMark         = template.TagMark
	LightMainTag = template.LightMainTag
	ProtobufTag     = "protobuf"
	GRPCClientAddr  = "grpc-addr"

	MainTag                   = template.MainTag
)

func ListTemplatesForGen(ctx context.Context, iface *types.Interface, absOutPath, sourcePath string, genProto string, genMain bool) (units []*GenerationUnit, err error) {
	importPackagePath, err := resolvePackagePath(filepath.Dir(sourcePath))
	if err != nil {
		return nil, err
	}
	absSourcePath, err := filepath.Abs(sourcePath)
	if err != nil {
		return nil, err
	}
	outImportPath, err := resolvePackagePath(absOutPath)
	if err != nil {
		return nil, err
	}
	m := make(map[string]bool, len(iface.Methods))
	for _, fn := range iface.Methods {
		m[fn.Name] = !mstrings.ContainTag(mstrings.FetchTags(fn.Docs, TagMark+LightMainTag), "-")
	}
	info := &template.GenerationInfo{
		SourcePackageImport:   importPackagePath,
		SourceFilePath:        absSourcePath,
		Iface:                 iface,
		OutputPackageImport:   outImportPath,
		OutputFilePath:        absOutPath,
		ProtobufPackageImport: mstrings.FetchMetaInfo(TagMark+ProtobufTag, iface.Docs),
		FileHeader:            defaultFileHeader,
		AllowedMethods:        m,
		ProtobufClientAddr:    mstrings.FetchMetaInfo(TagMark+GRPCClientAddr, iface.Docs),
	}
	lg.Logger.Logln(3, "\nGeneration Info:", info.String())
	/*stubSvc, err := NewGenUnit(ctx, template.NewStubInterfaceTemplate(info), absOutPath)
	if err != nil {
		return nil, err
	}
	units = append(units, stubSvc)*/

	genTags := mstrings.FetchTags(iface.Docs, TagMark+LightMainTag)
	lg.Logger.Logln(2, "Tags:", strings.Join(genTags, ", "))
	uniqueTemplate := make(map[string]template.Template)
	for _, tag := range genTags {
		templates := tagToTemplate(tag, info)
		if templates == nil {
			lg.Logger.Logln(1, "Warning: Unexpected tag", tag)
			continue
		}
		for _, t := range templates {
			uniqueTemplate[t.DefaultPath()] = t
		}
	}
	for _, t := range uniqueTemplate {
		unit, err := NewGenUnit(ctx, t, absOutPath)
		if err != nil {
			return nil, fmt.Errorf("%s: %v", absOutPath, err)
		}
		units = append(units, unit)
	}
	if genProto != "" {
		u, err := NewGenUnit(ctx, template.NewProtoTemplate(info, genProto), absOutPath)
		if err != nil {
			return nil, fmt.Errorf("%s: %v", absOutPath, err)
		}
		units = append(units, u)
	}
	if genMain {
		u, err := NewGenUnit(ctx, template.NewMainTemplate(info), absOutPath)
		if err != nil {
			return nil, fmt.Errorf("%s: %v", absOutPath, err)
		}
		units = append(units, u)
	}
	return units, nil
}

func tagToTemplate(tag string, info *template.GenerationInfo) (tmpls []template.Template) {
	switch tag {
	case MiddlewareTag:
		return append(tmpls, template.NewMiddlewareTemplate(info))
	case LoggingMiddlewareTag:
		return append(
			append(tmpls, tagToTemplate(MiddlewareTag, info)...),
			template.NewLoggingTemplate(info),
		)
	case GrpcTag:
		return append(
			append(tmpls, tagToTemplate(Transport, info)...),
			template.NewGRPCClientTemplate(info),
			template.NewGRPCServerTemplate(info),
			template.NewGRPCEndpointConverterTemplate(info),
			template.NewStubGRPCTypeConverterTemplate(info),
		)
	case GrpcClientTag:
		return append(
			append(tmpls, tagToTemplate(TransportClient, info)...),
			template.NewGRPCClientTemplate(info),
			template.NewGRPCEndpointConverterTemplate(info),
			template.NewStubGRPCTypeConverterTemplate(info),
		)
	case GrpcServerTag:
		return append(
			append(tmpls, tagToTemplate(TransportServer, info)...),
			template.NewGRPCServerTemplate(info),
			template.NewGRPCEndpointConverterTemplate(info),
			template.NewStubGRPCTypeConverterTemplate(info),
		)
	case HttpTag:
		return append(
			append(tmpls, tagToTemplate(Transport, info)...),
			template.NewHttpServerTemplate(info),
			template.NewHttpClientTemplate(info),
			template.NewHttpConverterTemplate(info),
		)
	case HttpServerTag:
		return append(
			append(tmpls, tagToTemplate(TransportServer, info)...),
			template.NewHttpServerTemplate(info),
			template.NewHttpConverterTemplate(info),
		)
	case HttpClientTag:
		return append(
			append(tmpls, tagToTemplate(TransportClient, info)...),
			template.NewHttpClientTemplate(info),
			template.NewHttpConverterTemplate(info),
		)
	case RecoveringMiddlewareTag:
		return append(
			append(tmpls, tagToTemplate(MiddlewareTag, info)...),
			template.NewRecoverTemplate(info),
		)
	case MainTag:
		lg.Logger.Logln(1, "Warning: Tag main is deprecated, use flag -main instead.")
		return nil
	case ErrorLoggingMiddlewareTag:
		return append(
			append(tmpls, tagToTemplate(MiddlewareTag, info)...),
			template.NewErrorLoggingTemplate(info),
		)
	case CachingMiddlewareTag:
		return append(
			append(tmpls, tagToTemplate(MiddlewareTag, info)...),
			template.NewCacheMiddlewareTemplate(info),
		)
	case TracingMiddlewareTag:
		return append(tmpls, template.EmptyTemplate{})
	case MetricsMiddlewareTag:
		return append(tmpls, template.EmptyTemplate{})
	case ServiceDiscoveryTag:
		return append(tmpls, template.EmptyTemplate{})
	case Transport:
		return append(tmpls,
			template.NewExchangeTemplate(info),
			template.NewEndpointsTemplate(info),
			template.NewEndpointsClientTemplate(info),
			template.NewEndpointsServerTemplate(info),
		)
	case TransportClient:
		return append(tmpls,
			template.NewExchangeTemplate(info),
			template.NewEndpointsTemplate(info),
			template.NewEndpointsClientTemplate(info),
		)
	case TransportServer:
		return append(tmpls,
			template.NewExchangeTemplate(info),
			template.NewEndpointsTemplate(info),
			template.NewEndpointsServerTemplate(info),
		)
	}
	return nil
}

func resolvePackagePath(outPath string) (string, error) {
	lg.Logger.Logln(3, "Try to resolve path for", outPath, "package...")
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		return "", fmt.Errorf("GOPATH is empty")
	}
	lg.Logger.Logln(4, "GOPATH:", gopath)

	absOutPath, err := filepath.Abs(outPath)
	if err != nil {
		return "", err
	}
	lg.Logger.Logln(4, "Resolving path:", absOutPath)

	for _, path := range strings.Split(gopath, ":") {
		gopathSrc := filepath.Join(path, "src")
		if strings.HasPrefix(absOutPath, gopathSrc) {
			return absOutPath[len(gopathSrc)+1:], nil
		}
	}
	return "", fmt.Errorf("path(%s) not in GOPATH(%s)", absOutPath, gopath)
}