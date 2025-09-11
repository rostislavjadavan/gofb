package gofb

import (
	"errors"
	"os"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/gltext"
)

// RenderContext holds the current rendering context information
type RenderContext struct {
	scale float32
}

// Global render context that tracks HiDPI scaling
var globalRenderContext = &RenderContext{scale: 1.0}

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
	if !openglInitialized {
		panic(errors.New("cannot draw text: OpenGL context not initialized. Create a Window first"))
	}
	c.glColor()
	gl.LoadIdentity()
	// Apply HiDPI scaling using global render context
	scaledX := float32(x) * globalRenderContext.scale
	scaledY := float32(y) * globalRenderContext.scale
	f.Handle.Printf(scaledX, scaledY, str)
}
