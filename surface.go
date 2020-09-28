package gofb

import (
	"errors"
	"image"
	"image/draw"
	"os"

	"github.com/go-gl/gl/v2.1/gl"
)

// Surface represent pixel buffer
type Surface struct {
	Width          int
	Height         int
	Scale          int
	Rotation       float32
	FlipHorizontal bool
	FlipVertical   bool
	pixels         *[]byte
	texture        *Texture
	needsUpdate    bool
}

// NewSurface create new empty surface
func NewSurface(width int, height int) *Surface {
	pixels := make([]byte, width*height*4)
	tex := NewTextureFromBytes(width, height, &pixels)
	return &Surface{
		Width:          width,
		Height:         height,
		Scale:          1,
		FlipHorizontal: false,
		FlipVertical:   false,
		pixels:         &pixels,
		texture:        tex,
		needsUpdate:    false,
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

// IsInsidev check if given point is inside surface
func (s *Surface) IsInsidev(v Vec2) bool {
	return s.IsInside(v.X, v.Y)
}

// IsInside check if given point is inside surface
func (s *Surface) IsInside(x, y int) bool {
	return x > 0 && x < s.Width && y > 0 && y < s.Height
}

// SetPixelv draw pixel at surface
func (s *Surface) SetPixelv(v Vec2, c Color) {
	s.SetPixel(v.X, v.Y, c)
}

// SetPixel draw pixel at surface
func (s *Surface) SetPixel(x int, y int, c Color) {
	i := (y*s.Width + x) * 4
	p := *s.pixels
	p[i+0] = c.R
	p[i+1] = c.G
	p[i+2] = c.B
	p[i+3] = c.A
	s.needsUpdate = true
}

// GetPixelv return color of pixel
func (s *Surface) GetPixelv(v Vec2) Color {
	return s.GetPixel(v.X, v.Y)
}

// GetPixel return color of pixel
func (s *Surface) GetPixel(x int, y int) Color {
	i := (y*s.Width + x) * 4
	p := *s.pixels
	return Color{
		R: p[i+0],
		G: p[i+1],
		B: p[i+2],
		A: p[i+3],
	}
}

// Clear surface with provided color
func (s *Surface) Clear(c Color) {
	for i := 0; i < len(*s.pixels); i += 4 {
		(*s.pixels)[i+0] = c.R
		(*s.pixels)[i+1] = c.G
		(*s.pixels)[i+2] = c.B
		(*s.pixels)[i+3] = c.A
	}
}

// Draw surface on the screen
func (s *Surface) Draw(x int, y int) {
	s.draw(
		NewVec2(x, y),
		NewVec2(s.Width*s.Scale, s.Height*s.Scale),
		NewVec2(0, 0),
		NewVec2(1, 1),
		s.texture,
	)
}

// DrawRegion draw region of surface on the screen
func (s *Surface) DrawRegion(x int, y int, r Region) {
	s.draw(
		NewVec2(x, y),
		NewVec2(r.W*s.Scale, r.H*s.Scale),
		NewVec2(r.X/s.Width, r.Y/s.Height),
		NewVec2((r.X+r.W)/s.Width, (r.Y+r.H)/s.Height),
		s.texture,
	)
}

// Release surface from memory and gpu
func (s *Surface) Release() {
	s.pixels = nil
	s.texture.Release()
}

func (s *Surface) update() {
	s.texture.Update(s.Width, s.Height, s.pixels)
}

func (s *Surface) draw(pos Vec2, size Vec2, t1 Vec2, t2 Vec2, tex *Texture) {
	if s.needsUpdate {
		s.update()
		s.needsUpdate = false
	}
	if s.FlipHorizontal {
		t1.X, t2.X = t2.X, t1.X
	}
	if s.FlipVertical {
		t1.Y, t2.Y = t2.Y, t1.Y
	}

	gl.Enable(gl.TEXTURE_2D)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	tex.Bind()

	gl.LoadIdentity()
	gl.Translatef(pos.floatX()+size.floatX()/2, pos.floatY()+size.floatY()/2, 0)
	gl.Rotatef(s.Rotation, 0, 0, 1)

	ColorWhite.glColor()
	gl.Begin(gl.QUADS)
	gl.TexCoord2f(t1.floatX(), t1.floatY())
	gl.Vertex2f(-size.floatX()/2, -size.floatY()/2)
	gl.TexCoord2f(t2.floatX(), t1.floatY())
	gl.Vertex2f(size.floatX()/2, -size.floatY()/2)
	gl.TexCoord2f(t2.floatX(), t2.floatY())
	gl.Vertex2f(size.floatX()/2, size.floatY()/2)
	gl.TexCoord2f(t1.floatX(), t2.floatY())
	gl.Vertex2f(-size.floatX()/2, size.floatY()/2)
	gl.End()
}
