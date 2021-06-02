package mem

import (
	"fmt"

	"github.com/mewkiz/pkg/errutil"
	"github.com/sdkvictor/golang-compiler/types"
)

//Type of the address of the memory is an int
type Address int

//Constants that mark the start of the context
const Globalstart = 0
const Localstart = 10000
const Tempstart = 20000
const Constantstart = 30000
const Scopestart = 40000

//Constants that set the offset of the segment
const FloatOffset = 0
const CharOffset = 1000
const BoolOffset = 2000
const IntOffset = 3000
const StringOffset = 4000
const SquareOffset = 5000
const CircleOffset = 6000
const ImageOffset = 7000
const TextOffset = 8000
const BackgroundOffset = 9000

//Constant that defines de size of the segment
const segmentsize = 1000

//String If the address is not in range or not useful it must be set to -1 to be declared invalid
func (a Address) String() string {
	if a < 0 {
		return "-1"
	}

	return fmt.Sprintf("%d", a)
}

var DefaultConstants = map[string]string{
	"1": "0.0",
	"2": "a",
	"3": "false",
	"4": "0",
	"5": "-",
}

/*
Floats -> Chars -> Booleans -> Ints -> Strings -> Squares -> Circles -> Images -> Texts -> Backgrounds
========Global = 0
0-999     Float
1000-1999 Char
2000-2999 Bool
3000-3999 Int
4000-4999 String
5000-5999 Square
6000-6999 Circle
7000-7999 Image
8000-8999 Text
9000-9999 Background
========Local = 10000
10000-10999 Float
11000-11999 Char
12000-12999 Bool
13000-13999 Int
14000-14999 String
15000-15999 Square
16000-16999 Circle
17000-17999 Image
18000-18999 Text
19000-19999 Background
======= Temp = 20000
20000-20999 Float
21000-21999 Char
22000-22999 Bool
23000-23999 Int
24000-24999 String
25000-25999 Square
26000-26999 Circle
27000-27999 Image
28000-28999 Text
29000-29999 Background
========Constant = 30000
30000-30999 Float
31000-31999 Char
32000-32999 Bool
33000-33999 Int
34000-34999 String
35000-35999 Square
36000-36999 Circle
37000-37999 Image
38000-38999 Text
39000-39999 Background
*/

var objectSize = len(types.ObjectAttributesIndex)+1

//Struct that defines what the Virtual memory containts
type VirtualMemory struct {
	globalfloatcount      int
	globalcharcount       int
	globalboolcount       int
	globalintcount        int
	globalstringcount     int
	globalsquarecount     int
	globalcirclecount     int
	globalimagecount      int
	globaltextcount       int
	globalbackgroundcount int

	localfloatcount      int
	localcharcount       int
	localboolcount       int
	localintcount        int
	localstringcount     int
	localsquarecount     int
	localcirclecount     int
	localimagecount      int
	localtextcount       int
	localbackgroundcount int

	tempfloatcount      int
	tempcharcount       int
	tempboolcount       int
	tempintcount        int
	tempstringcount     int
	tempsquarecount     int
	tempcirclecount     int
	tempimagecount      int
	temptextcount       int
	tempbackgroundcount int

	constantfloatcount      int
	constantcharcount       int
	constantboolcount       int
	constantintcount        int
	constantstringcount     int
	constantsquarecount     int
	constantcirclecount     int
	constantimagecount      int
	constanttextcount       int
	constantbackgroundcount int

	scopefloatcount      int
	scopecharcount       int
	scopeboolcount       int
	scopeintcount        int
	scopestringcount     int
	scopesquarecount     int
	scopecirclecount     int
	scopeimagecount      int
	scopetextcount       int
	scopebackgroundcount int

	constantmap map[string]int
}

// NewVirtualMemory creates a new virtual memory
func NewVirtualMemory() *VirtualMemory {

	vm := &VirtualMemory{
		0, // globalfloatcount
		0, // globalcharcount
		0, // globalboolcount
		0, // globalintcount
		0, // globalstringcount
		0, // globalsquarecount
		0, // globalcirclecount
		0, // globalimagecount
		0, // globaltextcount
		0, // globalbackgroundcount

		0, // localfloatcount
		0, // localcharcount
		0, // localboolcount
		0, // localintcount
		0, // localstringcount
		0, // localsquarecount
		0, // localcirclecount
		0, // localimagecount
		0, // localtextcount
		0, // localbackgroundcount

		0, // tempfloatcount
		0, // tempcharcount
		0, // tempboolcount
		0, // tempintcount
		0, // tempstringcount
		0, // tempsquarecount
		0, // tempcirclecount
		0, // tempimagecount
		0, // temptextcount
		0, // tempbackgroundcount

		0, // constantfloatcount
		0, // constantcharcount
		0, // constantboolcount
		0, // constantintcount
		0, // constantstringcount
		0, // constantsquarecount
		0, // constantcirclecount
		0, // constantimagecount
		0, // constanttextcount
		0, // constantbackgroundcount

		0, // scopefloatcount
		0, // scopecharcount
		0, // scopeboolcount
		0, // scopeintcount
		0, // scopestringcount
		0, // scopesquarecount
		0, // scopecirclecount
		0, // scopeimagecount
		0, // scopetextcount
		0, // scopebackgroundcount

		make(map[string]int),
	}

	for t, v := range DefaultConstants {
		nt := types.NewDataType(types.ConvertInverse(t), 0, 0)
		_, _ = vm.AddConstant(v, nt)
	}

	return vm
}

// GetNextGlobal Receives the type and determines the next global variable available in the scope to assign it
func (vm *VirtualMemory) GetNextGlobal(t *types.Type) (Address, error) {
	nt, amount := normalizeType(t)
	objAmount := amount * objectSize
	switch nt.String() {
	// Float
	case "1":
		if vm.globalfloatcount+amount >= segmentsize {
			return Address(-1), errutil.Newf("Error: global variables for floats exceeded.")
		}
		result := vm.globalfloatcount + FloatOffset + Globalstart
		vm.globalfloatcount += amount
		return Address(result), nil
	// Char
	case "2":
		if vm.globalcharcount+amount >= segmentsize {
			return Address(-1), errutil.Newf("Error: global variables for chars exceeded.")
		}
		result := vm.globalcharcount + CharOffset + Globalstart
		vm.globalcharcount += amount
		return Address(result), nil
	// Bool
	case "3":
		if vm.globalboolcount+amount >= segmentsize {
			return Address(-1), errutil.Newf("Error: global variables for bools exceeded.")
		}
		result := vm.globalboolcount + BoolOffset + Globalstart
		vm.globalboolcount += amount
		return Address(result), nil
	// Int
	case "4":
		if vm.globalintcount+amount >= segmentsize {
			return Address(-1), errutil.Newf("Error: global variables for ints exceeded.")
		}
		result := vm.globalintcount + IntOffset + Globalstart
		vm.globalintcount += amount
		return Address(result), nil
	// String
	case "5":
		if vm.globalstringcount+amount >= segmentsize {
			return Address(-1), errutil.Newf("Error: global variables for strings exceeded.")
		}
		result := vm.globalstringcount + StringOffset + Globalstart
		vm.globalstringcount += amount
		return Address(result), nil
	// Square
	case "7":
		if vm.globalsquarecount+objAmount >= segmentsize {
			return Address(-1), errutil.Newf("Error: global variables for squares exceeded.")
		}
		result := vm.globalsquarecount + SquareOffset + Globalstart
		vm.globalsquarecount += objAmount
		return Address(result), nil
	// Circle
	case "8":
		if vm.globalcirclecount+objAmount >= segmentsize {
			return Address(-1), errutil.Newf("Error: global variables for circles exceeded.")
		}
		result := vm.globalcirclecount + CircleOffset + Globalstart
		vm.globalcirclecount += objAmount
		return Address(result), nil
	// Image
	case "9":
		if vm.globalimagecount+objAmount >= segmentsize {
			return Address(-1), errutil.Newf("Error: global variables for bools exceeded.")
		}
		result := vm.globalimagecount + ImageOffset + Globalstart
		vm.globalimagecount += objAmount
		return Address(result), nil
	// Text
	case "a":
		if vm.globaltextcount+objAmount >= segmentsize {
			return Address(-1), errutil.Newf("Error: global variables for texts exceeded.")
		}
		result := vm.globaltextcount + TextOffset + Globalstart
		vm.globaltextcount += objAmount
		return Address(result), nil
	// Background
	case "b":
		if vm.globalbackgroundcount+objAmount >= segmentsize {
			return Address(-1), errutil.Newf("Error: global variables for backgrounds exceeded.")
		}
		result := vm.globalbackgroundcount + BackgroundOffset + Globalstart
		vm.globalbackgroundcount += objAmount
		return Address(result), nil
	}

	return Address(-1), errutil.Newf("Error: variable type not identified.")
}

// GetNextLocal Receives the type and determines the next local variable available in the scope to assign it
func (vm *VirtualMemory) GetNextLocal(t *types.Type) (Address, error) {
	nt, amount := normalizeType(t)
	objAmount := amount * objectSize
	switch nt.String() {
	// Float
	case "1":
		if vm.localfloatcount+amount >= segmentsize {
			return Address(-1), errutil.Newf("Error: local variables for floats exceeded.")
		}
		result := vm.localfloatcount + FloatOffset + Localstart
		vm.localfloatcount += amount
		return Address(result), nil
	// Char
	case "2":
		if vm.localcharcount+amount >= segmentsize {
			return Address(-1), errutil.Newf("Error: local variables for chars exceeded.")
		}
		result := vm.localcharcount + CharOffset + Localstart
		vm.localcharcount += amount
		return Address(result), nil
	// Bool
	case "3":
		if vm.localboolcount+amount >= segmentsize {
			return Address(-1), errutil.Newf("Error: local variables for bools exceeded.")
		}
		result := vm.localboolcount + BoolOffset + Localstart
		vm.localboolcount += amount
		return Address(result), nil
	// Int
	case "4":
		if vm.localintcount+amount >= segmentsize {
			return Address(-1), errutil.Newf("Error: local variables for ints exceeded.")
		}
		result := vm.localintcount + IntOffset + Localstart
		vm.localintcount += amount
		return Address(result), nil
	// String
	case "5":
		if vm.localstringcount+amount >= segmentsize {
			return Address(-1), errutil.Newf("Error: local variables for strings exceeded.")
		}
		result := vm.localstringcount + StringOffset + Localstart
		vm.localstringcount += amount
		return Address(result), nil
	// Square
	case "7":
		if vm.localsquarecount+objAmount >= segmentsize {
			return Address(-1), errutil.Newf("Error: local variables for squares exceeded.")
		}
		result := vm.localsquarecount + SquareOffset + Localstart
		vm.localsquarecount += objAmount
		return Address(result), nil
	// Circle
	case "8":
		if vm.localcirclecount+objAmount >= segmentsize {
			return Address(-1), errutil.Newf("Error: local variables for circles exceeded.")
		}
		result := vm.localcirclecount + CircleOffset + Localstart
		vm.localcirclecount += objAmount
		return Address(result), nil
	// Image
	case "9":
		if vm.localimagecount+objAmount >= segmentsize {
			return Address(-1), errutil.Newf("Error: local variables for bools exceeded.")
		}
		result := vm.localimagecount + ImageOffset + Localstart
		vm.localimagecount += objAmount
		return Address(result), nil
	// Text
	case "a":
		if vm.localtextcount+objAmount >= segmentsize {
			return Address(-1), errutil.Newf("Error: local variables for texts exceeded.")
		}
		result := vm.localtextcount + TextOffset + Localstart
		vm.localtextcount += objAmount
		return Address(result), nil
	// Background
	case "b":
		if vm.localbackgroundcount+objAmount >= segmentsize {
			return Address(-1), errutil.Newf("Error: local variables for backgrounds exceeded.")
		}
		result := vm.localbackgroundcount + BackgroundOffset + Localstart
		vm.localbackgroundcount += objAmount
		return Address(result), nil
	}

	return Address(-1), errutil.Newf("Error: variable type not identified.")
}

// GetNextTemp Receives the type and determines the next temp variable available in the scope to assign it
func (vm *VirtualMemory) GetNextTemp(t *types.Type) (Address, error) {
	nt, amount := normalizeType(t)
	objAmount := amount * objectSize
	switch nt.String() {
	// Float
	case "1":
		if vm.tempfloatcount+amount >= segmentsize {
			return Address(-1), errutil.Newf("Error: temp variables for floats exceeded.")
		}
		result := vm.tempfloatcount + FloatOffset + Tempstart
		vm.tempfloatcount += amount
		return Address(result), nil
	// Char
	case "2":
		if vm.tempcharcount+amount >= segmentsize {
			return Address(-1), errutil.Newf("Error: temp variables for chars exceeded.")
		}
		result := vm.tempcharcount + CharOffset + Tempstart
		vm.tempcharcount += amount
		return Address(result), nil
	// Bool
	case "3":
		if vm.tempboolcount+amount >= segmentsize {
			return Address(-1), errutil.Newf("Error: temp variables for bools exceeded.")
		}
		result := vm.tempboolcount + BoolOffset + Tempstart
		vm.tempboolcount += amount
		return Address(result), nil
	// Int
	case "4":
		if vm.tempintcount+amount >= segmentsize {
			return Address(-1), errutil.Newf("Error: temp variables for ints exceeded.")
		}
		result := vm.tempintcount + IntOffset + Tempstart
		vm.tempintcount += amount
		return Address(result), nil
	// String
	case "5":
		if vm.tempstringcount+amount >= segmentsize {
			return Address(-1), errutil.Newf("Error: temp variables for strings exceeded.")
		}
		result := vm.tempstringcount + StringOffset + Tempstart
		vm.tempstringcount += amount
		return Address(result), nil
	// Void
	case "6":
		return Address(-1), nil
	// Square
	case "7":
		if vm.tempsquarecount+objAmount >= segmentsize {
			return Address(-1), errutil.Newf("Error: temp variables for squares exceeded.")
		}
		result := vm.tempsquarecount + SquareOffset + Tempstart
		vm.tempsquarecount += objAmount
		return Address(result), nil
	// Circle
	case "8":
		if vm.tempcirclecount+objAmount >= segmentsize {
			return Address(-1), errutil.Newf("Error: temp variables for circles exceeded.")
		}
		result := vm.tempcirclecount + CircleOffset + Tempstart
		vm.tempcirclecount += objAmount
		return Address(result), nil
	// Image
	case "9":
		if vm.tempimagecount+objAmount >= segmentsize {
			return Address(-1), errutil.Newf("Error: temp variables for images exceeded.")
		}
		result := vm.tempimagecount + ImageOffset + Tempstart
		vm.tempimagecount += objAmount
		return Address(result), nil
	// Text
	case "a":
		if vm.temptextcount+objAmount >= segmentsize {
			return Address(-1), errutil.Newf("Error: temp variables for texts exceeded.")
		}
		result := vm.temptextcount + TextOffset + Tempstart
		vm.temptextcount += objAmount
		return Address(result), nil
	// Background
	case "b":
		if vm.tempbackgroundcount+objAmount >= segmentsize {
			return Address(-1), errutil.Newf("Error: temp variables for backgrounds exceeded.")
		}
		result := vm.tempbackgroundcount + BackgroundOffset + Tempstart
		vm.tempbackgroundcount += objAmount
		return Address(result), nil
	}

	return Address(-1), errutil.Newf("Error: variable type not identified.")
}

// getNextConstant Receives the type and determines the next constant variable available in the scope to assign it
func (vm *VirtualMemory) getNextConstant(t *types.Type) (Address, error) {
	switch t.String() {
	// Float
	case "1":
		if vm.constantfloatcount >= segmentsize {
			return Address(-1), errutil.Newf("Error: constant variables for floats exceeded.")
		}
		result := vm.constantfloatcount + FloatOffset + Constantstart
		vm.constantfloatcount++
		return Address(result), nil
	// Char
	case "2":
		if vm.constantcharcount >= segmentsize {
			return Address(-1), errutil.Newf("Error: constant variables for chars exceeded.")
		}
		result := vm.constantcharcount + CharOffset + Constantstart
		vm.constantcharcount++
		return Address(result), nil
	// Bool
	case "3":
		if vm.constantboolcount >= segmentsize {
			return Address(-1), errutil.Newf("Error: constant variables for bools exceeded.")
		}
		result := vm.constantboolcount + BoolOffset + Constantstart
		vm.constantboolcount++
		return Address(result), nil
	// Int
	case "4":
		if vm.constantintcount >= segmentsize {
			return Address(-1), errutil.Newf("Error: constant variables for ints exceeded.")
		}
		result := vm.constantintcount + IntOffset + Constantstart
		vm.constantintcount++
		return Address(result), nil
	// String
	case "5":
		if vm.constantstringcount >= segmentsize {
			return Address(-1), errutil.Newf("Error: constant variables for strings exceeded.")
		}
		result := vm.constantstringcount + StringOffset + Constantstart
		vm.constantstringcount++
		return Address(result), nil
	// Square
	case "7":
		if vm.constantsquarecount >= segmentsize {
			return Address(-1), errutil.Newf("Error: constant variables for squares exceeded.")
		}
		result := vm.constantsquarecount + SquareOffset + Constantstart
		vm.constantsquarecount++
		return Address(result), nil
	// Circle
	case "8":
		if vm.constantcirclecount >= segmentsize {
			return Address(-1), errutil.Newf("Error: constant variables for circles exceeded.")
		}
		result := vm.constantcirclecount + CircleOffset + Constantstart
		vm.constantcirclecount++
		return Address(result), nil
	// Image
	case "9":
		if vm.constantimagecount >= segmentsize {
			return Address(-1), errutil.Newf("Error: constant variables for images exceeded.")
		}
		result := vm.constantimagecount + ImageOffset + Constantstart
		vm.constantimagecount++
		return Address(result), nil
	// Text
	case "a":
		if vm.constanttextcount >= segmentsize {
			return Address(-1), errutil.Newf("Error: constant variables for texts exceeded.")
		}
		result := vm.constanttextcount + TextOffset + Constantstart
		vm.constanttextcount++
		return Address(result), nil
	// Background
	case "b":
		if vm.constantbackgroundcount >= segmentsize {
			return Address(-1), errutil.Newf("Error: constant variables for backgrounds exceeded.")
		}
		result := vm.constantbackgroundcount + BackgroundOffset + Constantstart
		vm.constantbackgroundcount++
		return Address(result), nil
	}

	return Address(-1), errutil.Newf("Error: variable type not identified.")
}

// normalizeType converts Type list into normal type and returns the amount of elements in the
// list, together with the new type. If the type is already a non-list type, the function
// returns the same type with a 1 amount
func normalizeType(t *types.Type) (*types.Type, int) {
	nt := t.Copy()
	amount := 1
	if nt.List() > 0 {
		nt.DecreaseList()
		amount = nt.Size()
	}

	return nt, amount
}

func (vm *VirtualMemory) ResetLocal() {
	vm.localfloatcount = 0
	vm.localcharcount = 0
	vm.localboolcount = 0
	vm.localintcount = 0
	vm.localstringcount = 0
	vm.localsquarecount = 0
	vm.localcirclecount = 0
	vm.localimagecount = 0
	vm.localtextcount = 0
	vm.localbackgroundcount = 0
}

func (vm *VirtualMemory) ResetTemp() {
	vm.tempfloatcount = 0
	vm.tempcharcount = 0
	vm.tempboolcount = 0
	vm.tempintcount = 0
	vm.tempstringcount = 0
	vm.tempsquarecount = 0
	vm.tempcirclecount = 0
	vm.tempimagecount = 0
	vm.temptextcount = 0
	vm.tempbackgroundcount = 0
}

func (vm *VirtualMemory) ConstantExists(c string) bool {
	_, ok := vm.constantmap[c]
	return ok
}

func (vm *VirtualMemory) AddConstant(c string, t *types.Type) (Address, error) {
	if vm.ConstantExists(c) {
		addr := Address(vm.constantmap[c])
		return addr, nil
	}

	// TODO: Determine the type of the constant and address accordingly
	nextAddr, err := vm.getNextConstant(t)
	if err != nil {
		return Address(-1), err
	}
	vm.constantmap[c] = int(nextAddr)

	return Address(nextAddr), nil
}

//GetConstantAddress gets the address of the constant from the map of constants
func (vm *VirtualMemory) GetConstantAddress(c string) Address {
	a, ok := vm.constantmap[c]
	if !ok {
		return Address(-1)
	}

	return Address(a)
}

func (vm *VirtualMemory) GetConstantMap() map[string]int {
	return vm.constantmap
}

func (vm *VirtualMemory) GetDefaultAddress(t *types.Type) Address {
	defValue, ok := DefaultConstants[t.String()]
	if !ok {
		return Address(-1)
	}

	return vm.GetConstantAddress(defValue)
}
