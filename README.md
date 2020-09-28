![](gofb_500.png)

# gofb

Framebuffer library for golang

[![Go Report Card](https://goreportcard.com/badge/github.com/rostislavjadavan/gofb)](https://goreportcard.com/report/github.com/rostislavjadavan/gofb)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/rostislavjadavan/gofb)](https://pkg.go.dev/github.com/rostislavjadavan/gofb)

## Example

```go
package main

import (
	"github.com/rostislavjadavan/gofb"
)

func main() {
	// Create Window
	w := gofb.NewWindow("go-fb", 600, 600, false)

	// Create pixel buffer
	surface := gofb.NewSurface(600, 600)

	// Draw pixel into buffer
	surface.SetPixel(300, 300, gofb.NewColor3(255, 255, 255))

	for w.IsRunning() {
		w.StartFrame()
		w.Clear(gofb.ColorBlack)

		// Draw buffer on the screen
		surface.Draw(0, 0)

		w.FinalizeFrame()
	}

	defer surface.Release()
	defer w.Destroy()
}
```

## Features

- draw pixel by pixel
- load `png`, `jpg` images
- scale to mimic pixel art
- use custom `ttf` font 
- sprite sheets (for tile set rendering)
- keyboard and mouse input

## More examples

http://github.com/rostislavjadavan/gofb-examples

![](https://github.com/rostislavjadavan/gofb-examples/raw/master/simple/preview.jpg)
![](https://github.com/rostislavjadavan/gofb-examples/raw/master/animation/preview.gif)
![](https://github.com/rostislavjadavan/gofb-examples/raw/master/tunnel/preview.jpg)
![](https://github.com/rostislavjadavan/gofb-examples/raw/master/mouse/preview.gif)
![](https://github.com/rostislavjadavan/gofb-examples/raw/master/roguelike/preview.jpg)

## Libraries used

- https://github.com/go-gl/glfw
- https://github.com/go-gl/gltext

## Logo

http://github.com/rostislavjadavan/gofb-logo

## License

[MIT](LICENSE)
