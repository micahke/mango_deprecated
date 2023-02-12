package core

import "github.com/go-gl/glfw/v3.3/glfw"


type m_Timer struct  {

  // PRIVATE:
  deltaTime float64
  lastFrame float64
} 


var Timer *m_Timer


func InitTimer() {

  if Timer != nil {
    panic("Timer already initialized")
  } 

  Timer = new(m_Timer)
  Timer.deltaTime = 0
}

// Gets the game time
func GetTime() float64 {
  return glfw.GetTime()
}


func UpdateDeltaTime() {
  now := glfw.GetTime()
  Timer.deltaTime = now - Timer.lastFrame
  Timer.lastFrame = now
}

func DeltaTime() float64 {
  return Timer.deltaTime
}
