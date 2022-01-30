package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/claris/light/generator"
	"github.com/claris/light/generator/template"

	mstrings "github.com/claris/light/generator/strings"

	"os"
	"strings"

	lg "github.com/claris/light/logger"
	"github.com/vetcher/go-astra"
	"github.com/vetcher/go-astra/types"
)

var (
	flagFileName     = flag.String("file", "service.go", "Path to input file with interface.")
)


func main() {
	flag.Parse()

	fmt.Println(*flagFileName)
	info, err := astra.ParseFile( *flagFileName)
	if err != nil {
		lg.Logger.Logln(0, "fatal:", err)
		os.Exit(1)
	}
	fmt.Println(info)

	i := findInterface(info)
	if i == nil {
		lg.Logger.Logln(0, "fatal: could not find interface with @microgen tag")
		lg.Logger.Logln(4, "All founded interfaces:")
		lg.Logger.Logln(4, listInterfaces(info.Interfaces))
		os.Exit(1)
	}
	fmt.Println(i)
	ctx, err := prepareContext(*flagFileName, i)
	if err != nil {
		lg.Logger.Logln(0, "fatal:", err)
		os.Exit(1)
	}
	fmt.Println(ctx)
}

func prepareContext(filename string, iface *types.Interface) (context.Context, error) {
	ctx := context.Background()
	p, err := astra.ResolvePackagePath(filename)
	if err != nil {
		return nil, err
	}
	fmt.Println(p)
	ctx = template.WithSourcePackageImport(ctx, p)

	set := template.TagsSet{}
	genTags := mstrings.FetchTags(iface.Docs, generator.TagMark+generator.LightMainTag)
	for _, tag := range genTags {
		set.Add(tag)
	}
	ctx = template.WithTags(ctx, set)

	return ctx, nil
}

func listInterfaces(ii []types.Interface) string {
	var s string
	for _, i := range ii {
		s = s + fmt.Sprintf("\t%s(%d methods, %d embedded interfaces)\n", i.Name, len(i.Methods), len(i.Interfaces))
	}
	return s
}

func findInterface(file *types.File) *types.Interface {
	for i := range file.Interfaces {
		fmt.Printf("doc is %v\n", file.Interfaces[i].Docs)
		if docsContainLightTag(file.Interfaces[i].Docs) {
			return &file.Interfaces[i]
		}
	}
	return nil
}

func docsContainLightTag(strs []string) bool {
	for _, str := range strs {
		fmt.Printf("s = %v\n", str)
		if strings.HasPrefix(str, generator.TagMark+generator.LightMainTag) {
			return true
		}
	}
	return false
}

