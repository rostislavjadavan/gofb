package gofb

import (
	"errors"
	"github.com/go-gl/gl/v2.1/gl"
	"image"
	"image/draw"
	"os"
)

// Surface represent pixel buffer
type Surface struct {
	Width       int
	Height      int
	Scale       int
	Rotation    float32
	pixels      *[]byte
	texture     *Texture
	needsUpdate bool
}

// NewSurface create new empty surface
func NewSurface(width int, height int) *Surface {
	pixels := make([]byte, width*height*4)
	tex := NewTextureFromBytes(width, height, &pixels)
	return &Surface{
		Width:       width,
		Height:      height,
		Scale:       1,
		pixels:      &pixels,
		texture:     tex,
		needsUpdate: false,
	}
}

// NewSurfaceFromBytes create new surface from given input byte array (expecting RGBA format)
func NewSurfaceFromBytes(width int, height int, bytes *[]byte) *Surface {
	tex := NewTextureFromBytes(width, height, bytes)
	return &Surface{
		Width:       width,
		Height:      height,
		Scale:       1,
		Rotation:    0,
		pixels:      bytes,
		texture:     tex,
		needsUpdate: false,
	}
}

// NewSurfaceFromFile create surface from image file
func NewSurfaceFromFile(file string) (*Surface, error) {
	imgFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, err
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return nil, errors.New("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	return NewSurfaceFromBytes(rgba.Rect.Size().X, rgba.Rect.Size().Y, &rgba.Pix), nil
}

// SetPixel draw pixel at surface
func (s *Surface) SetPixel(x int, y int, c Color) {
	i := (y*s.Width + x) * 4
	p := *s.pixels
	p[i+0] = c.r
	p[i+1] = c.g
	p[i+2] = c.b
	p[i+3] = c.a
	s.needsUpdate = true
}

// GetPixel return color of pixel
func (s *Surface) GetPixel(x int, y int) Color {
	i := (y*s.Width + x) * 4
	p := *s.pixels
	return Color{
		r: p[i+0],
		g: p[i+1],
		b: p[i+2],
		a: p[i+3],
	}
}

// Draw surface on the screen
func (s *Surface) Draw(x int, y int) {
	s.draw(
		NewPoint2(float32(x), float32(y)),
		NewPoint2(float32(s.Width*s.Scale), float32(s.Height*s.Scale)),
		NewPoint2(0, 0),
		NewPoint2(1, 1),
		s.texture,
	)
}

// DrawRegion draw region of surface on the screen
func (s *Surface) DrawRegion(x int, y int, r Region) {
	s.draw(
		NewPoint2(float32(x), float32(y)),
		NewPoint2(float32(r.w*s.Scale), float32(r.h*s.Scale)),
		NewPoint2(float32(r.x)/float32(s.Width), float32(r.y)/float32(s.Height)),
		NewPoint2(float32(r.x+r.w)/float32(s.Width), float32(r.y+r.h)/float32(s.Height)),
		s.texture,
	)
}

// Release surface from memory and gpu
func (s *Surface) Release() {
	s.pixels = nil
	s.texture.Release()
}

func (s *Surface) update() {
	s.texture.Release()
	s.texture = NewTextureFromBytes(s.Width, s.Height, s.pixels)
}

func (s *Surface) draw(pos Point2, size Point2, t1 Point2, t2 Point2, tex *Texture) {
	if s.needsUpdate {
		s.update()
		s.needsUpdate = false
	}

	gl.Enable(gl.TEXTURE_2D)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	tex.Bind()

	gl.PushMatrix()
	gl.LoadIdentity()
	gl.Translatef(pos.X+size.X/2, pos.Y+size.Y/2, 0)
	gl.Rotatef(s.Rotation, 0, 0, 1)

	ColorWhite.GL()
	gl.Begin(gl.QUADS)
	gl.TexCoord2f(t1.X, t1.Y)
	gl.Vertex2f(-size.X/2, -size.Y/2)
	gl.TexCoord2f(t2.X, t1.Y)
	gl.Vertex2f(size.X/2, -size.Y/2)
	gl.TexCoord2f(t2.X, t2.Y)
	gl.Vertex2f(size.X/2, size.Y/2)
	gl.TexCoord2f(t1.X, t2.Y)
	gl.Vertex2f(-size.X/2, size.Y/2)
	gl.End()
	gl.PopMatrix()
}
