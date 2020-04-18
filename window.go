package gofb

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"runtime"
	"time"
)

// Window represents window object
type Window struct {
	Width      int
	Height     int
	Fullscreen bool
	window     *glfw.Window
	running    bool

	lastFrameTime   time.Time
	frameDeltaMs    int64
	globalElapsedMs int64
	framesCount     int64

	inputPress []bool
	cursorPos  Point2
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

	var monitor *glfw.Monitor
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

	w := Window{Width: width, Height: height, Fullscreen: fullscreen, window: window}
	w.set2DProjection()
	w.lastFrameTime = time.Now()
	w.inputPress = make([]bool, 505)
	window.SetKeyCallback(w.keyCallback)
	w.cursorPos = NewPoint2(0, 0)
	window.SetCursorPosCallback(w.cursorPositionCallback)
	window.SetMouseButtonCallback(w.mouseButtonCallback)
	w.running = true

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

// GetDeltaTimeMs return number of milliseconds elapsed when rendering last frame
func (w *Window) GetDeltaTimeMs() int64 {
	return w.frameDeltaMs
}

// GetTotalElapsedMs return number of milliseconds elapsed since application started
func (w *Window) GetTotalElapsedMs() int64 {
	return w.globalElapsedMs
}

// GetFPS get average frames per second
func (w *Window) GetFPS() float32 {
	return float32(w.framesCount) / float32(w.globalElapsedMs) * 1000
}

// IsInput check if input was pressed
func (w *Window) IsInput(inputCode int) bool {
	return w.inputPress[inputCode]
}

// GetCursorPos get mouse cursor position
func (w *Window) GetCursorPos() Point2 {
	return w.cursorPos
}

// IsRunning check if application is running
func (w *Window) IsRunning() bool {
	return !w.window.ShouldClose() && w.running
}

// Destroy will close window
func (w *Window) Destroy() {
	defer glfw.Terminate()
}

// Stop running application
func (w *Window) Stop() {
	w.running = false
}

func (w *Window) keyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	switch action {
	case glfw.Repeat:
	case glfw.Press:
		w.inputPress[key] = true
		break
	default:
		w.inputPress[key] = false
	}

	w.setModKey(mods)
}

func (w *Window) cursorPositionCallback(window *glfw.Window, x float64, y float64) {
	w.cursorPos.X = float32(x)
	w.cursorPos.Y = float32(y)
}

func (w *Window) mouseButtonCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	switch action {
	case glfw.Repeat:
	case glfw.Press:
		w.inputPress[w.glfwMouseToInput(button)] = true
		break
	default:
		w.inputPress[w.glfwMouseToInput(button)] = false
	}

	w.setModKey(mods)
}

func (w *Window) glfwMouseToInput(button glfw.MouseButton) int {
	switch button {
	case glfw.MouseButtonLeft:
		return MouseButtonLeft
	case glfw.MouseButtonMiddle:
		return MouseButtonMiddle
	case glfw.MouseButtonRight:
		return MouseButtonRight
	}
	return InputUnknown
}

func (w *Window) setModKey(mods glfw.ModifierKey) {
	w.inputPress[KeyAlt] = false
	w.inputPress[KeyControl] = false
	w.inputPress[KeyShift] = false
	w.inputPress[KeySuper] = false

	switch mods {
	case glfw.ModAlt:
		w.inputPress[KeyAlt] = true
		break
	case glfw.ModControl:
		w.inputPress[KeyControl] = true
		break
	case glfw.ModShift:
		w.inputPress[KeyShift] = true
		break
	case glfw.ModSuper:
		w.inputPress[KeySuper] = true
		break
	}
}

func (w *Window) set2DProjection() {
	gl.Disable(gl.DEPTH_TEST)
	gl.Enable(gl.TEXTURE)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, float64(w.Width), float64(w.Height), 0, 0, 1)

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
}
