package gofb

import (
	"math"

	"github.com/go-gl/gl/v2.1/gl"
)

// Color RGBA byte color
type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

// NewColor3 create new Color (without alpha)
func NewColor3(r uint8, g uint8, b uint8) Color {
	return Color{R: r, G: g, B: b, A: 255}
}

// NewColor4 create new Color with alpha component
func NewColor4(r uint8, g uint8, b uint8, a uint8) Color {
	return Color{R: r, G: g, B: b, A: a}
}

// ColorWhite is white Color
var ColorWhite = Color{R: 255, G: 255, B: 255, A: 255}

// ColorBlack is black Color
var ColorBlack = Color{R: 0, G: 0, B: 0, A: 255}

// ColorOpaque is fully transparent color
var ColorOpaque = Color{R: 0, G: 0, B: 0, A: 0}

// GetAsInt return color as RGBA int value
func (c *Color) GetAsInt() int {
	return int(c.R)<<24 | int(c.G)<<16 | int(c.B)<<8 | int(c.A)
}

// GL set color to OpenGL context
func (c *Color) glColor() {
	gl.Color4ub(c.R, c.G, c.B, c.A)
}

// GLClear set color as OpenGL color for clearing the screen
func (c *Color) glClear() {
	gl.ClearColor(float32(c.R)/255, float32(c.G)/255, float32(c.B)/255, float32(c.A)/255)
}

// Region Screen region
type Region struct {
	X int
	Y int
	W int
	H int
}

// NewRegion create new region
func NewRegion(x int, y int, w int, h int) Region {
	return Region{X: x, Y: y, W: w, H: h}
}

// IsInside return if given vector is inside region
func (r *Region) IsInside(v Vec2) bool {
	return v.X >= r.X && v.X < r.X+r.W && v.Y >= v.Y && v.Y < r.Y+r.H
}

// Vec2 2D vector
type Vec2 struct {
	X int
	Y int
}

// NewVec2 return new instance
func NewVec2(x int, y int) Vec2 {
	return Vec2{X: x, Y: y}
}

// Abs absolute value
func (v *Vec2) Abs() Vec2 {
	return NewVec2(int(math.Abs(float64(v.X))), int(math.Abs(float64(v.Y))))
}

// Sub substract two vectors
func (v *Vec2) Sub(v2 *Vec2) *Vec2 {
	return &Vec2{v.X - v2.X, v.Y - v2.Y}
}

// Subi substract int from both components
func (v *Vec2) Subi(i int) *Vec2 {
	return &Vec2{v.X - i, v.Y - i}
}

// Distance return distance to given vector
func (v *Vec2) Distance(v2 Vec2) float64 {
	v3 := v2.Sub(v).Abs()
	return math.Sqrt(float64(v3.X*v3.X) + float64(v3.Y*v3.Y))
}

// IsInside check if vector is inside region
func (v *Vec2) IsInside(a Vec2) bool {
	return v.X >= 0 && v.X < a.X && v.Y >= 0 && v.Y < a.Y
}

// IsInside2 check if vector is inside region
func (v *Vec2) IsInside2(a Vec2, b Vec2) bool {
	if a.X > b.X {
		a.X, b.X = b.X, a.X
	}
	if a.Y > b.Y {
		a.Y, b.Y = b.Y, a.Y
	}

	return v.X >= a.X && v.X < b.X && v.Y >= a.Y && v.Y < b.Y
}

// ClipInside2 Clip vector into given region
func (v *Vec2) ClipInside2(a Vec2, b Vec2) {
	if a.X > b.X {
		a.X, b.X = b.X, a.X
	}
	if a.Y > b.Y {
		a.Y, b.Y = b.Y, a.Y
	}

	if v.X < a.X {
		v.X = a.X
	}
	if v.X > b.X {
		v.X = b.X
	}
	if v.Y < a.Y {
		v.Y = a.Y
	}
	if v.Y > b.Y {
		v.Y = b.Y
	}
}

// Clone return copy of vector
func (v *Vec2) Clone() Vec2 {
	return Vec2{X: v.X, Y: v.Y}
}

// Lerp linear interpolation between two vectors, t = [0,1]
func (v *Vec2) Lerp(t float64, target Vec2) Vec2 {
	if t < 0 {
		t = 0
	}
	if t > 1 {
		t = 1
	}
	return NewVec2(
		int((1.0-t)*float64(v.X)+t*float64(target.X)),
		int((1.0-t)*float64(v.Y)+t*float64(target.Y)),
	)
}

func (v *Vec2) floatX() float32 {
	return float32(v.X)
}

func (v *Vec2) floatY() float32 {
	return float32(v.Y)
}
