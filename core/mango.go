package core

import (

	"github.com/micahke/mango/window"
)

type MangoInstance struct {
  window *window.Window
}

// ENGINE INSTANCING
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

    scene := win.GetScene()
    // renderer.render(scene)
    
    win.SwapBuffers()
    window.PollEvents()

  }


}
