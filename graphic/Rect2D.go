package graphic

import "github.com/micahke/mango/opengl"

type Rect2D struct {
	*Mesh // the eobmbedded Mesh object

	// POSITION:
	x, y   float32
	width  float32
	height float32
}

// Return a new Rect2D
func NewRect2D(x, y, width, height float32) *Rect2D {
	rect2d := new(Rect2D)

  // SUPER: build mesh
  mesh := new(Mesh)
  mesh.vao = opengl.NewVertexArray()
  mesh.vbo = opengl.NewVertexBuffer(rect2d.GetVertices())
  mesh.vbl = opengl.NewVertexBufferLayout()

  mesh.vbl.Pushf(3)
  mesh.vao.AddBuffer(*mesh.vbo, *mesh.vbl)

  mesh.ibo = opengl.NewIndexBuffer(rect2d.GetIndeces())

  rect2d.Mesh = mesh
	rect2d.x = x
	rect2d.y = y
	rect2d.width = width
	rect2d.height = height

	return rect2d
}

func (rect *Rect2D) Render() {

}

// get the vertices for a quad
func (rect *Rect2D) GetVertices() []float32 {
	vertices := []float32{
		0.5, 0.5, 0.0,
		0.5, -0.5, 0.0,
		-0.5, -0.5, 0.0,
		-0.5, 0.5, 0.0,
	}
  return vertices
}

// get the indeces for a quad
func (rect *Rect2D) GetIndeces() []uint32 {
	indeces := []uint32{
    0, 1, 3,
    1, 2, 3,
	}
  return indeces
}
