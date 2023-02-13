package core

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	glm "github.com/go-gl/mathgl/mgl32"
	"github.com/micahke/mango/window"
)

type MangoInstance struct {
	window *window.Window

  backgroundColor glm.Vec4

}

// Engine instance
var mango *MangoInstance


// Get the running engine instance
func GetInstance() *MangoInstance {
	return mango
}

// create a new mango instance and
func NewMangoInstance() *MangoInstance {
	if mango != nil {
		panic("Mango Instance already running...")
	}
	mango = new(MangoInstance)
	InitTimer()
	// initialize GLFW and OpenGL
	// Initialize GLFW and OpenGL
	window.InitGLFW(3, 3)

	return mango
}

// Attaches a window to the mango instance
func (instance *MangoInstance) AttachWindow(window *window.Window) {
	instance.window = window
}

// Main engine loop
func (instance *MangoInstance) Start() {

	var win = instance.window

	for !win.ShouldClose() {

		UpdateDeltaTime()

		if &instance.backgroundColor != nil {
      color := instance.backgroundColor
			gl.ClearColor(color[0], color[1], color[2], color[3])
		}
		gl.Clear(gl.COLOR_BUFFER_BIT)

		// scene := win.GetScene()
		// renderer.render(scene)

		win.SwapBuffers()
		window.PollEvents()

	}

}

func (instance *MangoInstance) SetBackgroundColor(r, g, b float32) {
  instance.backgroundColor = glm.Vec4{r, g, b, 1.0}
}
