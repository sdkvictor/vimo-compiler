package mem

import (
	"fmt"

	"github.com/sdkvictor/golang-compiler/types"
	"github.com/mewkiz/pkg/errutil"
)

//Type of the address of the memory is an int
type Address int

//Constants that mark the start of the context
const Globalstart = 0
const Localstart = 5000
const Tempstart = 10000
const Constantstart = 15000
const Scopestart = 20000

//Constants that set the offset of the segment
const NumOffset = 0
const CharOffset = 1000
const BoolOffset = 2000
const FunctionOffset = 3000
const ListOffset = 4000

//Constant that defines de size of the segment
const segmentsize = 1000