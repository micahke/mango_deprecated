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
  window := window.NewWindow(800, 600, "Hello, world!")


  mainScene := window.CreateScene()

  rect := graphic.NewRect2D(100, 100, 100, 100)
  gameObject := scene.NewGameObjectM(rect.Mesh)

  rect2 := graphic.NewRect2D(400, 400, 50, 50)
  gameObject2 := scene.NewGameObjectM(rect2.Mesh)

  mainScene.AddGameObject(gameObject)
  mainScene.AddGameObject(gameObject2)
  mainScene.InitCamera()
  
  mango.SetBackgroundColor(0.5, 0.5, 0.5)

  // Attack the window and start the program
  mango.AttachWindow(window)
  mango.Start()
}
