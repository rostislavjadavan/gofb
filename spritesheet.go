package gofb

type SpriteSheet struct {
	surface    *Surface
	tileWidth  int
	tileHeight int
}

func NewSpriteSheetFromFile(file string) (*SpriteSheet, error) {
	s, err := NewSurfaceFromFile(file)
	if err != nil {
		return nil, err
	}
	return &SpriteSheet{
		surface:    s,
		tileWidth:  0,
		tileHeight: 0,
	}, nil
}

func (s *SpriteSheet) SetTileSize(width int, height int) {
	s.tileWidth = width
	s.tileHeight = height
}

func (s *SpriteSheet) Surface() *Surface {
	return s.surface
}

func (s *SpriteSheet) DrawTile(x int, y int, tileX int, tileY int) {
	r := NewRegion(tileX*s.tileWidth, tileY*s.tileHeight, s.tileWidth, s.tileHeight)
	s.surface.DrawRegion(x, y, r)
}

func (s *SpriteSheet) Release() {
	s.surface.Release()
}
