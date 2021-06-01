package semantics

import (
	"github.com/sdkvictor/golang-compiler/ast"
	"github.com/sdkvictor/golang-compiler/directories"
	"github.com/sdkvictor/golang-compiler/types"
	"github.com/mewkiz/pkg/errutil"
)

//buildFuncDirProgram receives the program and the function directory to start building the funcdir
func buildFuncDirProgram(program *ast.Program, funcdir *directories.FuncDirectory) error {
	for _, f := range program.Functions() {
		if err := buildFuncDirFunction(f, funcdir); err != nil { // CASI OK
			return err
		}
	}



	return nil
}

//buildFuncDirFunction creates a new FuncEntry for the function and adds it to the directory
func buildFuncDirFunction(function *ast.Function, funcdir *directories.FuncDirectory) error {
	id := function.Id()
	t := function.Type()
	vardir := directories.NewVarDirectory()
	params := make([]*types.Type, 0)

	if funcdir.Exists(id) {
		return errutil.NewNoPosf("%+v: Redeclaration of function %s", function.Token(), id)
	}

	if IdIsReserved(id) {
		return errutil.NewNoPosf("%+v: Cannot declare a function with reserved keyword %s", function.Token(), id)
	}

	FixParams(function.Params())

	for _, p := range function.Params() {
		params = append(params, p.Type())
		if err := buildVarDirFunction(p, vardir); err != nil {
			return err
		}
	}

	// ?? y la variable reservedFunctions de funcutil?
	if kw, ok := checkVarDirReserved(vardir); !ok {
		return errutil.NewNoPosf("%+v: Cannot declare variable with reserved keyword %s", function.Token(), kw)
	}

	fe := directories.NewFuncEntry(id, t, params, vardir)

	if ok := funcdir.Add(fe); !ok {
		return errutil.NewNoPosf("%+v: Invalid Function. This Function already exists.", function.Token())
	}

	return nil
}

//buildVarDirFunction
func buildVarDirFunction(ve *directories.VarEntry, vardir *directories.VarDirectory) error {
	if ok := vardir.Add(ve); !ok {
		return errutil.NewNoPosf("%+v: Invalid parameter. This parameter has already been declared.", ve.Token())
	}
	return nil
}

// FixParams invierte el id de pocision en los params ya que se construyen al reves en el AST.
func FixParams(params []*directories.VarEntry) {

	positions := make([]int, 0)

	for _, p := range params {
		positions = append(positions, p.Pos())
	}

	for i, _ := range params {
		params[len(params)-1-i].SetPos(positions[i])
	}
}