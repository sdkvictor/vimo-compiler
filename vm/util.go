package vm

import (
	"os"

	"github.com/sdkvictor/golang-compiler/mem"
	"github.com/sdkvictor/golang-compiler/semantics"
	"github.com/sdkvictor/golang-compiler/vm/ar"
	"github.com/mewkiz/pkg/errutil"
	"github.com/sdkvictor/golang-compiler/objects"
)

var objSize = semantics.ObjectSize

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

func getFloat(v interface{}) (float64, error) {
	f, ok := v.(float64)
	if !ok {
		return 0, errutil.NewNoPosf("Cannot convert current value to float")
	}

	return f, nil
}

func getFloats(v1, v2 interface{}) (float64, float64, error) {
	f1, ok := v1.(float64)
	if !ok {
		return 0, 0, errutil.NewNoPosf("Cannot convert current value to float")
	}

	f2, ok := v2.(float64)
	if !ok {
		return 0, 0, errutil.NewNoPosf("Cannot convert current value to float")
	}

	return f1, f2, nil
}

func getChar(v interface{}) (rune, error) {
	char, ok := v.(rune)
	if !ok {
		return 0, errutil.NewNoPosf("Cannot convert current value to char")
	}

	return char, nil
}

func getChars(v1, v2 interface{}) (rune, rune, error) {
	char1, ok := v1.(rune)
	if !ok {
		return 0, 0, errutil.NewNoPosf("Cannot convert current value to char")
	}

	char2, ok := v2.(rune)
	if !ok {
		return 0, 0, errutil.NewNoPosf("Cannot convert current value to char")
	}

	return char1, char2, nil
}

func getBool(v interface{}) (bool, error) {
	boo, ok := v.(bool)
	if !ok {
		return false, errutil.NewNoPosf("Cannot convert current value to bool")
	}

	return boo, nil
}

func getBools(v1, v2 interface{}) (bool, bool, error) {
	bool1, ok := v1.(bool)
	if !ok {
		return false, false, errutil.NewNoPosf("Cannot convert current value to bool")
	}

	bool2, ok := v2.(bool)
	if !ok {
		return false, false, errutil.NewNoPosf("Cannot convert current value to bool")
	}

	return bool1, bool2, nil
}

func getInt(v interface{}) (int, error) {
	in, ok := v.(int)
	if !ok {
		return 0, errutil.NewNoPosf("Cannot convert current value to address")
	}

	return in, nil
}

func getInts(v1, v2 interface{}) (int, int, error) {
	i1, ok := v1.(int)
	if !ok {
		return 0, 0, errutil.NewNoPosf("Cannot convert current value to int")
	}

	i2, ok := v2.(int)
	if !ok {
		return 0, 0, errutil.NewNoPosf("Cannot convert current value to int")
	}

	return i1, i2, nil
}

func getString(v interface{}) (string, error) {
	in, ok := v.(string)
	if !ok {
		return "", errutil.NewNoPosf("Cannot convert current value to address")
	}

	return in, nil
}


func getStrings(v1, v2 interface{}) (string, string, error) {
	i1, ok := v1.(string)
	if !ok {
		return "", "", errutil.NewNoPosf("Cannot convert current value to string")
	}

	i2, ok := v2.(string)
	if !ok {
		return "", "", errutil.NewNoPosf("Cannot convert current value to string")
	}

	return i1, i2, nil
}

func getSquare(v interface{}) (objects.Square, error) {
	in, ok := v.(objects.Square)
	if !ok {
		return objects.Square{}, errutil.NewNoPosf("Cannot convert current value to address")
	}

	return in, nil
}

func getCircle(v interface{}) (objects.Circle, error) {
	in, ok := v.(objects.Circle)
	if !ok {
		return objects.Circle{}, errutil.NewNoPosf("Cannot convert current value to address")
	}

	return in, nil
}

func getImage(v interface{}) (objects.Image, error) {
	in, ok := v.(objects.Image)
	if !ok {
		return objects.Image{}, errutil.NewNoPosf("Cannot convert current value to address")
	}

	return in, nil
}

func getText(v interface{}) (objects.Text, error) {
	in, ok := v.(objects.Text)
	if !ok {
		return objects.Text{}, errutil.NewNoPosf("Cannot convert current value to address")
	}

	return in, nil
}

func getBackground(v interface{}) (objects.Background, error) {
	in, ok := v.(objects.Background)
	if !ok {
		return objects.Background{}, errutil.NewNoPosf("Cannot convert current value to address")
	}

	return in, nil
}


func getTypeAddr(addr mem.Address) int {
	switch {
	case addr < mem.Localstart:
		return int(addr) - mem.Globalstart
	case addr < mem.Tempstart:
		return int(addr) - mem.Localstart
	case addr < mem.Constantstart:
		return int(addr) - mem.Tempstart
	case addr < mem.Scopestart:
		return int(addr) - mem.Constantstart
	default:
		return int(addr) - mem.Scopestart
	}
}

// copyTempoAR copies all the contents of the temp memory to the current activation record
// so that it can be restored later
func (vm *VirtualMachine) copyTempToAR(a *ar.ActivationRecord) error {
	mstemp := vm.mm.memtemp

	a.ResetTemps()

	for i, f := range mstemp.floats {
		a.AddFloatTemp(f, mem.Address(i+mem.FloatOffset+mem.Tempstart))
	}

	for i, in := range mstemp.ints {
		a.AddIntTemp(in, mem.Address(i+mem.IntOffset+mem.Tempstart))
	}

	for i, char := range mstemp.chars {
		a.AddCharTemp(char, mem.Address(i+mem.CharOffset+mem.Tempstart))
	}

	for i, str := range mstemp.strings {
		a.AddStringTemp(str, mem.Address(i+mem.StringOffset+mem.Tempstart))
	}

	for i, b := range mstemp.booleans {
		a.AddBoolTemp(b, mem.Address(i+mem.BoolOffset+mem.Tempstart))
	}

	for i, s := range mstemp.squares {
		a.AddSquareTemp(s, mem.Address((i*objSize)+mem.SquareOffset+mem.Tempstart))
	}

	for i, c := range mstemp.circles {
		a.AddCircleTemp(c, mem.Address((i*objSize)+mem.CircleOffset+mem.Tempstart))
	}

	for i, im := range mstemp.images {
		a.AddImageTemp(im, mem.Address((i*objSize)+mem.ImageOffset+mem.Tempstart))
	}

	for i, t := range mstemp.texts {
		a.AddTextTemp(t, mem.Address((i*objSize)+mem.TextOffset+mem.Tempstart))
	}

	for i, b := range mstemp.backgrounds {
		a.AddBackgroundTemp(b, mem.Address((i*objSize)+mem.BackgroundOffset+mem.Tempstart))
	}

	return nil
}

// copyTempoAR copies all the contents of the temp memory to the current activation record
// so that it can be restored later
func (vm *VirtualMachine) copyLocalToAR(a *ar.ActivationRecord) error {
	mslocal := vm.mm.memlocal

	a.ResetTemps()

	for i, f := range mslocal.floats {
		a.AddFloatLocal(f, mem.Address(i+mem.FloatOffset+mem.Localstart))
	}

	for i, in := range mslocal.ints {
		a.AddIntLocal(in, mem.Address(i+mem.IntOffset+mem.Localstart))
	}

	for i, char := range mslocal.chars {
		a.AddCharLocal(char, mem.Address(i+mem.CharOffset+mem.Localstart))
	}

	for i, str := range mslocal.strings {
		a.AddStringLocal(str, mem.Address(i+mem.StringOffset+mem.Localstart))
	}

	for i, b := range mslocal.booleans {
		a.AddBoolLocal(b, mem.Address(i+mem.BoolOffset+mem.Localstart))
	}

	for i, s := range mslocal.squares {
		a.AddSquareLocal(s, mem.Address((i*objSize)+mem.SquareOffset+mem.Localstart))
	}

	for i, c := range mslocal.circles {
		a.AddCircleLocal(c, mem.Address((i*objSize)+mem.CircleOffset+mem.Localstart))
	}

	for i, im := range mslocal.images {
		a.AddImageLocal(im, mem.Address((i*objSize)+mem.ImageOffset+mem.Localstart))
	}

	for i, t := range mslocal.texts {
		a.AddTextLocal(t, mem.Address((i*objSize)+mem.TextOffset+mem.Localstart))
	}

	for i, b := range mslocal.backgrounds {
		a.AddBackgroundLocal(b, mem.Address((i*objSize)+mem.BackgroundOffset+mem.Localstart))
	}

	return nil
}


func (vm *VirtualMachine) copyTempToMemory(a *ar.ActivationRecord) error {
	for _, p := range a.Floattemps() {
		if err := vm.mm.SetValue(p.Value(), p.Addr()); err != nil {
			return err
		}
	}

	for _, p := range a.Inttemps() {
		if err := vm.mm.SetValue(p.Value(), p.Addr()); err != nil {
			return err
		}
	}

	for _, p := range a.Chartemps() {
		if err := vm.mm.SetValue(p.Value(), p.Addr()); err != nil {
			return err
		}
	}

	for _, p := range a.Stringtemps() {
		if err := vm.mm.SetValue(p.Value(), p.Addr()); err != nil {
			return err
		}
	}

	for _, p := range a.Booltemps() {
		if err := vm.mm.SetValue(p.Value(), p.Addr()); err != nil {
			return err
		}
	}

	for _, p := range a.Squaretemps() {
		if err := vm.mm.SetValue(p.Value(), p.Addr()); err != nil {
			return err
		}
	}

	for _, p := range a.Circletemps() {
		if err := vm.mm.SetValue(p.Value(), p.Addr()); err != nil {
			return err
		}
	}

	for _, p := range a.Imagetemps() {
		if err := vm.mm.SetValue(p.Value(), p.Addr()); err != nil {
			return err
		}
	}

	for _, p := range a.Texttemps() {
		if err := vm.mm.SetValue(p.Value(), p.Addr()); err != nil {
			return err
		}
	}

	for _, p := range a.Backgroundtemps() {
		if err := vm.mm.SetValue(p.Value(), p.Addr()); err != nil {
			return err
		}
	}

	return nil
}

func (vm *VirtualMachine) copyParamsToLocal(a *ar.ActivationRecord) error {
	for _, p := range a.Floatparams() {
		if err := vm.mm.SetValue(p.Value(), p.Addr()); err != nil {
			return err
		}
	}

	for _, p := range a.Intparams() {
		if err := vm.mm.SetValue(p.Value(), p.Addr()); err != nil {
			return err
		}
	}

	for _, p := range a.Charparams() {
		if err := vm.mm.SetValue(p.Value(), p.Addr()); err != nil {
			return err
		}
	}

	for _, p := range a.Stringparams() {
		if err := vm.mm.SetValue(p.Value(), p.Addr()); err != nil {
			return err
		}
	}

	for _, p := range a.Boolparams() {
		if err := vm.mm.SetValue(p.Value(), p.Addr()); err != nil {
			return err
		}
	}

	for _, p := range a.Squareparams() {
		if err := vm.mm.SetValue(p.Value(), p.Addr()); err != nil {
			return err
		}
	}

	for _, p := range a.Circleparams() {
		if err := vm.mm.SetValue(p.Value(), p.Addr()); err != nil {
			return err
		}
	}

	for _, p := range a.Imageparams() {
		if err := vm.mm.SetValue(p.Value(), p.Addr()); err != nil {
			return err
		}
	}

	for _, p := range a.Textparams() {
		if err := vm.mm.SetValue(p.Value(), p.Addr()); err != nil {
			return err
		}
	}

	for _, p := range a.Backgroundparams() {
		if err := vm.mm.SetValue(p.Value(), p.Addr()); err != nil {
			return err
		}
	}

	return nil
}

func getOffsetObject(offset int, addr mem.Address) (int, error) {
	switch {
	case addr < mem.Globalstart: // Error
		return 0, errutil.Newf("Address out of scope")
	case addr < mem.Localstart: // Global
		addr -= mem.Globalstart
	case addr < mem.Tempstart: // Local
		addr -= mem.Localstart
	case addr < mem.Constantstart: // Temp
		addr -= mem.Tempstart
	case addr < mem.Scopestart: // Constant
		addr -= mem.Constantstart
	case addr < mem.Scopestart+5000: // Scope
		addr -= mem.Scopestart
	default: // Error
		return 0, errutil.Newf("Address out of scope")
	}

	// If addr is less than SquareOffset then it is a basic type
	if addr < mem.SquareOffset {
		return offset, nil
	}
	
	return offset * objSize, nil
}