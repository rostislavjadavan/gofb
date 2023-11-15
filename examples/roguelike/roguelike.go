package main

import "github.com/rostislavjadavan/gofb"

const TILE_WIDTH = 12
const TILE_HEIGHT = 12
const TILE_SCALE = 4

type WorldMap struct {
	Tiles  []uint16
	Width  int
	Height int
}

func NewWorldMap(width, height int) *WorldMap {
	return &WorldMap{
		Tiles:  make([]uint16, width*height),
		Width:  width,
		Height: height,
	}
}

func (m *WorldMap) Render(posX, posY int, sheet *gofb.SpriteSheet) {
	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			tile := m.Tiles[y*m.Width+x]
			tx := tile%20 - 1
			ty := tile / 20
			sheet.DrawTile(posX+x*TILE_WIDTH*TILE_SCALE, posY+y*TILE_HEIGHT*TILE_SCALE, int(tx), int(ty))
		}
	}
}

func main() {
	w := gofb.NewWindow("go-fb", 1200, 900, false)

	textBig, err := gofb.NewFont("../assets/uni0553-webfont.ttf", 60)
	if err != nil {
		panic(err)
	}

	text, err := gofb.NewFont("../assets/uni0553-webfont.ttf", 30)
	if err != nil {
		panic(err)
	}

	tileset, err := gofb.NewSpriteSheetFromFile("../assets/roguelike.png", TILE_WIDTH, TILE_HEIGHT)
	if err != nil {
		panic(err)
	}
	tileset.Surface().Scale = TILE_SCALE

	m := NewWorldMap(16, 10)
	m.Tiles = []uint16{
		82, 85, 84, 84, 73, 128, 149, 48, 75, 54, 82, 78, 78, 78, 78, 54, 64, 63, 63, 85, 75, 48, 48, 48, 75, 56, 56, 34, 171, 92, 78, 53, 65, 32, 61, 95, 75, 75, 48, 173, 75, 56, 55, 66, 66, 66, 66, 71, 87, 102, 50, 82, 85, 76, 48, 48, 76, 56, 86, 66, 350, 132, 123, 66, 169, 101, 52, 84, 96, 73, 48, 74, 75, 56, 95, 69, 48, 48, 48, 69, 54, 101, 54, 53, 52, 55, 49, 17, 17, 56, 83, 66, 48, 48, 351, 66, 52, 102, 54, 52, 55, 49, 49, 17, 17, 17, 84, 66, 69, 26, 69, 66, 54, 102, 53, 93, 49, 49, 49, 49, 162, 35, 49, 51, 51, 101, 92, 84, 103, 112, 103, 104, 103, 103, 103, 104, 104, 104, 104, 104, 104, 113, 104, 104, 54, 81, 53, 316, 316, 316, 91, 335, 84, 93, 92, 85, 54, 102, 55, 56,
	}

	for w.IsRunning() {
		w.StartFrame()
		w.Clear(gofb.NewColor3(0, 0, 0))

		textBig.Draw("Roguelike", 216, 40, gofb.NewColor3(198, 159, 40))

		m.Render(216, 150, tileset)

		text.Draw("HP: 100", 216, 700, gofb.NewColor3(201, 102, 40))

		w.FinalizeFrame()
	}

	defer tileset.Release()
	defer w.Destroy()
}
