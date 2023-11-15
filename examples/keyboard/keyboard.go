package main

import (
	"github.com/rostislavjadavan/gofb"
)

func main() {
	w := gofb.NewWindow("go-fb", 1200, 900, false)

	text, err := gofb.NewFont("../assets/uni0553-webfont.ttf", 40)
	if err != nil {
		panic(err)
	}

	star, err := gofb.NewSurfaceFromFile("../assets/pixel_star.png")
	if err != nil {
		panic(err)
	}

	starPos := gofb.NewVec2(500, 400)

	for w.IsRunning() {
		w.StartFrame()
		w.Clear(gofb.NewColor3(84, 197, 211))

		star.Drawv(starPos)

		speed := int(w.GetDeltaTimeMs())
		if w.IsInput(gofb.KeyShift) {
			speed = int(w.GetDeltaTimeMs()) / 5
		}

		if w.IsInput(gofb.KeyUp) {
			if starPos.Y > 0 {
				starPos.Y -= speed
			}
		}
		if w.IsInput(gofb.KeyDown) {
			if starPos.Y < w.Height-star.Height {
				starPos.Y += speed
			}
		}
		if w.IsInput(gofb.KeyLeft) {
			if starPos.X > 0 {
				starPos.X -= speed
			}
		}
		if w.IsInput(gofb.KeyRight) {
			if starPos.X < w.Width-star.Width {
				starPos.X += speed
			}
		}
		if w.IsInput(gofb.KeyEscape) {
			w.Stop()
		}

		text.Draw("Use cursor keys to move", 100, 720, gofb.NewColor3(0, 0, 0))
		text.Draw("Hold shift to move slower", 100, 760, gofb.NewColor3(0, 0, 0))
		text.Draw("Press escape to exit", 100, 800, gofb.NewColor3(0, 0, 0))

		w.FinalizeFrame()
	}

	defer star.Release()
	defer w.Destroy()
}
