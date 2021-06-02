package engine

// importar pixel
import (
	"image"
	"os"
	"fmt"
	"strconv"

	_ "image/png"

	"github.com/sdkvictor/golang-compiler/objects"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/colornames"
	"github.com/nfnt/resize"
)

/*
const (
	windowWidth  = 1024
	windowHeight = 768
)
*/

type Engine struct {
	win *pixelgl.Window
	imd *imdraw.IMDraw

}

type RGB struct {
	Red   uint8
	Green uint8
	Blue  uint8
}


func toRGB(h string) (RGB, error) {
	return Hex2RGB(h)
}

func Hex2RGB(hex string) (RGB, error) {
	var rgb RGB
	values, err := strconv.ParseUint(hex[1:len(hex)-1], 16, 32)

	if err != nil {
		return RGB{}, err
	}

	rgb = RGB{
		Red:   uint8(values >> 16),
		Green: uint8((values >> 8) & 0xFF),
		Blue:  uint8(values & 0xFF),
	}

	return rgb, nil
}


func loadPicture(path string, i objects.Image) (pixel.Picture, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	newImage := resize.Resize(uint(i.Width()), uint(i.Height()), img, resize.Lanczos3)
	return pixel.PictureDataFromImage(newImage), nil
}

func (e *Engine) DrawSquare(s objects.Square) {
	c, _:= toRGB(s.Color())
	e.imd.Color = pixel.RGB(float64(c.Red)/255, float64(c.Green)/255, float64(c.Blue)/255)
	r := pixel.R(s.X(), s.Y(), s.X()+s.Width(), s.Y()+s.Height())
	e.imd.Push(r.Min, r.Max)
	e.imd.Rectangle(0)
	e.imd.Draw(e.win)
}

func (e *Engine) DrawCircle(c objects.Circle) {
	r,_ := toRGB(c.Color())
	e.imd.Color = pixel.RGB(float64(r.Red)/255, float64(r.Green)/255, float64(r.Blue)/255)
	e.imd.Push(pixel.V(c.X(), c.Y()))
	e.imd.Ellipse(pixel.V(c.Width()/2, c.Height()/2), 0)
	e.imd.Draw(e.win)
}

func (e *Engine) DrawText(t objects.Text) {
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(t.X(), t.Y()), basicAtlas)

	fmt.Fprintln(basicTxt, t.Message()[1:len(t.Message())-1])
	basicTxt.Draw(e.win, pixel.IM)
}

func (e *Engine) DrawImage(i objects.Image) {
	img := string(i.Image())
	pic, err := loadPicture(img[1:len(img)-1], i)
	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(pic, pic.Bounds())
	sprite.Draw(e.win, pixel.IM.Moved(pixel.V(i.X(), i.Y())))
}

func (e *Engine) WindowClosed() bool {
	return e.win.Closed()
}


func (e *Engine) KeyPressed(k string) bool {
	key := k[1:len(k)-1]
	switch key {
	case "Space":
		return e.win.Pressed(pixelgl.KeySpace)
	case "Up":
		return e.win.Pressed(pixelgl.KeyUp)	
	case "Down":
		return e.win.Pressed(pixelgl.KeyDown)
	case "Right":
		return e.win.Pressed(pixelgl.KeyRight)
	case "Left":
		return e.win.Pressed(pixelgl.KeyLeft)
	case "W":
		return e.win.Pressed(pixelgl.KeyW)
	case "A":
		return e.win.Pressed(pixelgl.KeyA)
	case "S":
		return e.win.Pressed(pixelgl.KeyS)
	case "D":
		return e.win.Pressed(pixelgl.KeyD)
	}
	return false;
}

func (e *Engine) IntersectSquare(s1, s2 objects.Square) bool{
	r1 := pixel.R(s1.X(), s1.Y(), s1.X() + s1.Width(), s1.Y() + s1.Height())
	
	r2 := pixel.R(s2.X(), s2.Y(), s2.X() + s2.Width(), s2.Y() + s2.Height())
	
	return r1.Intersects(r2)
}

func (e *Engine) IntersectCircle(s1, s2 objects.Circle) bool{
	r1 := pixel.R(s1.X(), s1.Y(), s1.X() + s1.Width(), s1.Y() + s1.Height())
	
	r2 := pixel.R(s2.X(), s2.Y(), s2.X() + s2.Width(), s2.Y() + s2.Height())
	
	return r1.Intersects(r2)
}

func (e *Engine) IntersectCS(s1 objects.Circle, s2 objects.Square) bool{
	r1 := pixel.R(s1.X(), s1.Y(), s1.X() + s1.Width(), s1.Y() + s1.Height())
	//fmt.Printf(r1.String())
	r2 := pixel.R(s2.X(), s2.Y(), s2.X() + s2.Width(), s2.Y() + s2.Height())
	//fmt.Printf(r2.String())
	return r1.Intersects(r2)

}

func (e *Engine) IntersectSC(s1 objects.Square, s2 objects.Circle) bool{
	r1 := pixel.R(s1.X(), s1.Y(), s1.X() + s1.Width(), s1.Y() + s1.Height())
	
	r2 := pixel.R(s2.X(), s2.Y(), s2.X() + s2.Width(), s2.Y() + s2.Height())
	
	return r1.Intersects(r2)
}

func (e *Engine) Clear() {
	e.win.Clear(colornames.Black)
	e.imd.Clear()
	//e.imd.Reset()
}

func (e *Engine) Update() {
	e.win.Update()
}

func NewEngine(name string, windowWidth float64, windowHeight float64) *Engine {
	cfg := pixelgl.WindowConfig{
        Title:  name,
        Bounds: pixel.R(0, 0, windowWidth, windowHeight),
        VSync:  true,
    }
    win, err := pixelgl.NewWindow(cfg)
    if err != nil {
        panic(err)
    }

	win.Clear(colornames.Black)

	return &Engine{
		win,
		imdraw.New(nil),
	}
}