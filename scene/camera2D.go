package scene

import glm "github.com/go-gl/mathgl/mgl32"


type Camera2D struct {
  projectionMatrix glm.Mat4
  viewMatrix glm.Mat4
}


func NewCamera2D(screenWidth, screenHeight float32) *Camera2D {
  camera := new(Camera2D)
  camera.projectionMatrix = glm.Ortho2D(0, screenWidth, 0, screenHeight)
  camera.viewMatrix = glm.Translate3D(0, 0, 0)
  return camera
}


// Get the projection matrix
func (camera *Camera2D) Proj2D() glm.Mat4 {
  return camera.projectionMatrix
}


// Get the view matrix
func (camera *Camera2D) View2D() glm.Mat4 {
  return camera.viewMatrix
}


