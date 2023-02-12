package scene


type MComponent interface {
  Update()
  Draw()
}

type GameObject struct {
  // PUBLIC:


  // PRIVATE:
  id int
  components []MComponent

}

func (scene *Scene) NewGameObject() *GameObject {
  gameObject := new(GameObject) 
  scene.AddGameObject(*gameObject)
  return gameObject
}


