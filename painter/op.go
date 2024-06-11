package painter

import (
	"golang.org/x/image/draw"
	"image"
	"image/color"

	"golang.org/x/exp/shiny/screen"
)

type Operation interface {
	Do(t screen.Texture) (ready bool)
}
type OperationList []Operation

func (ol OperationList) Do(t screen.Texture) (ready bool) {
	for _, o := range ol {
		ready = o.Do(t) || ready
	}
	return
}

var UpdateOp = updateOp{}

type updateOp struct{}

func (op updateOp) Do(t screen.Texture) bool { return true }

type OperationFunc func(t screen.Texture)

func (f OperationFunc) Do(t screen.Texture) bool {
	f(t)
	return false
}

type BgRect struct {
	FirstPoint  image.Point
	SecondPoint image.Point
}

func (op *BgRect) Do(t screen.Texture) bool {
	c := color.Black
	t.Fill(image.Rect(op.FirstPoint.X, op.FirstPoint.Y, op.SecondPoint.X, op.SecondPoint.Y), c, screen.Src)
	return false
}

type Figure struct {
	CentralPoint image.Point
}

func (op *Figure) Do(t screen.Texture) bool {
	c := color.RGBA{R: 255, G: 255}
	t.Fill(image.Rect(op.CentralPoint.X-75, op.CentralPoint.Y-50, op.CentralPoint.X+75, op.CentralPoint.Y), c, draw.Src)
	t.Fill(image.Rect(op.CentralPoint.X-25, op.CentralPoint.Y, op.CentralPoint.X+25, op.CentralPoint.Y+50), c, draw.Src)
	return false
}

type Move struct {
	X            int
	Y            int
	FiguresArray []*Figure
}

func (op *Move) Do(t screen.Texture) bool {
	for i := range op.FiguresArray {
		op.FiguresArray[i].CentralPoint.X += op.X
		op.FiguresArray[i].CentralPoint.Y += op.Y
	}
	return false
}

func WhiteFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.White, screen.Src)
}
func GreenFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.RGBA{G: 0xff, A: 0xff}, screen.Src)
}
func Reset(t screen.Texture) {
	t.Fill(t.Bounds(), color.Black, screen.Src)
}
