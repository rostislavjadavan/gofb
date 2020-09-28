package gofb

import (
	"os"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/gltext"
)

// Font handle
type Font struct {
	Handle *gltext.Font
}

// NewFont create new Font from given filename (.ttf expected)
func NewFont(filename string, size int) (*Font, error) {
	fd, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer fd.Close()

	font, _ := gltext.LoadTruetype(fd, int32(size), 32, 127, gltext.LeftToRight)
	return &Font{Handle: font}, nil
}

// Drawv text on the screen
func (f *Font) Drawv(str string, v Vec2, c Color) {
	f.Draw(str, v.X, v.Y, c)
}

// Draw text on the screen
func (f *Font) Draw(str string, x, y int, c Color) {
	c.glColor()
	gl.LoadIdentity()
	f.Handle.Printf(float32(x), float32(y), str)
}
