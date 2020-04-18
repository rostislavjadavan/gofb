package gofb

import (
	"github.com/go-gl/gl/v2.1/gl"
)

// Color RGBA byte color
type Color struct {
	r uint8
	g uint8
	b uint8
	a uint8
}

// NewColor create new Color
func NewColor(r uint8, g uint8, b uint8, a uint8) Color {
	return Color{r: r, g: g, b: b, a: a}
}

// ColorWhite is white Color
var ColorWhite = Color{r: 255, g: 255, b: 255, a: 255}

// ColorBlack is black Color
var ColorBlack = Color{r: 0, g: 0, b: 0, a: 255}

// GL set color to OpenGL context
func (c *Color) GL() {
	gl.Color4ub(c.r, c.g, c.b, c.a)
}

// GLClear set color as OpenGL color for clearing the screen
func (c *Color) GLClear() {
	gl.ClearColor(float32(c.r)/255, float32(c.g)/255, float32(c.b)/255, float32(c.a)/255)
}

// Region Screen region
type Region struct {
	x int
	y int
	w int
	h int
}

// NewRegion create new region
func NewRegion(x int, y int, w int, h int) Region {
	return Region{x: x, y: y, w: w, h: h}
}

// Point2 2D point
type Point2 struct {
	X float32
	Y float32
}

// NewPoint2 create new Point
func NewPoint2(x float32, y float32) Point2 {
	return Point2{X: x, Y: y}
}

// GL set point as OpenGL vertex
func (p *Point2) GL() {
	gl.Vertex2f(p.X, p.Y)
}
