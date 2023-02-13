package graphic

import "github.com/micahke/mango/opengl"


type Mesh struct {
  vertices []float32
  shader *opengl.Shader
  vao *opengl.VertexArray
  vbo *opengl.VertexBuffer
  vbl *opengl.VertexBufferLayout
  ibo *opengl.IndexBuffer
}

type i_Mesh interface {
  Render()
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
