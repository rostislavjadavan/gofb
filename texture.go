package gofb

import (
	"github.com/go-gl/gl/v2.1/gl"
	_ "image/jpeg"
	_ "image/png"
	"strconv"
)

type Texture struct {
	Handle uint32
	Name   string
	Width  int
	Height int
}

func (t *Texture) Bind() {
	gl.BindTexture(gl.TEXTURE_2D, t.Handle)
}

func (t *Texture) Release() {
	gl.DeleteTextures(1, &t.Handle)
}

func NewTextureFromBytes(width int, height int, pixels *[]byte) *Texture {
	var handle uint32
	gl.Enable(gl.TEXTURE_2D)
	gl.GenTextures(1, &handle)
	gl.BindTexture(gl.TEXTURE_2D, handle)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(width),
		int32(height),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(*pixels))

	return &Texture{
		Handle: handle,
		Name:   "texture_" + strconv.Itoa(int(handle)),
		Width:  width,
		Height: height,
	}
}
