package gofb

import (
	"github.com/go-gl/gltext"
	"os"
)

// Font handle
type Font struct {
	Handle *gltext.Font
}

// Create new Font from given filename (.ttf expected)
func NewFont(filename string, size int32) (*Font, error) {
	fd, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer fd.Close()

	font, _ := gltext.LoadTruetype(fd, size, 32, 127, gltext.LeftToRight)
	return &Font{Handle: font}, nil
}

// Draw text on the screen
func (f *Font) Draw(str string, x int, y int, c Color) {
	c.GL()
	f.Handle.Printf(float32(x), float32(y), str)
}
