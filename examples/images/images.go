package main

import "github.com/rostislavjadavan/gofb"

func main() {
	w := gofb.NewWindow("go-fb", 1200, 900, false)

	bird, err := gofb.NewSurfaceFromFile("../assets/bird.png")
	if err != nil {
		panic(err)
	}

	star, err := gofb.NewSurfaceFromFile("../assets/pixel_star.png")
	if err != nil {
		panic(err)
	}

	for w.IsRunning() {
		w.StartFrame()
		w.Clear(gofb.NewColor3(120, 220, 230))

		bird.Draw(200, 150)

		for i := 0; i < 4; i++ {
			star.Scale = i + 1
			star.Draw(10, 10+i*130)
		}

		w.FinalizeFrame()
	}

	defer bird.Release()
	defer star.Release()
	defer w.Destroy()
}
