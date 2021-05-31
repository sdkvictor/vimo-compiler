//Package ic provides the generation of intermediate code
package ic

import (
	"github.com/sdkvictor/golang-compiler/ast"
	"github.com/sdkvictor/golang-compiler/directories"
	"github.com/sdkvictor/golang-compiler/mem"
	"github.com/sdkvictor/golang-compiler/semantics"
)

//GenerationContext djsknfkjsdfkj
type GenerationContext struct {
	funcdir *directories.FuncDirectory
	globals *directories.VarDirectory
	semcube *semantics.SemanticCube
	gen     *Generator
	vm      *mem.VirtualMemory
}

// FuncDir ...
func (ctx *GenerationContext) FuncDir() *directories.FuncDirectory {
	return ctx.funcdir
}

// FuncDir ...
func (ctx *GenerationContext) Globals() *directories.VarDirectory {
	return ctx.globals
}

// SemCube ...
func (ctx *GenerationContext) SemCube() *semantics.SemanticCube {
	return ctx.semcube
}

// Generator ...
func (ctx *GenerationContext) Generator() *Generator {
	return ctx.gen
}

//VM ...
func (ctx *GenerationContext) VM() *mem.VirtualMemory {
	return ctx.vm
}

//GenerateIntermediateCode calls the two main code generation functions, first the function to generate addresses and then the
// function to generate the code itself
func GenerateIntermediateCode(program *ast.Program, funcdir *directories.FuncDirectory, globals *directories.VarDirectory) (*Generator, *mem.VirtualMemory, error) {
	ctx := &GenerationContext{funcdir, globals, semantics.NewSemanticCube(), NewGenerator(), mem.NewVirtualMemory()}

	// GenerateAddresses intilializes all entries in every VarDirectory with an address
	// assigned by the VirtualMemory manager
	// This function must be called before the generateCode function to ensure every
	// variable and constant has a valid address
	if err := generateAddressesProgram(program, ctx); err != nil {
		return nil, nil, err
	}

	// GenerateCodeProgram
	if err := generateCodeProgram(program, ctx); err != nil {
		return nil, nil, err
	}

	return ctx.gen, ctx.vm, nil
}