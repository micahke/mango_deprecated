package window

import (
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type GLFW_Instance struct {

	// PUBLIC:
	Major int
	Minor int
}

var glfwInstance *GLFW_Instance

func GetGLFWInstance() *GLFW_Instance {
	return glfwInstance
}

func InitGLFW(major int, minor int) {

  runtime.LockOSThread()

	// GLFW:
	glfwInstance = new(GLFW_Instance)
	glfwInstance.Major = major
	glfwInstance.Minor = minor

	if err := glfw.Init(); err != nil {
		panic("Error initializing GLFW")
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

}

func InitGL() {
	// OPENGL:
	if err := gl.Init(); err != nil {
		panic("Error initializing OpenGL")
	}
}


func PollEvents() {
  glfw.PollEvents()
}
