package ic

import (
	//"fmt"
	"os"
	"testing"

	"github.com/sdkvictor/golang-compiler/ast"
	"github.com/sdkvictor/golang-compiler/gocc/lexer"
	"github.com/sdkvictor/golang-compiler/gocc/parser"
	"github.com/sdkvictor/golang-compiler/semantics"

	"github.com/davecgh/go-spew/spew"
)

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

func TestGenerateIntermediateCode(t *testing.T) {
	p := parser.NewParser()
	tests := []string{
		"test/test2.vm",
	}

	for _, test := range tests {
		input, err := readFile(test)

		if err != nil {
			t.Fatalf("Error reading file %s", test)
		}

		s := lexer.NewLexer(input)
		pro, err := p.Parse(s)
		if err != nil {
			t.Fatalf("%s: %v", test, err)
		}

		program, ok := pro.(*ast.Program)
		if !ok {
			t.Fatalf("Cannot cast to Program")
		}

		funcdir, globals, err := semantics.SemanticCheck(program)
		if err != nil {
			t.Fatalf("Error from semantic: %v", err)
		}

		gen, _, err := GenerateIntermediateCode(program, funcdir, globals)
		if err != nil {
			t.Fatalf("Error from generate code: %v", err)
		}

		spew.Dump(gen)
		//fmt.Printf("%s\n", vm)
	}
}