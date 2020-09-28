package gofb

// SpriteSheet represents multiple sprites on one image
type SpriteSheet struct {
	surface    *Surface
	tileWidth  int
	tileHeight int
}

// NewSpriteSheetFromSurface create sprite sheet from surface
func NewSpriteSheetFromSurface(s *Surface, tileWidth, tileHeight int) (*SpriteSheet, error) {
	return &SpriteSheet{
		surface:    s,
		tileWidth:  tileWidth,
		tileHeight: tileHeight,
	}, nil
}

// NewSpriteSheetFromFile create sprite sheet from file
func NewSpriteSheetFromFile(file string, tileWidth, tileHeight int) (*SpriteSheet, error) {
	s, err := NewSurfaceFromFile(file)
	if err != nil {
		return nil, err
	}
	return NewSpriteSheetFromSurface(s, tileWidth, tileHeight)
}

// SetTileSize set size of one sprite
func (s *SpriteSheet) SetTileSize(width int, height int) {
	s.tileWidth = width
	s.tileHeight = height
}

// Surface return surface instance
func (s *SpriteSheet) Surface() *Surface {
	return s.surface
}

// DrawTile draw tile
func (s *SpriteSheet) DrawTile(x int, y int, tileX int, tileY int) {
	r := NewRegion(tileX*s.tileWidth, tileY*s.tileHeight, s.tileWidth, s.tileHeight)
	s.surface.DrawRegion(x, y, r)
}

// Release release structure from memory
func (s *SpriteSheet) Release() {
	s.surface.Release()
}
