package ar

import (
	"github.com/sdkvictor/golang-compiler/mem"
	"github.com/sdkvictor/golang-compiler/objects"
)

type FloatParam struct {
	value float64
	addr  mem.Address
}

func (fp FloatParam) Value() float64 {
	return fp.value
}

func (fp FloatParam) Addr() mem.Address {
	return fp.addr
}

type IntParam struct {
	value int
	addr  mem.Address
}

func (ip IntParam) Value() int {
	return ip.value
}

func (ip IntParam) Addr() mem.Address {
	return ip.addr
}

type CharParam struct {
	value rune
	addr  mem.Address
}

func (cp CharParam) Value() rune {
	return cp.value
}

func (cp CharParam) Addr() mem.Address {
	return cp.addr
}

type BoolParam struct {
	value bool
	addr  mem.Address
}

func (bp BoolParam) Value() bool {
	return bp.value
}

func (bp BoolParam) Addr() mem.Address {
	return bp.addr
}

type StringParam struct {
	value string
	addr  mem.Address
}

func (sp StringParam) Value() string {
	return sp.value
}

func (sp StringParam) Addr() mem.Address {
	return sp.addr
}

type SquareParam struct {
	value objects.Square
	addr  mem.Address
}

func (sp SquareParam) Value() objects.Square {
	return sp.value
}

func (sp SquareParam) Addr() mem.Address {
	return sp.addr
}

type CircleParam struct {
	value objects.Circle
	addr  mem.Address
}

func (cp CircleParam) Value() objects.Circle {
	return cp.value
}

func (cp CircleParam) Addr() mem.Address {
	return cp.addr
}

type ImageParam struct {
	value objects.Image
	addr  mem.Address
}

func (ip ImageParam) Value() objects.Image {
	return ip.value
}

func (ip ImageParam) Addr() mem.Address {
	return ip.addr
}

type TextParam struct {
	value objects.Text
	addr  mem.Address
}

func (tp TextParam) Value() objects.Text {
	return tp.value
}

func (tp TextParam) Addr() mem.Address {
	return tp.addr
}

type BackgroundParam struct {
	value objects.Background
	addr  mem.Address
}

func (bp BackgroundParam) Value() objects.Background {
	return bp.value
}

func (bp BackgroundParam) Addr() mem.Address {
	return bp.addr
}

type ActivationRecord struct {
	retip      			int
	floatparams  		[]FloatParam
	intparams  			[]IntParam
	charparams 			[]CharParam
	boolparams 			[]BoolParam
	stringparams 		[]StringParam
	squareparams		[]SquareParam
	circleparams		[]CircleParam
	imageparams			[]ImageParam
	textparams			[]TextParam
	backgroundparams	[]BackgroundParam
	floattemps   		[]FloatParam
	inttemps   			[]IntParam
	chartemps  			[]CharParam
	booltemps  			[]BoolParam
	stringtemps  		[]StringParam
	squaretemps			[]SquareParam
	circletemps			[]CircleParam
	imagetemps			[]ImageParam
	texttemps			[]TextParam
	backgroundtemps		[]BackgroundParam
	floatcount   		int
	intcount			int
	charcount  			int
	boolcount  			int
	stringcount  		int
	squarecount			int
	circlecount			int
	imagecount			int
	textcount			int
	backgroundcount		int
}

func (a *ActivationRecord) SetRetIp(ip int) {
	a.retip = ip
}

func (a *ActivationRecord) AddFloatParam(f float64) {
	addr := mem.Address(a.floatcount + mem.Localstart + mem.FloatOffset)
	a.floatparams = append(a.floatparams, FloatParam{f, addr})
	a.floatcount++
}

func (a *ActivationRecord) AddIntParam(i int) {
	addr := mem.Address(a.intcount + mem.Localstart + mem.IntOffset)
	a.intparams = append(a.intparams, IntParam{i, addr})
	a.intcount++
}

func (a *ActivationRecord) AddCharParam(char rune) {
	addr := mem.Address(a.charcount + mem.Localstart + mem.CharOffset)
	a.charparams = append(a.charparams, CharParam{char, addr})
	a.charcount++
}

func (a *ActivationRecord) AddBoolParam(b bool) {
	addr := mem.Address(a.boolcount + mem.Localstart + mem.BoolOffset)
	a.boolparams = append(a.boolparams, BoolParam{b, addr})
	a.boolcount++
}

func (a *ActivationRecord) AddStringParam(s string) {
	addr := mem.Address(a.stringcount + mem.Localstart + mem.StringOffset)
	a.stringparams = append(a.stringparams, StringParam{s, addr})
	a.stringcount++
}

func (a *ActivationRecord) AddSquareParam(s objects.Square) {
	addr := mem.Address(a.squarecount + mem.Localstart + mem.SquareOffset)
	a.squareparams = append(a.squareparams, SquareParam{s, addr})
	a.squarecount++
}

func (a *ActivationRecord) AddCircleParam(c objects.Circle) {
	addr := mem.Address(a.circlecount + mem.Localstart + mem.CircleOffset)
	a.circleparams = append(a.circleparams, CircleParam{c, addr})
	a.circlecount++
}

func (a *ActivationRecord) AddImageParam(i objects.Image) {
	addr := mem.Address(a.imagecount + mem.Localstart + mem.ImageOffset)
	a.imageparams = append(a.imageparams, ImageParam{i, addr})
	a.imagecount++
}

func (a *ActivationRecord) AddTextParam(t objects.Text) {
	addr := mem.Address(a.textcount + mem.Localstart + mem.TextOffset)
	a.textparams = append(a.textparams, TextParam{t, addr})
	a.textcount++
}

func (a *ActivationRecord) AddBackgroundParam(b objects.Background) {
	addr := mem.Address(a.backgroundcount + mem.Localstart + mem.BackgroundOffset)
	a.backgroundparams = append(a.backgroundparams, BackgroundParam{b, addr})
	a.backgroundcount++
}

func (a *ActivationRecord) AddFloatLocal(f float64, addr mem.Address) {
	a.floatparams = append(a.floatparams, FloatParam{f, addr})
	a.floatcount++
}

func (a *ActivationRecord) AddIntLocal(i int, addr mem.Address) {
	a.intparams = append(a.intparams, IntParam{i, addr})
	a.intcount++
}

func (a *ActivationRecord) AddCharLocal(char rune, addr mem.Address) {
	a.charparams = append(a.charparams, CharParam{char, addr})
	a.charcount++
}

func (a *ActivationRecord) AddBoolLocal(b bool, addr mem.Address) {
	a.boolparams = append(a.boolparams, BoolParam{b, addr})
	a.boolcount++
}

func (a *ActivationRecord) AddStringLocal(s string, addr mem.Address) {
	a.stringparams = append(a.stringparams, StringParam{s, addr})
	a.stringcount++
}

func (a *ActivationRecord) AddSquareLocal(s objects.Square, addr mem.Address) {
	a.squareparams = append(a.squareparams, SquareParam{s, addr})
	a.squarecount++
}

func (a *ActivationRecord) AddCircleLocal(c objects.Circle, addr mem.Address) {
	a.circleparams = append(a.circleparams, CircleParam{c, addr})
	a.circlecount++
}

func (a *ActivationRecord) AddImageLocal(i objects.Image, addr mem.Address) {
	a.imageparams = append(a.imageparams, ImageParam{i, addr})
	a.imagecount++
}

func (a *ActivationRecord) AddTextLocal(t objects.Text, addr mem.Address) {
	a.textparams = append(a.textparams, TextParam{t, addr})
	a.textcount++
}

func (a *ActivationRecord) AddBackgroundLocal(b objects.Background, addr mem.Address) {
	a.backgroundparams = append(a.backgroundparams, BackgroundParam{b, addr})
	a.backgroundcount++
}

func (a *ActivationRecord) AddFloatTemp(f float64, addr mem.Address) {
	a.floattemps = append(a.floattemps, FloatParam{f, addr})
}

func (a *ActivationRecord) AddIntTemp(i int, addr mem.Address) {
	a.inttemps = append(a.inttemps, IntParam{i, addr})
}

func (a *ActivationRecord) AddCharTemp(char rune, addr mem.Address) {
	a.chartemps = append(a.chartemps, CharParam{char, addr})
}

func (a *ActivationRecord) AddBoolTemp(b bool, addr mem.Address) {
	a.booltemps = append(a.booltemps, BoolParam{b, addr})
}

func (a *ActivationRecord) AddStringTemp(s string, addr mem.Address) {
	a.stringtemps = append(a.stringtemps, StringParam{s, addr})
}

func (a *ActivationRecord) AddSquareTemp(s objects.Square, addr mem.Address) {
	a.squaretemps = append(a.squaretemps, SquareParam{s, addr})
}

func (a *ActivationRecord) AddCircleTemp(c objects.Circle, addr mem.Address) {
	a.circletemps = append(a.circletemps, CircleParam{c, addr})
}

func (a *ActivationRecord) AddImageTemp(i objects.Image, addr mem.Address) {
	a.imagetemps = append(a.imagetemps, ImageParam{i, addr})
}

func (a *ActivationRecord) AddTextTemp(t objects.Text, addr mem.Address) {
	a.texttemps = append(a.texttemps, TextParam{t, addr})
}

func (a *ActivationRecord) AddBackgroundTemp(b objects.Background, addr mem.Address) {
	a.backgroundtemps = append(a.backgroundtemps, BackgroundParam{b, addr})
}

func (a *ActivationRecord) ResetTemps() {
	a.floattemps = make([]FloatParam, 0)
	a.inttemps = make([]IntParam, 0)
	a.chartemps = make([]CharParam, 0)
	a.booltemps = make([]BoolParam, 0)
	a.stringtemps = make([]StringParam, 0)
	a.squaretemps = make([]SquareParam, 0)
	a.circletemps = make([]CircleParam, 0)
	a.imagetemps = make([]ImageParam, 0)
	a.texttemps = make([]TextParam, 0)
	a.backgroundtemps = make([]BackgroundParam, 0)
}

func (a *ActivationRecord) ResetParams() {
	a.floatparams = make([]FloatParam, 0)
	a.intparams = make([]IntParam, 0)
	a.charparams = make([]CharParam, 0)
	a.boolparams = make([]BoolParam, 0)
	a.stringparams = make([]StringParam, 0)
	a.squareparams = make([]SquareParam, 0)
	a.circleparams = make([]CircleParam, 0)
	a.imageparams = make([]ImageParam, 0)
	a.textparams = make([]TextParam, 0)
	a.backgroundparams = make([]BackgroundParam, 0)
}

func NewActivationRecord() *ActivationRecord {
	return &ActivationRecord{
		0,									// tip
		make([]FloatParam, 0),				// floatparams
		make([]IntParam, 0),				// intparams
		make([]CharParam, 0),				// charparams
		make([]BoolParam, 0),				// boolparams
		make([]StringParam, 0),				// stringparams
		make([]SquareParam, 0),				// squareparams
		make([]CircleParam, 0),				// circleparams
		make([]ImageParam, 0),				// imageparams
		make([]TextParam, 0),				// textparams
		make([]BackgroundParam, 0),			// backgroundparams
		make([]FloatParam, 0),				// floattemps
		make([]IntParam, 0),				// inttemps
		make([]CharParam, 0),				// chartemps
		make([]BoolParam, 0),				// booltemps
		make([]StringParam, 0),				// stringtemps
		make([]SquareParam, 0),				// squaretemps
		make([]CircleParam, 0),				// circletemps
		make([]ImageParam, 0),				// imagetemps
		make([]TextParam, 0),				// textparams
		make([]BackgroundParam, 0),			// backgroundtemps
		0,									// floatcount
		0,									// intcount
		0,									// charcount
		0,									// boolcount
		0,									// stringcount
		0,									// squarecount
		0,									// circlecount
		0,									// imagecount
		0,									// textcount
		0,									// backgroundcount
	}
}

// stringcount getter and setter
func (a *ActivationRecord) Stringcount() int {
	return a.stringcount
}

func (a *ActivationRecord) SetStringcount(stringcount int) {
	a.stringcount = stringcount
}

// boolcount getter and setter
func (a *ActivationRecord) Boolcount() int {
	return a.boolcount
}

func (a *ActivationRecord) SetBoolcount(boolcount int) {
	a.boolcount = boolcount
}

// charcount getter and setter
func (a *ActivationRecord) Charcount() int {
	return a.charcount
}

func (a *ActivationRecord) SetCharcount(charcount int) {
	a.charcount = charcount
}

// intcount getter and setter
func (a *ActivationRecord) Intcount() int {
	return a.intcount
}

func (a *ActivationRecord) SetIntcount(intcount int) {
	a.intcount = intcount
}

// floatcount getter and setter
func (a *ActivationRecord) Floatcount() int {
	return a.floatcount
}

func (a *ActivationRecord) SetFloatcount(floatcount int) {
	a.floatcount = floatcount
}

// squarecount getter and setter
func (a *ActivationRecord) Squarecount() int {
	return a.squarecount
}

func (a *ActivationRecord) SetSquarecount(squarecount int) {
	a.squarecount = squarecount
}

// circlecount getter and setter
func (a *ActivationRecord) Circlecount() int {
	return a.circlecount
}

func (a *ActivationRecord) SetCirclecount(circlecount int) {
	a.circlecount = circlecount
}

// imagecount getter and setter
func (a *ActivationRecord) Imagecount() int {
	return a.imagecount
}

func (a *ActivationRecord) SetImagecount(imagecount int) {
	a.imagecount = imagecount
}

// textcount getter and setter
func (a *ActivationRecord) Textcount() int {
	return a.textcount
}

func (a *ActivationRecord) SetTextcount(textcount int) {
	a.textcount = textcount
}

// backgroundcount getter and setter
func (a *ActivationRecord) Backgroundcount() int {
	return a.backgroundcount
}

func (a *ActivationRecord) SetBackgroundcount(backgroundcount int) {
	a.backgroundcount = backgroundcount
}

// stringtemps getter and setter
func (a *ActivationRecord) Stringtemps() []StringParam {
	return a.stringtemps
}

func (a *ActivationRecord) SetStringtemps(stringtemps []StringParam) {
	a.stringtemps = stringtemps
}

// booltemps getter and setter
func (a *ActivationRecord) Booltemps() []BoolParam {
	return a.booltemps
}

func (a *ActivationRecord) SetBooltemps(booltemps []BoolParam) {
	a.booltemps = booltemps
}

// chartemps getter and setter
func (a *ActivationRecord) Chartemps() []CharParam {
	return a.chartemps
}

func (a *ActivationRecord) SetChartemps(chartemps []CharParam) {
	a.chartemps = chartemps
}

// inttemps getter and setter
func (a *ActivationRecord) Inttemps() []IntParam {
	return a.inttemps
}

func (a *ActivationRecord) SetInttemps(inttemps []IntParam) {
	a.inttemps = inttemps
}

// floattemps getter and setter
func (a *ActivationRecord) Floattemps() []FloatParam {
	return a.floattemps
}

func (a *ActivationRecord) SetFloattemps(floattemps []FloatParam) {
	a.floattemps = floattemps
}

// squaretemps getter and setter
func (a *ActivationRecord) Squaretemps() []SquareParam {
	return a.squaretemps
}

func (a *ActivationRecord) SetSquaretemps(squaretemps []SquareParam) {
	a.squaretemps = squaretemps
}

// circletemps getter and setter
func (a *ActivationRecord) Circletemps() []CircleParam {
	return a.circletemps
}

func (a *ActivationRecord) SetCircletemps(circletemps []CircleParam) {
	a.circletemps = circletemps
}

// imagetemps getter and setter
func (a *ActivationRecord) Imagetemps() []ImageParam {
	return a.imagetemps
}

func (a *ActivationRecord) SetImagetemps(imagetemps []ImageParam) {
	a.imagetemps = imagetemps
}

// texttemps getter and setter
func (a *ActivationRecord) Texttemps() []TextParam {
	return a.texttemps
}

func (a *ActivationRecord) SetTexttemps(texttemps []TextParam) {
	a.texttemps = texttemps
}

// backgroundtemps getter and setter
func (a *ActivationRecord) Backgroundtemps() []BackgroundParam {
	return a.backgroundtemps
}

func (a *ActivationRecord) SetBackgroundtemps(backgroundtemps []BackgroundParam) {
	a.backgroundtemps = backgroundtemps
}

// stringparams getter and setter
func (a *ActivationRecord) Stringparams() []StringParam {
	return a.stringparams
}

func (a *ActivationRecord) SetStringparams(stringparams []StringParam) {
	a.stringparams = stringparams
}

// boolparams getter and setter
func (a *ActivationRecord) Boolparams() []BoolParam {
	return a.boolparams
}

func (a *ActivationRecord) SetBoolparams(boolparams []BoolParam) {
	a.boolparams = boolparams
}

// charparams getter and setter
func (a *ActivationRecord) Charparams() []CharParam {
	return a.charparams
}

func (a *ActivationRecord) SetCharparams(charparams []CharParam) {
	a.charparams = charparams
}

// intparams getter and setter
func (a *ActivationRecord) Intparams() []IntParam {
	return a.intparams
}

func (a *ActivationRecord) SetIntparams(intparams []IntParam) {
	a.intparams = intparams
}

// floatparams getter and setter
func (a *ActivationRecord) Floatparams() []FloatParam {
	return a.floatparams
}

func (a *ActivationRecord) SetFloatparams(floatparams []FloatParam) {
	a.floatparams = floatparams
}

// squareparams getter and setter
func (a *ActivationRecord) Squareparams() []SquareParam {
	return a.squareparams
}

func (a *ActivationRecord) SetSquareparams(squareparams []SquareParam) {
	a.squareparams = squareparams
}

// circleparams getter and setter
func (a *ActivationRecord) Circleparams() []CircleParam {
	return a.circleparams
}

func (a *ActivationRecord) SetCircleparams(circleparams []CircleParam) {
	a.circleparams = circleparams
}

// imageparams getter and setter
func (a *ActivationRecord) Imageparams() []ImageParam {
	return a.imageparams
}

func (a *ActivationRecord) SetImageparams(imageparams []ImageParam) {
	a.imageparams = imageparams
}

// textparams getter and setter
func (a *ActivationRecord) Textparams() []TextParam {
	return a.textparams
}

func (a *ActivationRecord) SetTextparams(textparams []TextParam) {
	a.textparams = textparams
}

// backgroundparams getter and setter
func (a *ActivationRecord) Backgroundparams() []BackgroundParam {
	return a.backgroundparams
}

func (a *ActivationRecord) SetBackgroundparams(backgroundparams []BackgroundParam) {
	a.backgroundparams = backgroundparams
}

// retip getter and setter
func (a *ActivationRecord) Retip() int {
	return a.retip
}

func (a *ActivationRecord) SetRetip(retip int) {
	a.retip = retip
}