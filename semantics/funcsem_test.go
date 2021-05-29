package sem

import (
	"github.com/sdkvictor/golang-compiler/ast"
	"github.com/sdkvictor/golang-compiler/directories"
	"github.com/sdkvictor/golang-compiler/gocc/lexer"
	"github.com/sdkvictor/golang-compiler/gocc/parser"
	//"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestBuildFuncDirProgram(t *testing.T) {
	p := parser.NewParser()
	tests := []string{
		"test/test1.vm",
	}

	for _, test := range tests {
		input, err := readFile(test)

		if err != nil {
			t.Fatalf("Error reading file %s", test)
		}

		s := lexer.NewLexer(input)
		pro, err := p.Parse(s)
		if err != nil {
			t.Errorf("%s: %v", test, err)
		}

		program, ok := pro.(*ast.Program)
		if !ok {
			t.Fatalf("Cannot cast to Program")
		}

		funcdir := directories.NewFuncDirectory()

		err = buildFuncDirProgram(program, funcdir)
		if err != nil {
			t.Errorf("buildFuncDirProgram: %v", err)
		}

		//spew.Dump(funcdir)
	}
}