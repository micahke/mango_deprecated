package scene

type GameObject struct {

	// PRIVATE:
	x float32
	y float32
	z float32
}

func NewGameObject() *GameObject {
	gameObject := new(GameObject)

	return gameObject

}

// Get the x posiiton of the game object
func (gameObject *GameObject) X() float32 {
  return gameObject.x
}

// Get the y posiiton of the game object
func (gameObject *GameObject) Y() float32 {
  return gameObject.y
}


// Get the z posiiton of the game object
func (gameObject *GameObject) Z() float32 {
  return gameObject.z
}
