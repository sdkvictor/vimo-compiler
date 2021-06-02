package objects

import "fmt"

type Object interface {
	X() float64
	Y() float64
	Width() float64
	Height() float64
	Color() string
	Image() string
	Size() float64
	Message() string
	SetX(float64)
	SetY(float64) 
	SetWidth(float64) 
	SetHeight(float64) 
	SetColor(string) 
	SetImage(string) 
	SetSize(float64) 
	SetMessage(string)
	String() string
}


type Square struct {
	x float64
	y float64
	width float64
	height float64
	color string
}

func (o *Square) X() float64 {
	return o.x
}

func (o *Square) Y() float64 {
	return o.y
}

func (o *Square) Width() float64 {
	return o.width
}

func (o *Square) Height() float64 {
	return o.height
}

func (o *Square) Color() string {
	return o.color
}

func (o *Square) Image() string {
	return ""
}

func (o *Square) Size() float64 {
	return 0.0
}

func (o *Square) Message() string {
	return ""
}

func (o *Square) SetX(x float64) {
	o.x = x
}

func (o *Square) SetY(x float64) {
	o.y = x
}

func (o *Square) SetWidth(x float64) {
	o.width = x
}

func (o *Square) SetHeight(x float64) {
	o.height = x
}

func (o *Square) SetColor(x string) {
	o.color = x
}

func (o *Square) SetImage(x string) {
	
}

func (o *Square) SetSize(x float64) {
	
}

func (o *Square) SetMessage(x string) {
	
}

func (o *Square) String() string {
	return fmt.Sprintf("Square = x: %f, y: %f, width: %f, height %f, color %s", o.x, o.y, o.width, o.height, o.color)
}

type Circle struct {
	x float64
	y float64
	width float64
	height float64
	color string
}

func (o Circle) X() float64 {
	return o.x
}

func (o Circle) Y() float64 {
	return o.y
}

func (o Circle) Width() float64 {
	return o.width
}

func (o Circle) Height() float64 {
	return o.height
}

func (o Circle) Color() string {
	return o.color
}

func (o Circle) Image() string {
	return ""
}

func (o Circle) Size() float64 {
	return 0.0
}

func (o Circle) Message() string {
	return ""
}

func (o *Circle) SetX(x float64) {
	o.x = x
}

func (o *Circle) SetY(x float64) {
	o.y = x
}

func (o *Circle) SetWidth(x float64) {
	o.width = x
}

func (o *Circle) SetHeight(x float64) {
	o.height = x
}

func (o *Circle) SetColor(x string) {
	o.color = x
}

func (o *Circle) SetImage(x string) {
	
}

func (o *Circle) SetSize(x float64) {
	
}

func (o *Circle) SetMessage(x string) {
	
}

func (o *Circle) String() string {
	return fmt.Sprintf("Circle = x: %f, y: %f, width: %f, height %f, color %s", o.x, o.y, o.width, o.height, o.color)
}

type Image struct {
	x float64
	y float64
	width float64
	height float64
	image string
}

func (o Image) X() float64 {
	return o.x
}

func (o Image) Y() float64 {
	return o.y
}

func (o Image) Width() float64 {
	return o.width
}

func (o Image) Height() float64 {
	return o.height
}

func (o Image) Color() string {
	return ""
}

func (o Image) Image() string {
	return o.image
}

func (o Image) Size() float64 {
	return 0.0
}

func (o Image) Message() string {
	return ""
}

func (o *Image) SetX(x float64) {
	o.x = x
}

func (o *Image) SetY(x float64) {
	o.y = x
}

func (o *Image) SetWidth(x float64) {
	o.width = x
}

func (o *Image) SetHeight(x float64) {
	o.height = x
}

func (o *Image) SetColor(x string) {

}

func (o *Image) SetImage(x string) {
	o.image = x
}

func (o *Image) SetSize(x float64) {
	
}

func (o *Image) SetMessage(x string) {
	
}

func (o *Image) String() string {
	return fmt.Sprintf("Image = x: %f, y: %f, width: %f, height: %f, image: %s", o.x, o.y, o.width, o.height, o.image)
}

type Text struct {
	x float64
	y float64
	size float64
	color string
	message string
}

func (o Text) X() float64 {
	return o.x
}

func (o Text) Y() float64 {
	return o.y
}

func (o Text) Width() float64 {
	return 0.0
}

func (o Text) Height() float64 {
	return 0.0
}

func (o Text) Color() string {
	return o.color
}

func (o Text) Image() string {
	return ""
}

func (o Text) Size() float64 {
	return o.size
}

func (o Text) Message() string {
	return o.message
}

func (o *Text) SetX(x float64) {
	o.x = x
}

func (o *Text) SetY(x float64) {
	o.y = x
}

func (o *Text) SetWidth(x float64) {

}

func (o *Text) SetHeight(x float64) {

}

func (o *Text) SetColor(x string) {
	o.color = x
}

func (o *Text) SetImage(x string) {
	
}

func (o *Text) SetSize(x float64) {
	o.size = x
}

func (o *Text) SetMessage(x string) {
	o.message = x
}

func (o *Text) String() string {
	return fmt.Sprintf("Text = x: %f, y: %f, size: %f, message: %s, color: %s", o.x, o.y, o.size, o.message, o.color)
}

type Background struct {
	x float64
	y float64
	width float64
	height float64
	image string
}

func (o Background) X() float64 {
	return o.x
}

func (o Background) Y() float64 {
	return o.y
}

func (o Background) Width() float64 {
	return o.width
}

func (o Background) Height() float64 {
	return o.height
}

func (o Background) Color() string {
	return ""
}

func (o Background) Image() string {
	return o.image
}

func (o Background) Size() float64 {
	return 0.0
}

func (o Background) Message() string {
	return ""
}

func (o *Background) SetX(x float64) {
	o.x = x
}

func (o *Background) SetY(x float64) {
	o.y = x
}

func (o *Background) SetWidth(x float64) {
	o.width = x
}

func (o *Background) SetHeight(x float64) {
	o.height = x
}

func (o *Background) SetColor(x string) {

}

func (o *Background) SetImage(x string) {
	o.image = x
}

func (o *Background) SetSize(x float64) {
	
}

func (o *Background) SetMessage(x string) {
	
}

func (o *Background) String() string {
	return fmt.Sprintf("Background = x: %f, y: %f, width: %f, height: %f, image: %s", o.x, o.y, o.width, o.height, o.image)
}