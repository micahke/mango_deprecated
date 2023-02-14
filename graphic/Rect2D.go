package graphic

import (
	"github.com/micahke/mango/opengl"
)

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
  mesh := NewMesh()
  mesh.vao = opengl.NewVertexArray()
  mesh.vbo = opengl.NewVertexBuffer(rect2d.GetVertices())
  mesh.vbl = opengl.NewVertexBufferLayout()
  mesh.shader = opengl.NewShader("vertex2D.glsl", "fragment2D.glsl")

  mesh.vbl.Pushf(2)
  mesh.vao.AddBuffer(*mesh.vbo, *mesh.vbl)

  mesh.ibo = opengl.NewIndexBuffer(rect2d.GetIndeces())

  rect2d.Mesh = mesh
	rect2d.x = x
	rect2d.y = y
	rect2d.width = width
	rect2d.height = height

	return rect2d
}


// get the vertices for a quad
func (rect *Rect2D) GetVertices() []float32 {
	vertices := []float32{
		100.0, 100.0, 
		200.0, 100.0, 
		200.0, 200.0, 
		100.0, 200.0, 
	}
  return vertices
}

// get the indeces for a quad
func (rect *Rect2D) GetIndeces() []uint32 {
	indeces := []uint32{
    0, 1, 2,
    2, 3, 0,
	}
  return indeces
}
