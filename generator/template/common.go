package template

import (
	. "github.com/dave/jennifer/jen"
	"github.com/vetcher/go-astra/types"
)


const (
	TagMark         = "// @"
	LightMainTag = "light"
)

const (
	MiddlewareTag             = "middleware"
	LoggingMiddlewareTag      = "logging"
	RecoveringMiddlewareTag   = "recovering"
	HttpTag                   = "http"
	HttpServerTag             = "http-server"
	HttpClientTag             = "http-client"
	GrpcTag                   = "grpc"
	GrpcServerTag             = "grpc-server"
	GrpcClientTag             = "grpc-client"
	MainTag                   = "main"
	ErrorLoggingMiddlewareTag = "error-logging"
	TracingMiddlewareTag      = "tracing"
	CachingMiddlewareTag      = "caching"
	JSONRPCTag                = "json-rpc"
	JSONRPCServerTag          = "json-rpc-server"
	JSONRPCClientTag          = "json-rpc-client"
	Transport                 = "transport"
	TransportClient           = "transport-client"
	TransportServer           = "transport-server"
	MetricsMiddlewareTag      = "metrics"
	ServiceDiscoveryTag       = "service-discovery"
)

const (
	MicrogenExt    = ".microgen.go"
	PathService    = "service"
	PathTransport  = "transport"
	PathExecutable = "cmd"
)

type WriteStrategyState int

const (
	PackagePathGoKitEndpoint         = "github.com/go-kit/kit/endpoint"
	PackagePathContext               = "context"
	PackagePathGoKitLog              = "github.com/go-kit/kit/log"
	PackagePathTime                  = "time"
	PackagePathGoogleGRPC            = "google.golang.org/grpc"
	PackagePathGoogleGRPCStatus      = "google.golang.org/grpc/status"
	PackagePathGoogleGRPCCodes       = "google.golang.org/grpc/codes"
	PackagePathNetContext            = "golang.org/x/net/context"
	PackagePathGoKitTransportGRPC    = "github.com/go-kit/kit/transport/grpc"
	PackagePathHttp                  = "net/http"
	PackagePathGoKitTransportHTTP    = "github.com/go-kit/kit/transport/http"
	PackagePathBytes                 = "bytes"
	PackagePathJson                  = "encoding/json"
	PackagePathIOUtil                = "io/ioutil"
	PackagePathIO                    = "io"
	PackagePathStrings               = "strings"
	PackagePathUrl                   = "net/url"
	PackagePathEmptyProtobuf         = "github.com/golang/protobuf/ptypes/empty"
	PackagePathFmt                   = "fmt"
	PackagePathOs                    = "os"
	PackagePathOsSignal              = "os/signal"
	PackagePathSyscall               = "syscall"
	PackagePathErrors                = "errors"
	PackagePathNet                   = "net"
	PackagePathGorillaMux            = "github.com/gorilla/mux"
	PackagePathPath                  = "PackagePathOsSignalpath"
	PackagePathStrconv               = "strconv"
	PackagePathOpenTracingGo         = "github.com/opentracing/opentracing-go"
	PackagePathGoKitTracing          = "github.com/go-kit/kit/tracing/opentracing"
	PackagePathGoKitTransportJSONRPC = "github.com/go-kit/kit/transport/http/jsonrpc"
	PackagePathGoKitMetrics          = "github.com/go-kit/kit/metrics"
	PackagePathGoKitSD               = "github.com/go-kit/kit/sd"
	PackagePathGoKitLB               = "github.com/go-kit/kit/sd/lb"
	PackagePathSyncErrgroup          = "golang.org/x/sync/errgroup"

	TagMark         = "// @"
	MicrogenMainTag = "microgen"
	serviceAlias    = "service"
)

const (
	FileStrat WriteStrategyState = iota + 1
	AppendStrat
)

type GenerationInfo struct {
	Iface               *types.Interface
	SourcePackageImport string
	SourceFilePath      string
	OutputPackageImport string
	OutputFilePath      string
	FileHeader          string

	ProtobufPackageImport string
	ProtobufClientAddr    string
	AllowedMethods        map[string]bool
}

var ctx_contextContext = Id(_ctx_).Qual(PackagePathContext, "Context")



