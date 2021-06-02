package main

import (
	"os"
	"fmt"

	"github.com/sdkvictor/golang-compiler/ast"
	"github.com/sdkvictor/golang-compiler/gocc/lexer"
	"github.com/sdkvictor/golang-compiler/gocc/parser"
	"github.com/sdkvictor/golang-compiler/semantics"
	"github.com/sdkvictor/golang-compiler/ic"
	"github.com/sdkvictor/golang-compiler/vm"
	"github.com/faiface/pixel/pixelgl"
	"github.com/mewkiz/pkg/errutil"
	//"github.com/davecgh/go-spew/spew"
)

func usage() {
	fmt.Printf("Usage: run <vm source file>\n")
}

func readFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	_, err = file.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func compile(file string) (*ic.Generator, map[string]int, error) {
	p := parser.NewParser()
	input, err := readFile(file)

	if err != nil {
		return nil, nil, err
	}

	s := lexer.NewLexer(input)
	pro, err := p.Parse(s)

	if err != nil {
		return nil, nil, err
	}

	program, ok := pro.(*ast.Program)
	if !ok {
		return nil, nil, errutil.NewNoPos("Cannot cast program")
	}

	funcdir, globals, err := semantics.SemanticCheck(program)
	if err != nil {
		return nil, nil, err
	}

	gen, vm, err := ic.GenerateIntermediateCode(program, funcdir, globals)
	if err != nil {
		return nil, nil, err
	}

	return gen, vm.GetConstantMap(), nil
}

func run() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	file := os.Args[1]

	gen, consmap, err := compile(file)
	if err != nil {
		fmt.Printf("Compilation %v\n", err)
		return
	}

	fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")

	//spew.Dump(consmap)

	mach := vm.NewVirtualMachine(gen.Quadruples(), consmap)
	err = mach.LoadConstants(consmap)
	if err != nil {
		fmt.Printf("Setup %v\n", err)
		return
	}
	
	
	//fmt.Printf("%s\n", mach)

	err = mach.Run()
	if err != nil {
		fmt.Printf("Runtime %v\n", err)
		return
	}
}

func main() {
	pixelgl.Run(run)
}