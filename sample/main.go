package main

import (
	"context"
	"fmt"
	. "github.com/dave/jennifer/jen"
)

func sampme1() {
	f := NewFile("main")
	f.Func().Id("main").Params().Block(
		Qual("fmt", "Println").Call(Lit("Hello, world")),
	)
	fmt.Printf("%#v", f)
}

func sample2() {
	ctx := context.Background()
	//ts := ctx.Value("hello").(map[string]struct{})
	context.WithValue(ctx, "hello", "beijing")
	fmt.Printf("%v\n", ctx.Value("hello"))
	/*
	for _, t  := range ts {
		fmt.Println(t)
	}
	*/
}

func sample3() {

}

func main() {
	sample2()
}
