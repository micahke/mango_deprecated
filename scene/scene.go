package scene


type Scene struct {
  gameObjects []*GameObject
  camera *Camera2D

  width, height float32

}


// Constructor: returns a new blank scene
func NewScene(width, height float32) *Scene {
  scene := new(Scene)
  scene.width = width
  scene.height = height
  return scene
}

// Add a game object to the scene
func (scene *Scene) AddGameObject(gameObject *GameObject) {
  scene.gameObjects = append(scene.gameObjects, gameObject)
}


func (scene *Scene) GetGameObjects() []*GameObject {
  return scene.gameObjects
}

func (scene *Scene) InitCamera() {
  camera := NewCamera2D(scene.width, scene.height)
  scene.camera = camera
}


// render the scene
func (scene *Scene) Render() {

  for i := 0; i < len(scene.gameObjects); i++ {

    mesh := scene.gameObjects[i].GetMesh()
    mesh.Render(scene.camera.projectionMatrix)
  }

}
