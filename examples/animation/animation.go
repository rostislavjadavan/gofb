package main

import (
	"github.com/rostislavjadavan/gofb"
)

func main() {
	w := gofb.NewWindow("go-fb", 1200, 900, false)

	spriteSheet, err := gofb.NewSpriteSheetFromFile("../assets/person_atlas.png", 128, 163)
	if err != nil {
		panic(err)
	}

	spriteSheet.Surface().Scale = 2

	var frameUpdateTimeMs int64 = 0 // how much time elapsed
	var frame = 0                   // current animation frame

	for w.IsRunning() {
		w.StartFrame()
		w.Clear(gofb.NewColor3(120, 220, 230))

		spriteSheet.Surface().FlipHorizontal = false
		spriteSheet.DrawTile(200, 250, frame, 0)

		spriteSheet.Surface().FlipHorizontal = true
		spriteSheet.DrawTile(600, 250, frame, 1)

		frameUpdateTimeMs += w.GetDeltaTimeMs()
		// when 100ms elapsed switch to next frame from sprite sheet (and reset timer)
		if frameUpdateTimeMs > 100 {
			frame++
			frameUpdateTimeMs = 0
			// there are 8 frames in sprite sheet
			if frame > 8 {
				frame = 0
			}
		}

		w.FinalizeFrame()
	}

	defer spriteSheet.Release()
	defer w.Destroy()
}
