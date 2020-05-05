package gofb

type SpriteSheet struct {
	surface     *Surface
	frameWidth  int
	frameHeight int
}

func NewSpriteSheetFromFile(file string) (*SpriteSheet, error) {
	s, err := NewSurfaceFromFile(file)
	if err != nil {
		return nil, err
	}
	return &SpriteSheet{
		surface:     s,
		frameWidth:  0,
		frameHeight: 0,
	}, nil
}

func (s *SpriteSheet) SetFrameRegion(width int, height int) {
	s.frameWidth = width
	s.frameHeight = height
}

func (s *SpriteSheet) Surface() *Surface {
	return s.surface
}

func (s *SpriteSheet) DrawFrame(x int, y int, fx int, fy int) {
	r := NewRegion(fx*s.frameWidth, fy*s.frameHeight, s.frameWidth, s.frameHeight)
	s.surface.DrawRegion(x, y, r)
}

func (s *SpriteSheet) Release() {
	s.surface.Release()
}
