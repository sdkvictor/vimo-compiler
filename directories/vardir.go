package directories

import (
	//"fmt"
	"github.com/sdkvictor/golang-compiler/gocc/token"
	"github.com/sdkvictor/golang-compiler/mem"
	"github.com/sdkvictor/golang-compiler/types"
	"github.com/mewkiz/pkg/errutil"
)

type VarEntry struct {
	id   string
	t    *types.Type
	tok  *token.Token
	addr mem.Address
	pos  int
}

func (ve *VarEntry) Id() string {
	return ve.id
}

func (ve *VarEntry) Type() *types.Type {
	return ve.t
}

func (ve *VarEntry) Token() *token.Token {
	return ve.tok
}

func (ve *VarEntry) Address() mem.Address {
	return ve.addr
}

func (ve *VarEntry) SetAddress(addr mem.Address) {
	ve.addr = addr
}

func (ve *VarEntry) Pos() int {
	return ve.pos
}

func (ve *VarEntry) SetPos(i int) {
	ve.pos = i
}

type VarDirectory struct {
	table map[string]*VarEntry
}
/*
func (e *VarEntry) String() string {
	return fmt.Sprintf("%s", e.id)
}
*/
//NewVarEntry Initialization of one entry of the variable with its attributes
func NewVarEntry(id string, t *types.Type, tok *token.Token, pos int) *VarEntry {
	return &VarEntry{id, t, tok, 0, pos}
}

//Add Add a varentry to the directory variables using the toString function as key
func (vd *VarDirectory) Add(e *VarEntry) bool {

	_, ok := vd.table[e.Id()]
	if !ok {
		vd.table[e.Id()] = e
	}
	return !ok
}

func (vd *VarDirectory) Get(key string) *VarEntry {

	if result, ok := vd.table[key]; ok {
		return result
	}

	return nil
}

func (vd *VarDirectory) Exists(key string) bool {

	_, ok := vd.table[key]
	return ok
}

func (vd *VarDirectory) Table() map[string]*VarEntry {
	return vd.table
}

func CreateVarDirectoryFromVarEntries(ves []*VarEntry) (*VarDirectory, error) {
	vd := NewVarDirectory()
	for _, ve := range ves {
		if !vd.Add(ve) {
			return nil, errutil.Newf("Redeclaration of %s in %v", ve.Id(), ve.Token())
		}
	}

	return vd, nil
}

//NewVarDirectory New directory of variables that stores the var entry
func NewVarDirectory() *VarDirectory {
	return &VarDirectory{make(map[string]*VarEntry)}
}