package gofb

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"runtime"
	"time"
)

type Window struct {
	width      int
	height     int
	fullscreen bool
	window     *glfw.Window

	lastFrameTime   time.Time
	frameDeltaMs    int64
	globalElapsedMs int64
	framesCount		int64
}

// NewWindow create new OpenGL window
func NewWindow(name string, width int, height int, fullscreen bool) *Window {
	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)

	var monitor *glfw.Monitor = nil
	if fullscreen {
		monitor = glfw.GetPrimaryMonitor()
	}

	window, err := glfw.CreateWindow(width, height, name, monitor, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	w := Window{width: width, height: height, fullscreen: fullscreen, window: window}
	w.set2DProjection()
	w.lastFrameTime = time.Now()
	window.SetKeyCallback(w.keyCallback)

	return &w
}

// Clear screen
func (w *Window) Clear(c Color) {
	c.GLClear()
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

// StartFrame needs to be at the start of the render loop
func (w *Window) StartFrame() {
	w.frameDeltaMs = time.Since(w.lastFrameTime).Milliseconds()
	w.globalElapsedMs += w.frameDeltaMs
	w.lastFrameTime = time.Now()
}

// FinalizeFrame will swap buffers and poll for events
// This function needs to be called at the end of the render loop
func (w *Window) FinalizeFrame() {
	w.framesCount++
	w.window.SwapBuffers()
	glfw.PollEvents()
}

// GetDeltaTime return number of milliseconds elapsed when rendering last frame
func (w *Window) GetDeltaTimeMs() int64  {
	return w.frameDeltaMs
}

// GetTotalElapsedMs return number of milliseconds elapsed since application started
func (w *Window) GetTotalElapsedMs() int64  {
	return w.globalElapsedMs
}

func (w *Window) GetFPS() float32 {
	return float32(w.framesCount) / float32(w.globalElapsedMs) * 1000;
}

// IsRunning check if application is running
func (w *Window) IsRunning() bool {
	return !w.window.ShouldClose()
}

// Destroy will close window
func (w *Window) Destroy() {
	defer glfw.Terminate()
}

func (w *Window) keyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	// TODO
}

func (w *Window) set2DProjection() {
	gl.Disable(gl.DEPTH_TEST)
	gl.Enable(gl.TEXTURE)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, float64(w.width), float64(w.height), 0, 0, 1)

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
}
