package main

import (
	"github.com/rostislavjadavan/gofb"
	"math"
	"strconv"
)

const (
	W = 600
	H = 450
)

func main() {
	w := gofb.NewWindow("go-fb", 1200, 900, false)

	text, err := gofb.NewFont("../assets/uni0553-webfont.ttf", 40)
	if err != nil {
		panic(err)
	}

	texture, err := gofb.NewSurfaceFromFile("../assets/pattern.jpg")
	if err != nil {
		panic(err)
	}

	surface := gofb.NewSurface(W, H)
	surface.Scale = 2

	distances := [W][H]int{}
	angles := [W][H]int{}

	fw := float64(W)
	fh := float64(H)
	fth := float64(texture.Height)
	ftw := float64(texture.Width)

	for y := 0; y < H; y++ {
		fy := float64(y)
		for x := 0; x < W; x++ {
			fx := float64(x)
			distances[x][y] = int(50.0*fth/math.Sqrt((fx-fw/2.0)*(fx-fw/2.0)+(fy-fh/2.0)*(fy-fh/2.0))) % texture.Height
			angles[x][y] = int(0.5 * ftw * math.Atan2(fy-fh/2.0, fx-fw/2.0) / math.Pi)
		}
	}

	var animation float64 = 0
	var movement float64 = 0

	for w.IsRunning() {
		w.StartFrame()
		w.Clear(gofb.NewColor3(0, 0, 0))

		shiftX := (int)(ftw + animation)
		shiftY := (int)(fth + movement)

		for y := 0; y < H; y++ {
			for x := 0; x < W; x++ {
				tX := (distances[x][y] + shiftX) % texture.Width
				tY := (angles[x][y] + shiftY) % texture.Height

				surface.SetPixel(x, y, texture.GetPixel(tX, tY))
			}
		}

		surface.Draw(0, 0)

		animation += float64(w.GetDeltaTimeMs() / 10)
		movement += float64(w.GetDeltaTimeMs() / 50)

		fps := strconv.Itoa(int(w.GetFPS()))
		text.Draw("GOFB TUNNEL: Running "+fps+" frames per second", 100, 800, gofb.NewColor3(0, 0, 0))

		w.FinalizeFrame()
	}

	defer surface.Release()
	defer texture.Release()
	defer w.Destroy()
}
