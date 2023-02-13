package graphic

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	glm "github.com/go-gl/mathgl/mgl32"
	"github.com/micahke/mango/opengl"
)

type Mesh struct {
	vertices []float32
	shader   *opengl.Shader
	vao      *opengl.VertexArray
	vbo      *opengl.VertexBuffer
	vbl      *opengl.VertexBufferLayout
	ibo      *opengl.IndexBuffer
  modelMatrix glm.Mat4
}

type i_Mesh interface {
	Render()
}

func NewMesh() *Mesh {
  mesh := new(Mesh)

  mesh.modelMatrix = glm.Translate3D(0, 0, 0)

  return mesh
}

func (mesh *Mesh) Render(projectionMatrix glm.Mat4, viewMatrix glm.Mat4) {
	if mesh.ibo != nil {
		mesh.ibo.Bind()
	}
	mesh.shader.Bind()
	mesh.vao.Bind()

  mesh.shader.SetUniformMat4f("model", mesh.modelMatrix)
  mesh.shader.SetUniformMat4f("view", viewMatrix)
  mesh.shader.SetUniformMat4f("proj", projectionMatrix)


	if mesh.ibo != nil {
		gl.DrawElements(gl.TRIANGLES, int32(mesh.ibo.GetCount()), gl.UNSIGNED_INT, nil)
	}
}

// NOTE: I think taking in a pointer here allows instancing
func (mesh *Mesh) SetVertices(vertices *[]float32) {
	mesh.vertices = *vertices
}

func (mesh *Mesh) SetShader(shader *opengl.Shader) {
	mesh.shader = shader
}

func (mesh *Mesh) SetVAO(vao *opengl.VertexArray) {
	mesh.vao = vao
}

func (mesh *Mesh) SetVBO(vbo *opengl.VertexBuffer) {
	mesh.vbo = vbo
}

func (mesh *Mesh) SetVBL(vbl *opengl.VertexBufferLayout) {
	mesh.vbl = vbl
}

func (mesh *Mesh) SetIBO(ibo *opengl.IndexBuffer) {
	mesh.ibo = ibo
}
