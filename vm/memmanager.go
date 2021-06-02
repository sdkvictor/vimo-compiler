package vm

import (
	"fmt"
	"strings"

	"github.com/sdkvictor/golang-compiler/mem"
	"github.com/sdkvictor/golang-compiler/objects"
	"github.com/sdkvictor/golang-compiler/semantics"
	"github.com/mewkiz/pkg/errutil"
)

var objectSize = semantics.ObjectSize

//MemorySegment represent a single segment in the Memory struct
type MemorySegment struct {
	floats  	[]float64
	ints      	[]int
	chars     	[]rune
	strings     []string
	booleans 	[]bool
	squares 	[]objects.Square
	circles		[]objects.Circle
	images		[]objects.Image
	texts		[]objects.Text
	backgrounds	[]objects.Background
	base     	mem.Address
	name     	string
}

// SetValue takes a value and an address and finds the corresponding position to that address and
// tries to save the given value
func (ms *MemorySegment) SetValue(v interface{}, addr mem.Address) error {
	baseaddr := addr - ms.base

	switch {
	case baseaddr < mem.FloatOffset: // Error
		return errutil.Newf("Address out of scope")
	case baseaddr < mem.CharOffset: // Float
		if f, ok := v.(float64); ok {
			typebaseaddr := int(baseaddr - mem.FloatOffset)
			// If the specified address is bigger than the current size of the array
			// we need to grow the array to that size
			if len(ms.floats) <= typebaseaddr {
				// First we create a new slice with the extra cells we need
				newslice := make([]float64, typebaseaddr-len(ms.floats)+1)
				ms.floats = append(ms.floats, newslice...)
				// Now we set the value to the specified address
				ms.floats[typebaseaddr] = f
			} else {
				ms.floats[typebaseaddr] = f
			}
			return nil
		}
		return errutil.Newf("Cannot set non-float in float address range")
	case baseaddr < mem.BoolOffset: // Character
		if c, ok := v.(rune); ok {
			typebaseaddr := int(baseaddr - mem.CharOffset)
			// If the specified address is bigger than the current size of the array
			// we need to grow the array to that size
			if len(ms.chars) <= typebaseaddr {
				// First we create a new slice with the extra cells we need
				newslice := make([]rune, typebaseaddr-len(ms.chars)+1)
				ms.chars = append(ms.chars, newslice...)
				// Now we set the value to the specified address
				ms.chars[typebaseaddr] = c
			} else {
				ms.chars[typebaseaddr] = c
			}
			return nil
		}
		return errutil.Newf("Cannot set non-char in char address range")
	case baseaddr < mem.IntOffset: // Boolean
		if b, ok := v.(bool); ok {
			typebaseaddr := int(baseaddr - mem.BoolOffset)
			// If the specified address is bigger than the current size of the array
			// we need to grow the array to that size
			if len(ms.booleans) <= typebaseaddr {
				// First we create a new slice with the extra cells we need
				newslice := make([]bool, typebaseaddr-len(ms.booleans)+1)
				ms.booleans = append(ms.booleans, newslice...)
				// Now we set the value to the specified address
				ms.booleans[typebaseaddr] = b
			} else {
				ms.booleans[typebaseaddr] = b
			}
			return nil
		}
		return errutil.Newf("Cannot set non-boolean in boolean address range")
	case baseaddr < mem.StringOffset: // Int
		if i, ok := v.(int); ok {
			typebaseaddr := int(baseaddr - mem.IntOffset)
			// If the specified address is bigger than the current size of the array
			// we need to grow the array to that size
			if len(ms.ints) <= typebaseaddr {
				// First we create a new slice with the extra cells we need
				newslice := make([]int, typebaseaddr-len(ms.ints)+1)
				ms.ints = append(ms.ints, newslice...)
				// Now we set the value to the specified address
				ms.ints[typebaseaddr] = i
			} else {
				ms.ints[typebaseaddr] = i
			}
			return nil
		}
		return errutil.Newf("Cannot set non-int in int address range")
	case baseaddr < mem.SquareOffset: // String
		if s, ok := v.(string); ok {
			typebaseaddr := int(baseaddr - mem.StringOffset)
			// If the specified address is bigger than the current size of the array
			// we need to grow the array to that size
			if len(ms.strings) <= typebaseaddr {
				// First we create a new slice with the extra cells we need
				newslice := make([]string, typebaseaddr-len(ms.strings)+1)
				ms.strings = append(ms.strings, newslice...)
				// Now we set the value to the specified address
				ms.strings[typebaseaddr] = s
			} else {
				ms.strings[typebaseaddr] = s
			}
			return nil
		}
		return errutil.Newf("Cannot set non-string in string address range")
	case baseaddr < mem.CircleOffset: // Square
		typebaseaddr := int(baseaddr - mem.SquareOffset)
		attributeAddress := typebaseaddr % objectSize
		objBase := typebaseaddr - attributeAddress

		if attributeAddress != 0 {
			sq := ms.squares[objBase / objectSize]

			if err := ms.SetAttribute(&sq, v, attributeAddress); err != nil {
				return err
			}
			ms.squares[objBase / objectSize] = sq
			return nil
		} else {
			if s, ok := v.(objects.Square); ok {
				realAddress := typebaseaddr / objectSize
				if len(ms.squares) <= realAddress {
					// First we create a new slice with the extra cells we need
					newslice := make([]objects.Square, realAddress-len(ms.squares)+1)
					ms.squares = append(ms.squares, newslice...)
					// Now we set the value to the specified address
					ms.squares[realAddress] = s
				} else {
					ms.squares[realAddress] = s
				}
				return nil
			}
		}

		return errutil.Newf("Cannot set non-square in Square address range")
	case baseaddr < mem.ImageOffset: // Circle
		typebaseaddr := int(baseaddr - mem.CircleOffset)
		attributeAddress := typebaseaddr % objectSize
		objBase := typebaseaddr - attributeAddress

		if attributeAddress != 0 {
			sq := ms.circles[objBase / objectSize]

			if err := ms.SetAttribute(&sq, v, attributeAddress); err != nil {
				return err
			}
			ms.circles[objBase / objectSize] = sq
			return nil
		} else {
			if s, ok := v.(objects.Circle); ok {
				realAddress := typebaseaddr / objectSize
				if len(ms.circles) <= realAddress {
					// First we create a new slice with the extra cells we need
					newslice := make([]objects.Circle, realAddress-len(ms.circles)+1)
					ms.circles = append(ms.circles, newslice...)
					// Now we set the value to the specified address
					ms.circles[realAddress] = s
				} else {
					ms.circles[realAddress] = s
				}
				return nil
			}
		}
		return errutil.Newf("Cannot set non-square in Circle address range")
	case baseaddr < mem.TextOffset: // Image
		typebaseaddr := int(baseaddr - mem.ImageOffset)
		attributeAddress := typebaseaddr % objectSize
		objBase := typebaseaddr - attributeAddress

		if attributeAddress != 0 {
			sq := ms.images[objBase / objectSize]

			if err := ms.SetAttribute(&sq, v, attributeAddress); err != nil {
				return err
			}
			ms.images[objBase / objectSize] = sq
			return nil
		} else {
			if s, ok := v.(objects.Image); ok {
				realAddress := typebaseaddr / objectSize
				if len(ms.images) <= realAddress {
					// First we create a new slice with the extra cells we need
					newslice := make([]objects.Image, realAddress-len(ms.images)+1)
					ms.images = append(ms.images, newslice...)
					// Now we set the value to the specified address
					ms.images[realAddress] = s
				} else {
					ms.images[realAddress] = s
				}
				return nil
			}
		}
		return errutil.Newf("Cannot set non-square in Image address range")
	case baseaddr < mem.BackgroundOffset: // Text
		typebaseaddr := int(baseaddr - mem.TextOffset)
		attributeAddress := typebaseaddr % objectSize
		objBase := typebaseaddr - attributeAddress

		if attributeAddress != 0 {
			sq := ms.texts[objBase / objectSize]

			if err := ms.SetAttribute(&sq, v, attributeAddress); err != nil {
				return err
			}
			ms.texts[objBase / objectSize] = sq
			return nil
		} else {
			if s, ok := v.(objects.Text); ok {
				realAddress := typebaseaddr / objectSize
				if len(ms.texts) <= realAddress {
					// First we create a new slice with the extra cells we need
					newslice := make([]objects.Text, realAddress-len(ms.texts)+1)
					ms.texts = append(ms.texts, newslice...)
					// Now we set the value to the specified address
					ms.texts[realAddress] = s
				} else {
					ms.texts[realAddress] = s
				}
				return nil
			}
		}
		return errutil.Newf("Cannot set non-square in Image address range")
	case baseaddr < mem.BackgroundOffset+1000: // Backgrounds
		typebaseaddr := int(baseaddr - mem.TextOffset)
		attributeAddress := typebaseaddr % objectSize
		objBase := typebaseaddr - attributeAddress

		if attributeAddress != 0 {
			sq := ms.backgrounds[objBase / objectSize]

			if err := ms.SetAttribute(&sq, v, attributeAddress); err != nil {
				return err
			}
			ms.backgrounds[objBase / objectSize] = sq
			return nil
		} else {
			if s, ok := v.(objects.Background); ok {
				realAddress := typebaseaddr / objectSize
				if len(ms.backgrounds) <= realAddress {
					// First we create a new slice with the extra cells we need
					newslice := make([]objects.Background, realAddress-len(ms.backgrounds)+1)
					ms.backgrounds = append(ms.backgrounds, newslice...)
					// Now we set the value to the specified address
					ms.backgrounds[realAddress] = s
				} else {
					ms.backgrounds[realAddress] = s
				}
				return nil
			}
		}
		return errutil.Newf("Cannot set non-square in Image address range")
	default: // Error
		return errutil.Newf("Address out of scope")
	}

}

func (ms *MemorySegment) SetAttribute(obj objects.Object, v interface{}, attAddr int) error {
	attribute := semantics.ObjectAttributes[attAddr - 1]

	//fmt.Printf("SETTING %s to %v\n", attribute, v)

	switch attribute {
	case "height":
		vf, ok := v.(float64)
		if !ok {
			return errutil.Newf("Cannot cast attribute to corresponding type %T", v)
		}
		obj.SetHeight(vf)
	case "width":
		vf, ok := v.(float64)
		if !ok {
			return errutil.Newf("Cannot cast attribute to corresponding type")
		}
		obj.SetWidth(vf)
	case "x":
		vf, ok := v.(float64)
		if !ok {
			return errutil.Newf("Cannot cast attribute to corresponding type")
		}
		obj.SetX(vf)
	case "y":
		vf, ok := v.(float64)
		if !ok {
			return errutil.Newf("Cannot cast attribute to corresponding type")
		}
		obj.SetY(vf)
	case "size":
		vf, ok := v.(float64)
		if !ok {
			return errutil.Newf("Cannot cast attribute to corresponding type")
		}
		obj.SetSize(vf)
	case "color":
		vf, ok := v.(string)
		if !ok {
			return errutil.Newf("Cannot cast attribute to corresponding type")
		}
		obj.SetColor(vf)
	case "message":
		vf, ok := v.(string)
		if !ok {
			return errutil.Newf("Cannot cast attribute to corresponding type")
		}
		obj.SetMessage(vf)
	case "image":
		vf, ok := v.(string)
		if !ok {
			return errutil.Newf("Cannot cast attribute to corresponding type")
		}
		obj.SetImage(vf)
	default:
		return errutil.Newf("Invalid object attribute")
	}

	return nil
}

// GetValue takes and address and tries to retreive the corresponding value in that position
func (ms *MemorySegment) GetValue(addr mem.Address) (interface{}, error) { // AQUI
	baseaddr := addr - ms.base

	switch {
	case baseaddr < mem.FloatOffset: // Error
		return nil, errutil.Newf("Address out of scope: %d", baseaddr)
	case baseaddr < mem.CharOffset: // Float
		typebaseaddr := int(baseaddr - mem.FloatOffset)
		if len(ms.floats) <= typebaseaddr {
			return nil, errutil.Newf("%s: Referencing address %d out of scope of %d", ms.name, typebaseaddr, len(ms.floats))
		}
		return ms.floats[typebaseaddr], nil
	case baseaddr < mem.BoolOffset: // Character
		typebaseaddr := int(baseaddr - mem.CharOffset)
		if len(ms.chars) <= typebaseaddr {
			return nil, errutil.Newf("Referencing address out of scope")
		}
		return ms.chars[typebaseaddr], nil
	case baseaddr < mem.IntOffset: // Boolean
		typebaseaddr := int(baseaddr - mem.BoolOffset)
		if len(ms.booleans) <= typebaseaddr {
			return nil, errutil.Newf("Referencing address out of scope")
		}
		return ms.booleans[typebaseaddr], nil
	case baseaddr < mem.StringOffset: // Int
		typebaseaddr := int(baseaddr - mem.IntOffset)
		if len(ms.ints) <= typebaseaddr {
			return nil, errutil.Newf("Referencing address out of scope")
		}
		return ms.ints[typebaseaddr], nil
	case baseaddr < mem.SquareOffset: // String
		typebaseaddr := int(baseaddr - mem.StringOffset)
		if len(ms.strings) <= typebaseaddr {
			return nil, errutil.Newf("Referencing address out of scope")
		}
		return ms.strings[typebaseaddr], nil
	case baseaddr < mem.CircleOffset: // Square
		typebaseaddr := int(baseaddr - mem.SquareOffset)
		attributeAddress := typebaseaddr % objectSize
		objBase := typebaseaddr - attributeAddress

		if len(ms.squares) <= objBase / objectSize {
			return nil, errutil.Newf("Referencing address out of scope")
		}

		if attributeAddress != 0 {
			s := ms.squares[objBase / objectSize]
			return GetAttribute(&s, attributeAddress)
		} else {
			return ms.squares[objBase / objectSize], nil
		}

	case baseaddr < mem.ImageOffset: // Circle
		typebaseaddr := int(baseaddr - mem.CircleOffset)
		attributeAddress := typebaseaddr % objectSize
		objBase := typebaseaddr - attributeAddress

		if len(ms.circles) <= objBase / objectSize {
			return nil, errutil.Newf("Referencing address out of scope")
		}

		if attributeAddress != 0 {
			s := ms.circles[objBase / objectSize]
			return GetAttribute(&s, attributeAddress)
		} else {
			return ms.circles[objBase / objectSize], nil
		}
	case baseaddr < mem.TextOffset: // Image
		typebaseaddr := int(baseaddr - mem.ImageOffset)
		attributeAddress := typebaseaddr % objectSize
		objBase := typebaseaddr - attributeAddress

		if len(ms.images) <= objBase / objectSize {
			return nil, errutil.Newf("Referencing address out of scope")
		}

		if attributeAddress != 0 {
			s := ms.images[objBase / objectSize]
			return GetAttribute(&s, attributeAddress)
		} else {
			return ms.images[objBase / objectSize], nil
		}
	case baseaddr < mem.BackgroundOffset: // Text
		typebaseaddr := int(baseaddr - mem.TextOffset)
		attributeAddress := typebaseaddr % objectSize
		objBase := typebaseaddr - attributeAddress

		if len(ms.texts) <= objBase / objectSize {
			return nil, errutil.Newf("Referencing address out of scope")
		}

		if attributeAddress != 0 {
			s := ms.texts[objBase / objectSize]
			return GetAttribute(&s, attributeAddress)
		} else {
			return ms.texts[objBase / objectSize], nil
		}
	case baseaddr < mem.BackgroundOffset+1000: // Background
		typebaseaddr := int(baseaddr - mem.BackgroundOffset)
		attributeAddress := typebaseaddr % objectSize
		objBase := typebaseaddr - attributeAddress

		if len(ms.backgrounds) <= objBase / objectSize {
			return nil, errutil.Newf("Referencing address out of scope")
		}

		if attributeAddress != 0 {
			s := ms.backgrounds[objBase / objectSize]
			return GetAttribute(&s, attributeAddress)
		} else {
			return ms.backgrounds[objBase / objectSize], nil
		}
	default: // Error
		return nil, errutil.Newf("Address out of scope")
	}

}

func GetAttribute(obj objects.Object, attAddr int) (interface{}, error) {
	attribute := semantics.ObjectAttributes[attAddr - 1]

	switch attribute {
	case "height":
		return obj.Height(), nil
	case "width":
		return obj.Width(), nil
	case "x":
		return obj.X(), nil
	case "y":
		return obj.Y(), nil
	case "size":
		return obj.Size(), nil
	case "color":
		return obj.Color(), nil
	case "message":
		return obj.Message(), nil
	case "image":
		return obj.Image(), nil
	default:
		return 0, errutil.Newf("Invalid object attribute")
	}
}

func (ms *MemorySegment) String() string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("  %s:\n", ms.name))

	builder.WriteString("    Floats:\n")
	for i, v := range ms.floats {
		builder.WriteString(fmt.Sprintf("      %d: %f\n", i, v))
	}
	builder.WriteString("    Chars:\n")
	for i, v := range ms.chars {
		builder.WriteString(fmt.Sprintf("      %d: %c\n", i, v))
	}
	builder.WriteString("    Booleans:\n")
	for i, v := range ms.booleans {
		builder.WriteString(fmt.Sprintf("      %d: %t\n", i, v))
	}
	builder.WriteString("    Ints:\n")
	for i, v := range ms.ints {
		builder.WriteString(fmt.Sprintf("      %d: %d\n", i, v))
	}
	builder.WriteString("    Strings:\n")
	for i, v := range ms.strings {
		builder.WriteString(fmt.Sprintf("      %d: %d\n", i, v))
	}
	builder.WriteString("    Squares:\n")
	for i, v := range ms.squares {
		builder.WriteString(fmt.Sprintf("      %d: %d\n", i, v))
	}
	builder.WriteString("    Circles:\n")
	for i, v := range ms.circles {
		builder.WriteString(fmt.Sprintf("      %d: %d\n", i, v))
	}
	builder.WriteString("    Images:\n")
	for i, v := range ms.images {
		builder.WriteString(fmt.Sprintf("      %d: %d\n", i, v))
	}
	builder.WriteString("    Texts:\n")
	for i, v := range ms.texts {
		builder.WriteString(fmt.Sprintf("      %d: %d\n", i, v))
	}
	builder.WriteString("    Backgrounds:\n")
	for i, v := range ms.backgrounds {
		builder.WriteString(fmt.Sprintf("      %d: %s\n", i, v))
	}
	return builder.String()
}

func NewMemorySegment(base int, name string) *MemorySegment {
	return &MemorySegment{
		make([]float64, 0),					// floats
		make([]int, 0),						// ints
		make([]rune, 0), 					// chars
		make([]string, 0),					// strings
		make([]bool, 0),					// booleans
		make([]objects.Square, 0),					// squares
		make([]objects.Circle, 0),					// circles
		make([]objects.Image, 0),					// images
		make([]objects.Text, 0),					// texts
		make([]objects.Background, 0),				// backgrounds
		mem.Address(base),
		name,
	}
}

// Memory represents the virtual memory for the virtual machine
// it contains one MemorySegment for each segment
type Memory struct {
	memglobal   *MemorySegment
	memlocal    *MemorySegment
	memtemp     *MemorySegment
	memconstant *MemorySegment
	memscope    *MemorySegment
}

func NewMemory() *Memory {
	return &Memory{
		NewMemorySegment(mem.Globalstart, "Global"),
		NewMemorySegment(mem.Localstart, "Local"),
		NewMemorySegment(mem.Tempstart, "Temp"),
		NewMemorySegment(mem.Constantstart, "Constant"),
		NewMemorySegment(mem.Scopestart, "Scope"),
	}
}

// Get value takes an address and tries to retrieve the saved value by consulting its memory segments
func (m *Memory) GetValue(addr mem.Address) (interface{}, error) {
	switch {
	case addr < mem.Globalstart: // Error
		return false, errutil.Newf("Address out of scope")
	case addr < mem.Localstart: // Global
		v, err := m.memglobal.GetValue(addr)
		if err != nil {
			return nil, err
		}
		return v, nil
	case addr < mem.Tempstart: // Local
		v, err := m.memlocal.GetValue(addr)
		if err != nil {
			return nil, err
		}
		return v, nil
	case addr < mem.Constantstart: // Temp
		v, err := m.memtemp.GetValue(addr)
		if err != nil {
			return nil, err
		}
		return v, nil
	case addr < mem.Scopestart: // Constant
		v, err := m.memconstant.GetValue(addr)
		if err != nil {
			return nil, err
		}
		return v, nil
	case addr < mem.Scopestart+5000: // Scope
		v, err := m.memscope.GetValue(addr)
		if err != nil {
			return nil, err
		}
		return v, nil
	default: // Error
		return nil, errutil.Newf("Address out of scope")
	}
}

// SetValue takes a value and an address and then consults its segments to
// try to save the value
func (m *Memory) SetValue(v interface{}, addr mem.Address) error {
	switch {
	case addr < mem.Globalstart: // Error
		return errutil.Newf("Address out of scope")
	case addr < mem.Localstart: // Global
		if err := m.memglobal.SetValue(v, addr); err != nil {
			return err
		}
		return nil
	case addr < mem.Tempstart: // Local
		if err := m.memlocal.SetValue(v, addr); err != nil {
			return err
		}
		return nil
	case addr < mem.Constantstart: // Temp
		if err := m.memtemp.SetValue(v, addr); err != nil {
			return err
		}
		return nil
	case addr < mem.Scopestart: // Constant
		if err := m.memconstant.SetValue(v, addr); err != nil {
			return err
		}
		return nil
	case addr < mem.Scopestart+5000: // Scope
		if err := m.memscope.SetValue(v, addr); err != nil {
			return err
		}
		return nil
	default: // Error
		return errutil.Newf("Address out of scope")
	}
}

func (m *Memory) String() string {
	var builder strings.Builder

	builder.WriteString("Memory:\n")
	builder.WriteString(m.memglobal.String())
	builder.WriteString(m.memlocal.String())
	builder.WriteString(m.memtemp.String())
	builder.WriteString(m.memconstant.String())
	builder.WriteString(m.memscope.String())

	return builder.String()
}