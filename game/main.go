package main

import (
	"github.com/micahke/mango/core"
	"github.com/micahke/mango/scene"
	"github.com/micahke/mango/window"
)



func main() {

  // Initialize the engine
  mango := core.NewMangoInstance()
  window := window.NewWindow(800, 600, "Hello, world!")


  mainScene := scene.NewScene()
  mainScene.NewGameObject()
  

  // Attack the window and start the program
  mango.AttachWindow(window)
  mango.Start()
}
