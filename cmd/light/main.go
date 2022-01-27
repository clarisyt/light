package main

import (
	"flag"
	"fmt"
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
		if docsContainMicrogenTag(file.Interfaces[i].Docs) {
			return &file.Interfaces[i]
		}
	}
	return nil
}

func docsContainMicrogenTag(strs []string) bool {
	for _, str := range strs {
		if strings.HasPrefix(str, generator.TagMark+generator.LightMainTag) {
			return true
		}
	}
	return false
}

