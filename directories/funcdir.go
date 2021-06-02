package directories

import (
	"github.com/sdkvictor/golang-compiler/types"
	"github.com/sdkvictor/golang-compiler/mem"
)

// A FuncEntry represents a declaration of a function that will be stored in the functions directory.
// Example of declaration:
//
// returntype id(params) {}
type FuncEntry struct {
	id			string
	returntype 	*types.Type
	params 		[]*types.Type
	vardir		*VarDirectory
	loc			mem.Address
	varcounter 	int
	era			int
}

// Key returns the key of the FuncEntry
func (fe *FuncEntry) Key() string {
	return fe.id
}

// Id returns the name of the FuncEntry
func (fe *FuncEntry) Id() string {
	return fe.id;
}

// ReturnType returns the type of the return value of the FuncEntry
func (fe *FuncEntry) ReturnType() *types.Type {
	return fe.returntype;
}

// Params return the list of parameters of the FuncEntry
func (fe *FuncEntry) Params() []*types.Type {
	return fe.params;
}

// VarDir returns the variables directory of the FuncEntry
func (fe *FuncEntry) VarDir() *VarDirectory {
	return fe.vardir;
}

// Era returns the size of the FuncEntry
func (fe *FuncEntry) Era() int {
	return fe.era
}

// Loc returns the memory address of the FuncEntry
func (fe *FuncEntry) Loc() mem.Address {
	return fe.loc
}

func (fe *FuncEntry) Varcounter() int {
	return fe.varcounter
}

func (fe *FuncEntry) SetVarcounter(i int) {
	fe.varcounter = i
}

func (fe *FuncEntry) IncreaseVarcounter() {
	fe.varcounter++
}

// SetLocation sets the memory address of a FuncEntry
func (fe *FuncEntry) SetLocation(loc int) {
	fe.loc = mem.Address(loc)
}

// CreateFuncEntry creates a new FuncEntry struct
func NewFuncEntry(id string, returntype *types.Type, params []*types.Type, vardir *VarDirectory) *FuncEntry {
	return &FuncEntry{id, returntype, params, vardir, mem.Address(-1), 0, 0}
}

// A FuncDirectory is the function directory which represents a table that stores all the instances of FuncEntry
type FuncDirectory struct {
	table map[string]*FuncEntry
}

// Table returns the table which stores all the functions declared
func (fd *FuncDirectory) Table() map[string]*FuncEntry {
	return fd.table
}

// Add inserts a new FuncEntry into the function directory
func (fd *FuncDirectory) Add(e *FuncEntry) bool {
	_, ok := fd.table[e.Key()]

	if !ok {
		fd.table[e.Key()] = e
	}

	return !ok
}

// Get returns a FuncEntry in case it exists in the function directory
func (fd *FuncDirectory) Get(key string) *FuncEntry {
	if result, ok := fd.table[key]; ok {
		return result
	}

	return nil
}

// Exists returns true if the FuncEntry was already added to the function directory, false otherwise
func (fd *FuncDirectory) Exists(key string) bool {
	_, ok := fd.table[key]

	return ok
}

// NewFuncDirectory creates a new empty function directory
func NewFuncDirectory() *FuncDirectory {
	return &FuncDirectory{make(map[string]*FuncEntry)}
}

// MainFuncEntry Initialization of the function directory with the initial parameters of the main program
func MainFuncEntry() *FuncEntry {
	return &FuncEntry{"main", types.NewDataType(types.Int, 0, 0), make([]*types.Type, 0), NewVarDirectory(), mem.Address(-1), 0, 0}
}