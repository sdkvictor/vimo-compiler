package vm

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sdkvictor/golang-compiler/mem"
	"github.com/sdkvictor/golang-compiler/quad"
	"github.com/sdkvictor/golang-compiler/engine"
	"github.com/sdkvictor/golang-compiler/objects"
	"github.com/sdkvictor/golang-compiler/vm/ar"
	"github.com/mewkiz/pkg/errutil"
)

// VirtualMachine contains the necesary attributes of a virtual machine to execute sequential code,
// manage memory, and call submodules
type VirtualMachine struct {
	ip           int
	quads        []*quad.Quadruple
	mm           *Memory
	ar           *ar.ArStack
	pendingcalls *ar.ArStack
	engine 	     *engine.Engine
}

// String represents the vm in a strctured format so that it can be easily debugged
func (vm *VirtualMachine) String() string {
	var builder strings.Builder

	builder.WriteString("VirtualMachine:\n")
	builder.WriteString(fmt.Sprintf("  IP: %d\n", vm.ip))
	builder.WriteString("  Quads:\n")


	for i, q := range vm.quads {
		builder.WriteString(fmt.Sprintf("    %d: %s\n", i, q))
	}

	builder.WriteString("\n")

	builder.WriteString(vm.mm.String())

	return builder.String()
}

// loadConstants takes an array of strings and converts them to the corresponding constants in memory
func (vm *VirtualMachine) LoadConstants(constantmap map[string]int) error {
	for cons, addr := range constantmap {
		switch {
		case addr < mem.Constantstart+mem.CharOffset: // Float
			f, err := strconv.ParseFloat(cons, 64)
			if err != nil {
				return err
			}

			if err := vm.mm.SetValue(f, mem.Address(addr)); err != nil {
				return err
			}
		case addr < mem.Constantstart+mem.BoolOffset: // Char
			char := rune(cons[0])

			if err := vm.mm.SetValue(char, mem.Address(addr)); err != nil {
				return err
			}
		case addr < mem.Constantstart+mem.IntOffset: // Bool
			boolean, err := strconv.ParseBool(cons)
			if err != nil {
				return err
			}

			if err := vm.mm.SetValue(boolean, mem.Address(addr)); err != nil {
				return err
			}
		case addr < mem.Constantstart+mem.StringOffset: // Int
			i, err := strconv.Atoi(cons)
			if err != nil {
				return err
			}

			if err := vm.mm.SetValue(i, mem.Address(addr)); err != nil {
				return err
			}
		case addr < mem.Constantstart+mem.SquareOffset: // String
			if err := vm.mm.SetValue(cons, mem.Address(addr)); err != nil {
				return err
			}
		default:
			return errutil.Newf("Cannot set non-constant value")
		}
	}

	return nil
}

func (vm *VirtualMachine) printOutput(v interface{}) {
	if c, ok := v.(rune); ok {
		fmt.Printf("%c\n", c)
	} else if f, ok := v.(float64); ok {
		if f == float64(int64(f)) {
			fmt.Printf("%d\n", int64(f))
		} else {
			fmt.Printf("%f\n", f)
		}
	} else if o, ok := v.(objects.Object); ok {
		fmt.Printf("%s\n", o.String())
	} else {
		fmt.Printf("%v\n", v)
	}
}

// executeNextInstruction indexes the next quadruple with the instruction pointer and
// proceeds to execute that instruction
func (vm *VirtualMachine) executeNextInstruction() error {
	q := vm.quads[vm.ip]

	//fmt.Printf("%d: Operation: %s %d %d %d\n", vm.ip, q.Op().String(), q.Lop(), q.Rop(), q.R())

	switch q.Op() {
	case quad.Add:
		if err := vm.operationAdd(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Sub:
		if err := vm.operationSub(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Mult:
		if err := vm.operationMult(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Div:
		if err := vm.operationDiv(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.And:
		if err := vm.operationAnd(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Or:
		if err := vm.operationOr(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Not:
		if err := vm.operationNot(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Gt:
		if err := vm.operationGt(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Lt:
		if err := vm.operationLt(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Equal:
		if err := vm.operationEqual(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Print:
		if err := vm.operationPrint(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Era:
		if err := vm.operationEra(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Param:
		if err := vm.operationParam(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Call:
		if err := vm.operationCall(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
	case quad.Ret:
		if err := vm.operationRet(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
	case quad.Assign:
		if err := vm.operationAssign(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Goto:
		if err := vm.operationGoto(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
	case quad.GotoT:
		if err := vm.operationGotoT(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
	case quad.GotoF:
		if err := vm.operationGotoF(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
	case quad.Pow:
		if err := vm.operationPow(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Sqrt:
		if err := vm.operationSqrt(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Init:
		if err := vm.operationInit(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.CheckBound:
		if err := vm.operationCheckBound(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.AddAddr:
		if err := vm.operationAddAddr(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.AssignIndex:
		if err := vm.operationAssignIndex(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.AssignIndexInv:
		if err := vm.operationAssignIndexInv(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Render:
		if err := vm.operationRender(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Clear:
		if err := vm.operationClear(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	case quad.Update:
		if err := vm.operationUpdate(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	
	case quad.KeyPressed:
		if err := vm.operationKeyPressed(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	
	case quad.CheckCollision:
		if err := vm.operationCheckCollision(q.Lop(), q.Rop(), q.R()); err != nil {
			return err
		}
		vm.ip++
	default:
		return errutil.Newf("Invalid Quad %s", q.Op().String())
	}

	if vm.engine.WindowClosed() {
		return nil
	}

	return nil
}

// Run starts executing the instructions in the virtual machine
func (vm *VirtualMachine) Run() error {
	if len(vm.quads) < 1 {
		return errutil.NewNoPosf("No instructions to execute")
	}
	

	//Setup pixel

	// Create the main activation record
	mainAR := ar.NewActivationRecord()
	// Set return to the end of the quads to end execution
	mainAR.SetRetIp(len(vm.quads))
	vm.ar.Push(mainAR)

	for vm.ip < len(vm.quads) {
		if err := vm.executeNextInstruction(); err != nil {
			return err
		}
	}

	return nil
}
/*
// NewVirtualMachine default
func NewVirtualMachine() *VirtualMachine {
	return &VirtualMachine{0, make([]*quad.Quadruple, 0), NewMemory(), ar.NewArStack(), ar.NewArStack(), 0}
}
*/
// NewVirtualMachine custom
func NewVirtualMachine(quads []*quad.Quadruple, consmap map[string]int) *VirtualMachine {
	return &VirtualMachine{0, quads, NewMemory(), ar.NewArStack(), ar.NewArStack(), engine.NewEngine("Ping Pong", 730, 500)}
}