package main

import (
	"github.com/rostislavjadavan/gofb"
	"math/rand"
)

func main() {
	w := gofb.NewWindow("go-fb", 1200, 900, false)
	bg := gofb.NewSurface(100, 100)
	bg.Scale = 6

	for y := 0; y < bg.Height; y++ {
		for x := 0; x < bg.Width; x++ {
			r := 100 + uint8(rand.Intn(155))
			g := 100 + uint8(rand.Intn(155))
			b := 100 + uint8(rand.Intn(135))
			bg.SetPixel(x, y, gofb.NewColor3(r, g, b))
		}
	}

	for w.IsRunning() {
		w.StartFrame()
		w.Clear(gofb.NewColor3(120, 220, 230))

		bg.Draw(300, 150)
		//bg.Rotation += float32(w.GetDeltaTimeMs() / 10)

		w.FinalizeFrame()
	}

	defer bg.Release()
	defer w.Destroy()
}
