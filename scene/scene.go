package scene


type Scene struct {
  gameObjects []GameObject
}


// Constructor: returns a new blank scene
func NewScene() *Scene {
  scene := new(Scene)
  return scene
}

// Add a game object to the scene
func (scene *Scene) AddGameObject(gameObject GameObject) {
  scene.gameObjects = append(scene.gameObjects, gameObject)
}


