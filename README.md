![](gofb_500.png)

# gofb

Simple framebuffer library for golang using OpenGL 2.1.

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
	surface.SetPixel(300, 300, gofb.NewColor(255, 255, 255, 255))

	for w.IsRunning() {
		w.StartFrame()
		w.Clear(gofb.NewColor(0, 0, 0, 255))

		// Draw buffer on screen
		surface.Draw(0, 0)

		w.FinalizeFrame()
	}

	defer surface.Release()
	defer w.Destroy()
}
```

## More examples

http://github.com/rostislavjadavan/gofb-examples

## Features

- draw pixel by pixel
- load png,jpg images
- scale to mimic pixel art
- use custom `ttf` font 
- draw sub region of image
- keyboard and mouse input
