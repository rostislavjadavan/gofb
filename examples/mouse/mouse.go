package main

import (
	"github.com/rostislavjadavan/gofb"
	"strconv"
)

// Linear interpolation between points p, q and time t=[0, 1]
func lerp(t float32, px, py, qx, qy int) (x, y int) {
	return int(float32(px) + t*float32(qx-px)), int(float32(py) + t*float32(qy-py))
}

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
	sourceStarPos := starPos
	targetStarPos := starPos

	var t float64 = 1.0

	for w.IsRunning() {
		w.StartFrame()
		w.Clear(gofb.NewColor3(84, 197, 211))

		speed := float64(w.GetDeltaTimeMs()) / 1000
		cursorPos := w.GetCursorPos()

		if w.IsInput(gofb.KeyEscape) {
			w.Stop()
		}
		if w.IsInput(gofb.MouseButtonLeft) {
			sourceStarPos = starPos.Clone()
			targetStarPos = cursorPos.Clone()
			t = 0
		}

		if t <= 1 {
			starPos = sourceStarPos.Lerp(t, targetStarPos)
			t += speed
			if t > 1 {
				starPos = targetStarPos
			}
			text.Draw("X", targetStarPos.X-15, targetStarPos.Y-30, gofb.NewColor3(0, 0, 0))
		}

		star.Draw(starPos.X-star.Width/2, starPos.Y-star.Height/2)

		posAsString := strconv.FormatFloat(float64(cursorPos.X), 'f', 0, 32) + ", " + strconv.FormatFloat(float64(cursorPos.Y), 'f', 0, 32)
		text.Draw("Mouse is at "+posAsString, 100, 760, gofb.NewColor3(0, 0, 0))
		text.Draw("Press escape to exit", 100, 800, gofb.NewColor3(0, 0, 0))

		w.FinalizeFrame()
	}

	defer star.Release()
	defer w.Destroy()
}
