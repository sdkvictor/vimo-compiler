package ic

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/sdkvictor/golang-compiler/directories"
	"github.com/sdkvictor/golang-compiler/mem"
	"github.com/sdkvictor/golang-compiler/quad"
	"github.com/sdkvictor/golang-compiler/types"
)

// Generator ...
type Generator struct {
	jumpStack       *AddressStack
	addrStack       *AddressStack
	typeStack		*TypeStack
	icounter        int
	pcounter        int
	quads           []*quad.Quadruple
	pendingFuncAddr map[int]string
	pendingEraSize  map[int]string
}

// NewGenerator ...
func NewGenerator() *Generator {
	return &Generator{NewAddressStack(), NewAddressStack(), NewTypeStack(), 0, 0, make([]*quad.Quadruple, 0), make(map[int]string), make(map[int]string)}
}

// JumpStack ...
func (g *Generator) JumpStack() *AddressStack {
	return g.jumpStack
}

// ICounter gets the current instruction counter
func (g *Generator) ICounter() int {
	return g.icounter
}

//GetNextPCounter ...
func (g *Generator) GetNextPCounter() int {
	val := g.pcounter
	g.pcounter++
	return val
}

//ResetPCounter ...
func (g *Generator) ResetPCounter() {
	g.pcounter = 0
}

// Quadruples ...
func (g *Generator) Quadruples() []*quad.Quadruple {
	return g.quads
}

// Generate creates a new quadruple with the given parameters
func (g *Generator) Generate(op quad.Operation, a1, a2, r mem.Address) {
	g.quads = append(g.quads, quad.NewQuadruple(op, a1, a2, r))
	//fmt.Printf("%d: Operation: %s %d %d %d\n", g.icounter, op, a1, a2, r)
	g.icounter++
}

//PushToAddrStack adds an address to the address stack
func (g *Generator) PushToAddrStack(a mem.Address) {
	g.addrStack.Push(a)
}

//GetFromAddrStack pops and returns the top address of the stack
func (g *Generator) GetFromAddrStack() mem.Address {
	val := g.addrStack.Top()
	g.addrStack.Pop()
	return val
}

//PushToJumpStack ...
func (g *Generator) PushToJumpStack(a mem.Address) {
	g.jumpStack.Push(a)
}

//GetFromJumpStack ...
func (g *Generator) GetFromJumpStack() mem.Address {
	val := g.jumpStack.Top()
	g.jumpStack.Pop()
	return val
}

//PushToAddrStack adds an address to the address stack
func (g *Generator) PushToTypeStack(a *types.Type) {
	g.typeStack.Push(a)
}

//GetFromAddrStack pops and returns the top address of the stack
func (g *Generator) GetFromTypeStack() *types.Type {
	val := g.typeStack.Top()
	g.typeStack.Pop()
	return val
}

//FillJumpQuadruple ...
func (g *Generator) FillJumpQuadruple(location mem.Address, jump mem.Address) {
	g.quads[int(location)].SetR(jump)
}

//AddPendingFuncAddr ...
func (g *Generator) AddPendingFuncAddr(loc int, id string) {
	g.pendingFuncAddr[loc] = id
}

//GetPendingFuncAddr ...
func (g *Generator) GetPendingFuncAddr() *map[int]string {
	return &g.pendingFuncAddr
}

//AddPendingEra ...
func (g *Generator) AddPendingEra(loc int, id string) {
	g.pendingEraSize[loc] = id
}

//GetPendingEraSize ...
func (g *Generator) GetPendingEraSize() *map[int]string {
	return &g.pendingEraSize
}

//FillPendingFuncAddr ...
func (g *Generator) FillPendingFuncAddr(funcdir *directories.FuncDirectory) {
	for loc, id := range g.pendingFuncAddr {
		fe := funcdir.Get(id)

		if g.quads[loc].Op() == quad.Goto {
			g.quads[loc].SetR(fe.Loc())
		} else {
			g.quads[loc].SetLop(fe.Loc())
		}
	}
}

//FillPendingEraFunctions ...
func (g *Generator) FillPendingEraFunctions(funcdir *directories.FuncDirectory) {

	for loc, id := range g.pendingEraSize {
		fe := funcdir.Get(id)
		g.quads[loc].SetLop(mem.Address(fe.Era()))
	}
}

//String ...
func (g *Generator) String() string {
	var builder strings.Builder
	builder.WriteString("Generator:\n")
	//builder.WriteString(fmt.Sprintf("  JumpStack: %s\n", g.jumpStack))
	//builder.WriteString(fmt.Sprintf("  AddrStack: %s\n", g.addrStack))
	builder.WriteString(fmt.Sprintf("  Instruction Counter: %d\n", g.icounter))
	builder.WriteString("  Quads:\n")

	for i, q := range g.quads {
		builder.WriteString(fmt.Sprintf("    %d: %s\n", i, q))
	}

	return builder.String()
}

func (g *Generator) CreateFile(file string, vm *mem.VirtualMemory) error {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("%d\n", len(g.quads)))

	for _, q := range g.quads {
		builder.WriteString(fmt.Sprintf("%s\n", q))
	}

	cm := vm.GetConstantMap()
	builder.WriteString(fmt.Sprintf("%d\n", len(cm)))

	for key, value := range cm {
		builder.WriteString(fmt.Sprintf("%s %d\n", key, value))
	}

	content := []byte(builder.String())
	path := fmt.Sprintf("./%s.obj", file)

	err := ioutil.WriteFile(path, content, 0644)

	return err
}