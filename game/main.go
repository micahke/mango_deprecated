package main

import (
	"github.com/micahke/mango/core"
	"github.com/micahke/mango/graphic"
	"github.com/micahke/mango/scene"
	"github.com/micahke/mango/window"
)



func main() {

  // Initialize the engine
  mango := core.NewMangoInstance()
  window := window.NewWindow(850, 540, "Hello, world!")


  mainScene := window.CreateScene()

  rect := graphic.NewRect2D(0, 0, 100, 100)
  gameObject := scene.NewGameObjectM(rect.Mesh)

  mainScene.AddGameObject(gameObject)
  mainScene.InitCamera()
  
  mango.SetBackgroundColor(0.5, 0.5, 0.5)

  // Attack the window and start the program
  mango.AttachWindow(window)
  mango.Start()
}
