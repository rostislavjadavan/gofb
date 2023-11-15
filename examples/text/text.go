package main

import (
	"github.com/rostislavjadavan/gofb"
	"strconv"
)

func main() {
	w := gofb.NewWindow("go-fb", 1200, 900, false)
	bg := gofb.NewSurface(1200, 900)

	header, err := gofb.NewFont("../assets/uni0553-webfont.ttf", 160)
	if err != nil {
		panic(err)
	}

	text, err := gofb.NewFont("../assets/uni0553-webfont.ttf", 40)
	if err != nil {
		panic(err)
	}

	for w.IsRunning() {
		w.StartFrame()
		w.Clear(gofb.NewColor3(100, 100, 100))

		header.Draw("GOFB", 100, 40, gofb.NewColor3(255, 255, 0))
		text.Draw("Simple framebuffer library", 100, 220, gofb.NewColor3(200, 200, 200))

		fps := strconv.Itoa(int(w.GetFPS()))
		text.Draw("Running "+fps+" frames per second", 100, 800, gofb.NewColor3(150, 150, 150))

		w.FinalizeFrame()
	}

	defer bg.Release()
	defer w.Destroy()
}
