package window

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/micahke/mango/scene"
)

type Window struct {

	// PUBLIC:
	*glfw.Window // the actual embedded window

	// PRIVATE:

	width  int
	height int
	title  string

	scene *scene.Scene

}

var instance *Window

// Returns the instance for the current window
func GetInstance() *Window {
	return instance
}



// Constructor function for the window instance
func NewWindow(width int, height int, title string) *Window {
	// Check to see if we already have a window instance running
	if instance != nil {
		panic("Window instance already exists")
	}


	// Create the window
	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		panic("Error creating GLFW window")
	}

	instance = new(Window)
	instance.Window = window
	instance.width = width
	instance.height = height

	instance.Window.MakeContextCurrent()
	glfw.SwapInterval(1)

	InitGL()

	return instance

}

// attach the scene to the window
func (window *Window) CreateScene() *scene.Scene {
  ns := scene.NewScene(float32(window.width), float32(window.height))
  window.scene = ns
  return window.scene
}


// get the current scene within the window
func (window *Window) GetScene() *scene.Scene {
  return window.scene
}


func (window *Window) Width() int {
  return window.width
}

func (window *Window) Height() int {
  return window.height
}












