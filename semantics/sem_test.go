package semantics

import (
	"github.com/sdkvictor/golang-compiler/ast"
	"github.com/sdkvictor/golang-compiler/gocc/lexer"
	"github.com/sdkvictor/golang-compiler/gocc/parser"
	//"github.com/davecgh/go-spew/spew"

	"os"
	"testing"
	//"fmt"
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

func TestSemanticCheck(t *testing.T) {
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

		//spew.Dump(program)

		_, _, err = SemanticCheck(program)
		if err != nil {
			t.Fatalf("Error from semantic: %v", err)
		}
		//spew.Dump(funcdir)
	}
}