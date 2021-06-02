package vm

import (
	"math"

	"github.com/sdkvictor/golang-compiler/mem"
	"github.com/sdkvictor/golang-compiler/objects"
	"github.com/sdkvictor/golang-compiler/semantics"
	"github.com/sdkvictor/golang-compiler/vm/ar"
	"github.com/mewkiz/pkg/errutil"
)

func (vm *VirtualMachine) operationAdd(lop, rop, r mem.Address) error {
	//fmt.Printf("ADD %s\n", vm.mm)
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	ropv, err := vm.mm.GetValue(rop)
	if err != nil {
		return err
	}


	if f1, f2, err := getFloats(lopv, ropv); err == nil {
		result := f1 + f2
		if err := vm.mm.SetValue(result, r); err != nil {
			return err
		}
		return nil
	} else if i1, i2, err := getInts(lopv, ropv); err == nil {
		result := i1 + i2
		if err := vm.mm.SetValue(result, r); err != nil {
			return err
		}
		return nil
	}

	return nil
}

func (vm *VirtualMachine) operationSub(lop, rop, r mem.Address) error {
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	ropv, err := vm.mm.GetValue(rop)
	if err != nil {
		return err
	}

	if f1, f2, err := getFloats(lopv, ropv); err == nil {
		result := f1 - f2
		if err := vm.mm.SetValue(result, r); err != nil {
			return err
		}
		return nil
	} else if i1, i2, err := getInts(lopv, ropv); err == nil {
		result := i1 - i2
		if err := vm.mm.SetValue(result, r); err != nil {
			return err
		}
		return nil
	}

	return nil
}

func (vm *VirtualMachine) operationMult(lop, rop, r mem.Address) error {
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	ropv, err := vm.mm.GetValue(rop)
	if err != nil {
		return err
	}

	if f1, f2, err := getFloats(lopv, ropv); err == nil {
		result := f1 * f2
		if err := vm.mm.SetValue(result, r); err != nil {
			return err
		}
		return nil
	} else if i1, i2, err := getInts(lopv, ropv); err == nil {
		result := i1 * i2
		if err := vm.mm.SetValue(result, r); err != nil {
			return err
		}
		return nil
	}

	return nil
}

func (vm *VirtualMachine) operationDiv(lop, rop, r mem.Address) error {
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	ropv, err := vm.mm.GetValue(rop)
	if err != nil {
		return err
	}

	if f1, f2, err := getFloats(lopv, ropv); err == nil {
		if f2 == 0 {
			return errutil.Newf("Arithmethic exception, division by 0")
		}
		result := f1 / f2
		if err := vm.mm.SetValue(result, r); err != nil {
			return err
		}
		return nil
	} else if i1, i2, err := getInts(lopv, ropv); err == nil {
		if i2 == 0 {
			return errutil.Newf("Arithmethic exception, division by 0")
		}
		result := i1 / i2
		if err := vm.mm.SetValue(result, r); err != nil {
			return err
		}
		return nil
	}

	return nil
}


func (vm *VirtualMachine) operationAnd(lop, rop, r mem.Address) error {
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	ropv, err := vm.mm.GetValue(rop)
	if err != nil {
		return err
	}

	b1, err := getBool(lopv)
	if err != nil {
		return err
	}

	b2, err := getBool(ropv)
	if err != nil {
		return err
	}

	result := b1 && b2

	if err := vm.mm.SetValue(result, r); err != nil {
		return err
	}

	return nil
}

func (vm *VirtualMachine) operationOr(lop, rop, r mem.Address) error {
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	ropv, err := vm.mm.GetValue(rop)
	if err != nil {
		return err
	}

	b1, err := getBool(lopv)
	if err != nil {
		return err
	}

	b2, err := getBool(ropv)
	if err != nil {
		return err
	}

	result := b1 || b2

	if err := vm.mm.SetValue(result, r); err != nil {
		return err
	}

	return nil
}

func (vm *VirtualMachine) operationNot(lop, rop, r mem.Address) error {
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	b1, err := getBool(lopv)
	if err != nil {
		return err
	}

	result := !b1

	if err := vm.mm.SetValue(result, r); err != nil {
		return err
	}

	return nil
}

func (vm *VirtualMachine) operationGt(lop, rop, r mem.Address) error {
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	ropv, err := vm.mm.GetValue(rop)
	if err != nil {
		return err
	}

	if f1, f2, err := getFloats(lopv, ropv); err == nil {
		result := f1 > f2
		if err := vm.mm.SetValue(result, r); err != nil {
			return err
		}
		return nil
	} else if i1, i2, err := getInts(lopv, ropv); err == nil {
		result := i1 > i2
		if err := vm.mm.SetValue(result, r); err != nil {
			return err
		}
		return nil
	}

	return nil
}

func (vm *VirtualMachine) operationLt(lop, rop, r mem.Address) error {
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	ropv, err := vm.mm.GetValue(rop)
	if err != nil {
		return err
	}

	if f1, f2, err := getFloats(lopv, ropv); err == nil {
		result := f1 < f2
		if err := vm.mm.SetValue(result, r); err != nil {
			return err
		}
		return nil
	} else if i1, i2, err := getInts(lopv, ropv); err == nil {
		result := i1 < i2
		if err := vm.mm.SetValue(result, r); err != nil {
			return err
		}
		return nil
	}

	return nil
}

func (vm *VirtualMachine) operationEqual(lop, rop, r mem.Address) error {
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	ropv, err := vm.mm.GetValue(rop)
	if err != nil {
		return err
	}

	if f1, f2, err := getFloats(lopv, ropv); err == nil {
		result := f1 == f2
		if err := vm.mm.SetValue(result, r); err != nil {
			return err
		}
		return nil
	} else if c1, c2, err := getChars(lopv, ropv); err == nil {
		result := c1 == c2
		if err := vm.mm.SetValue(result, r); err != nil {
			return err
		}
		return nil
	} else if b1, b2, err := getBools(lopv, ropv); err == nil {
		result := b1 == b2
		if err := vm.mm.SetValue(result, r); err != nil {
			return err
		}
		return nil
	} else if i1, i2, err := getInts(lopv, ropv); err == nil {
		result := i1 == i2
		if err := vm.mm.SetValue(result, r); err != nil {
			return err
		}
		return nil
	} else if s1, s2, err := getStrings(lopv, ropv); err == nil {
		result := s1 == s2
		if err := vm.mm.SetValue(result, r); err != nil {
			return err
		}
		return nil
	}

	return errutil.Newf("Cannot perform equal operation on given types")
}

func (vm *VirtualMachine) operationPrint(lop, rop, r mem.Address) error {
	result, err := vm.mm.GetValue(r)
	if err != nil {
		return err
	}

	vm.printOutput(result)

	return nil
}

func (vm *VirtualMachine) operationEra(lop, rop, r mem.Address) error {
	vm.pendingcalls.Push(ar.NewActivationRecord())
	return nil
}

func (vm *VirtualMachine) operationParam(lop, rop, r mem.Address) error {
	nextcall := vm.pendingcalls.Top()

	typedaddr := getTypeAddr(lop)

	size := int(rop)

	if size < 2 {
		size = 1
	}

	for i := 0; i < size; i++ {
		lopv, err := vm.mm.GetValue(mem.Address(int(lop) + i))
		if err != nil {
			return err
		}
	
		switch {
		case typedaddr < mem.CharOffset:
			f1, err := getFloat(lopv)
			if err != nil {
				return err
			}
			nextcall.AddFloatParam(f1)
		case typedaddr < mem.BoolOffset:
			c1, err := getChar(lopv)
			if err != nil {
				return err
			}
			nextcall.AddCharParam(c1)
		case typedaddr < mem.IntOffset:
			b1, err := getBool(lopv)
			if err != nil {
				return err
			}
			nextcall.AddBoolParam(b1)
		case typedaddr < mem.StringOffset:
			i1, err := getInt(lopv)
			if err != nil {
				return err
			}
			nextcall.AddIntParam(i1)	
		case typedaddr < mem.SquareOffset:
			s1, err := getString(lopv)
			if err != nil {
				return err
			}
			nextcall.AddStringParam(s1)
		case typedaddr < mem.CircleOffset:  // Square
			s1, err := getSquare(lopv)
			if err != nil {
				return err
			}
			nextcall.AddSquareParam(s1)
		case typedaddr < mem.ImageOffset:  // Circle
			s1, err := getCircle(lopv)
			if err != nil {
				return err
			}
			nextcall.AddCircleParam(s1)
		case typedaddr < mem.TextOffset: // Image
			s1, err := getImage(lopv)
			if err != nil {
				return err
			}
			nextcall.AddImageParam(s1)
		case typedaddr < mem.BackgroundOffset: // Text
			s1, err := getText(lopv)
			if err != nil {
				return err
			}
			nextcall.AddTextParam(s1)
		default:							// Background
			s1, err := getBackground(lopv)
			if err != nil {
				return err
			}
			nextcall.AddBackgroundParam(s1)
		}
		

	}

	return nil
}

func (vm *VirtualMachine) operationCall(lop, rop, r mem.Address) error {
	// First we get the current call
	currcall := vm.ar.Top()
	//fmt.Printf("CURR CALL %+v\n", currcall)
	// Now we get the call which is about to happen
	nextcall := vm.pendingcalls.Top()

	// And we remove it from the pending calls
	// This allows us to have calls within calls
	vm.pendingcalls.Pop()

	// Now we copy the temp values of the current call to its activation record
	// so that it can be restored later
	if err := vm.copyTempToAR(currcall); err != nil {
		return err
	}

	if err := vm.copyLocalToAR(currcall); err != nil {
		return err
	}

	// Now we initialize the local memory with the values of the parameters of the
	// function
	if err := vm.copyParamsToLocal(nextcall); err != nil {
		return err
	}


	// Now we set the return IP of the incoming call to the current ip
	nextcall.SetRetIp(vm.ip)

	// We add the incoming call to the activation stack
	vm.ar.Push(nextcall)

	// Now we get the location for the function
	jump := int(lop)

	if jump < 0 {
		return errutil.Newf("Invalid instruction address")
	}

	// And finally we set the current ip to the new location
	vm.ip = jump

	return nil
}

func (vm *VirtualMachine) operationRet(lop, rop, r mem.Address) error {
	//fmt.Printf("RETURN %s\n", vm.mm)
	// We get the return instruction pointer
	retip := vm.ar.Top().Retip()
	
	// We kill the current activation record
	vm.ar.Pop()
	
	newcurrcall := vm.ar.Top()
	
	// If the stack is empty, then we are out of the main function 
	// and we can end execution
	if newcurrcall == nil {
		//fmt.Println("FINISHED?")
		vm.ip = retip + 1
		return nil
	}

	void := false
	var retv interface{}
	var err error
	if int(r) == -1 {
		// If the return address is -1, then it's a void function
		void = true
	} else {
		// If its not void, We get the value of the return value
		retv, err = vm.mm.GetValue(r)
		if err != nil {
			return err
		}
	}
	

	// We reactivate the local variables to the main memory
	if err := vm.copyParamsToLocal(newcurrcall); err != nil {
		return err
	}
	
	// We copy back the temp values to the main memory
	if err := vm.copyTempToMemory(newcurrcall); err != nil {
		return err
	}

	// We check if the function is a void function
	if void {
		// If it's a void, then we just update the ip and return
		vm.ip = retip + 1
		return nil
	}
	
	// We get the address where we need to save the return value
	callSaveAddr := vm.quads[retip].R()

	if f, err := getFloat(retv); err == nil {
		if err := vm.mm.SetValue(f, callSaveAddr); err != nil {
			return err
		}
	} else if c, err := getChar(retv); err == nil {
		if err := vm.mm.SetValue(c, callSaveAddr); err != nil {
			return err
		}
	} else if b, err := getBool(retv); err == nil {
		if err := vm.mm.SetValue(b, callSaveAddr); err != nil {
			return err
		}
	} else if i, err := getInt(retv); err == nil {
		if err := vm.mm.SetValue(i, callSaveAddr); err != nil {
			return err
		}
	} else if s, err := getString(retv); err == nil {
		if err := vm.mm.SetValue(s, callSaveAddr); err != nil {
			return err
		}
	} else if i, err := getSquare(retv); err == nil {
		if err := vm.mm.SetValue(i, callSaveAddr); err != nil {
			return err
		}
	}else if i, err := getCircle(retv); err == nil {
		if err := vm.mm.SetValue(i, callSaveAddr); err != nil {
			return err
		}
	}else if i, err := getImage(retv); err == nil {
		if err := vm.mm.SetValue(i, callSaveAddr); err != nil {
			return err
		}
	}else if i, err := getText(retv); err == nil {
		if err := vm.mm.SetValue(i, callSaveAddr); err != nil {
			return err
		}
	} else if i, err := getBackground(retv); err == nil {
		if err := vm.mm.SetValue(i, callSaveAddr); err != nil {
			return err
		}
	}else {
		return errutil.Newf("Cannot get valid form for value")
	}

	// We update the ip with the saved value
	vm.ip = retip + 1

	return nil
}

func (vm *VirtualMachine) operationAssign(lop, rop, r mem.Address) error {
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	if err := vm.mm.SetValue(lopv, r); err != nil {
		return err
	}

	return nil
}

func (vm *VirtualMachine) operationGoto(lop, rop, r mem.Address) error {
	if r < 0 {
		return errutil.Newf("Invalid instruction address")
	}

	vm.ip = int(r)

	return nil
}

func (vm *VirtualMachine) operationGotoT(lop, rop, r mem.Address) error {
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	b1, err := getBool(lopv)
	if err != nil {
		return err
	}

	jump := int(r)

	if b1 {
		vm.ip = jump
	} else {
		vm.ip++
	}

	return nil
}

func (vm *VirtualMachine) operationGotoF(lop, rop, r mem.Address) error {
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	b1, err := getBool(lopv)
	if err != nil {
		return err
	}

	jump := int(r)

	if !b1 {
		vm.ip = jump
	} else {
		vm.ip++
	}

	return nil
}

func (vm *VirtualMachine) operationPow(lop, rop, r mem.Address) error {
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	ropv, err := vm.mm.GetValue(rop)
	if err != nil {
		return err
	}

	if f1, f2, err := getFloats(lopv, ropv); err == nil {
		result := math.Pow(f1, f2)
		if err := vm.mm.SetValue(result, r); err != nil {
			return err
		}
		return nil
	} else if i1, i2, err := getInts(lopv, ropv); err == nil {
		result := math.Pow(float64(i1), float64(i2))
		if err := vm.mm.SetValue(result, r); err != nil {
			return err
		}
		return nil
	}

	return nil
}

func (vm *VirtualMachine) operationSqrt(lop, rop, r mem.Address) error {
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	if f1, err := getFloat(lopv); err == nil {
		result := math.Sqrt(f1)
		if err := vm.mm.SetValue(result, r); err != nil {
			return err
		}
		return nil
	} else if i1, err := getInt(lopv); err == nil {
		result := math.Sqrt(float64(i1))
		if err := vm.mm.SetValue(result, r); err != nil {
			return err
		}
		return nil
	}

	return nil
}

func (vm *VirtualMachine) operationInit(lop, rop, r mem.Address) error {
	startAddr := int(lop)
	size := int(rop)
	t := int (r)
	for i := 0; i < size; i++ {
		switch t {
		case 1:
			if err := vm.mm.SetValue(0.0, mem.Address(startAddr+i)); err != nil {
				return err
			}
		case 2:
			if err := vm.mm.SetValue('a', mem.Address(startAddr+i)); err != nil {
				return err
			}
		case 3:
			if err := vm.mm.SetValue(false, mem.Address(startAddr+i)); err != nil {
				return err
			}
		case 4:
			if err := vm.mm.SetValue(0, mem.Address(startAddr+i)); err != nil {
				return err
			}
		case 5:
			if err := vm.mm.SetValue(" ", mem.Address(startAddr+i)); err != nil {
				return err
			}
		case 6:
		case 7:
			if err := vm.mm.SetValue(objects.Square{}, mem.Address(startAddr+(i*semantics.ObjectSize))); err != nil {
				return err
			}
		case 8:
			if err := vm.mm.SetValue(objects.Circle{}, mem.Address(startAddr+(i*semantics.ObjectSize))); err != nil {
				return err
			}
		case 9:
			if err := vm.mm.SetValue(objects.Image{}, mem.Address(startAddr+(i*semantics.ObjectSize))); err != nil {
				return err
			}
		case 10:
			if err := vm.mm.SetValue(objects.Text{}, mem.Address(startAddr+(i*semantics.ObjectSize))); err != nil {
				return err
			}
		case 11:
			if err := vm.mm.SetValue(objects.Background{}, mem.Address(startAddr+(i*semantics.ObjectSize))); err != nil {
				return err
			}
		default:
			return errutil.Newf("Type not identified")
		}
	}

	return nil
}

func (vm *VirtualMachine) operationCheckBound(lop, rop, r mem.Address) error {
	size := int(lop)
	o, err := vm.mm.GetValue(r)
	if err != nil {
		return err
	}

	offset, err := getInt(o)
	if err != nil {
		return err
	}

	//fmt.Printf("OFFSET: %d\n", offset)

	if offset >= size {
		return errutil.Newf("Index %d out of bounds for array of size %d", offset, size)
	}

	return nil
}

func (vm *VirtualMachine) operationAddAddr(lop, rop, r mem.Address) error {
	startAddr := int(lop)

	o, err := vm.mm.GetValue(rop)
	if err != nil {
		return err
	}

	offset, err := getInt(o)
	if err != nil {
		return err
	}

	realOffset, err := getOffsetObject(offset, lop)
	if err != nil {
		return err
	}

	finalAddr := realOffset + startAddr
	if err := vm.mm.SetValue(finalAddr, r); err != nil {
		return err
	}

	return nil
}

func (vm *VirtualMachine) operationAssignIndex(lop, rop, r mem.Address) error {
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	arrAddrv, err := vm.mm.GetValue(r)
	if err != nil {
		return err
	}

	arrAddr, err := getInt(arrAddrv)
	if err != nil {
		return err
	}

	if err := vm.mm.SetValue(lopv, mem.Address(arrAddr)); err != nil {
		return err
	}

	return nil
}

func (vm *VirtualMachine) operationAssignIndexInv(lop, rop, r mem.Address) error {
	finalAddrv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	finalAddr, err := getInt(finalAddrv)
	if err != nil {
		return err
	}

	res, err := vm.mm.GetValue(mem.Address(finalAddr))
	if err != nil {
		return err
	}


	if err := vm.mm.SetValue(res, r); err != nil {
		return err
	}

	return nil
}

func (vm *VirtualMachine) operationRender(lop, rop, r mem.Address) error { //ASSIGNED TO MOISES
	// lop address del objecto a renderizar
	// rop no se usa
	// ni r

	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	if i, err := getSquare(lopv); err == nil {
		vm.engine.DrawSquare(i)
	}else if i, err := getCircle(lopv); err == nil {
		vm.engine.DrawCircle(i)
	}else if i, err := getImage(lopv); err == nil {
		vm.engine.DrawImage(i)
	}else if i, err := getText(lopv); err == nil {
		vm.engine.DrawText(i)
	} else if _, err := getBackground(lopv); err == nil {
		return nil
	}else {
		return errutil.Newf("Cannot get valid form for value")
	}
	return nil
}

func (vm *VirtualMachine) operationClear(lop, rop, r mem.Address) error { 
	vm.engine.Clear()
	return nil
}

func (vm *VirtualMachine) operationUpdate(lop, rop, r mem.Address) error {
	vm.engine.Update()
	return nil
}
		

func (vm *VirtualMachine) operationKeyPressed(lop, rop, r mem.Address) error { 
	k, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}
	key, err := getString(k)
	if err != nil {
		return err
	}
	res := vm.engine.KeyPressed(key)
	if err := vm.mm.SetValue(res, r); err != nil {
		return err
	}

	return nil
}
func (vm *VirtualMachine) operationCheckCollision(lop, rop, r mem.Address) error { 
	// lop tiene objeto 1
	// rop tiene el objeto 2
	// r guardar el resultado de la colision
	lopv, err := vm.mm.GetValue(lop)
	if err != nil {
		return err
	}

	ropv, err := vm.mm.GetValue(rop)
	if err != nil {
		return err
	}

	if l, err := getSquare(lopv); err==nil{ 
		if l2, err := getCircle(ropv); err==nil{
			res := vm.engine.IntersectSC(l, l2) //is Square - Circle
			if err := vm.mm.SetValue(res, r); err != nil {
				return err
			}
		}else if l2, err := getSquare(ropv); err==nil{
			res := vm.engine.IntersectSquare(l,l2) //is Square - Square
			if err := vm.mm.SetValue(res, r); err != nil {
				return err
			}
		}else{
			return errutil.Newf("Cannot get valid form for collision ")
		}
	} else if l, err := getCircle(lopv); err==nil{
		if l2, err := getCircle(ropv); err==nil{
			res := vm.engine.IntersectCircle(l, l2) //is Circle - Circle
			if err := vm.mm.SetValue(res, r); err != nil {
				return err
			}
		}else if l2, err := getSquare(ropv); err==nil{
			res := vm.engine.IntersectCS(l,l2) //is Circle - Square
			if err := vm.mm.SetValue(res, r); err != nil {
				return err
			}
		}else{
			return errutil.Newf("Cannot get valid form for collision ")
		}
	}else{
		return errutil.Newf("Cannot get valid form for collision ")
	}
	return nil
}
