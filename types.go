package gofb

import (
	"github.com/go-gl/gl/v2.1/gl"
)

type Color struct {
	r uint8
	g uint8
	b uint8
	a uint8
}

func NewColor(r uint8, g uint8, b uint8, a uint8) Color {
	return Color{r: r, g: g, b: b, a: a}
}

var White = Color{r: 255, g: 255, b: 255, a: 255}
var Black = Color{r: 0, g: 0, b: 0, a: 255}

func (c *Color) GL() {
	gl.Color4ub(c.r, c.g, c.b, c.a)
}

func (c *Color) GLClear() {
	gl.ClearColor(float32(c.r)/255, float32(c.g)/255, float32(c.b)/255, float32(c.a)/255)
}

type Region struct {
	x int
	y int
	w int
	h int
}

func NewRegion(x int, y int, w int, h int) Region {
	return Region{x: x, y: y, w: w, h: h}
}

type Point2 struct {
	X float32
	Y float32
}

func NewPoint2(x float32, y float32) Point2 {
	return Point2{X: x, Y: y}
}

func (p *Point2) GL() {
	gl.Vertex2f(p.X, p.Y)
}
