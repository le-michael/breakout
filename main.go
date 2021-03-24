package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/le-michael/breakout/game"
	"github.com/le-michael/breakout/resmgr"
)

const (
	windowWidth  = 800
	windowHeight = 600
)

var breakout = game.New(windowWidth, windowHeight)

func init() {
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("Failed to initalize glfw:", err)
	}
	defer glfw.Terminate()

	//glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err := glfw.CreateWindow(windowWidth, windowHeight, "Breakout", nil, nil)
	if err != nil {
		log.Fatalln("Unable to create window:", err)
	}

	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		log.Fatalln("Unable to initalize opengl:", err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	window.SetKeyCallback(keyCallback)
	window.SetFramebufferSizeCallback(framebufferSizeCallback)

	gl.Viewport(0, 0, windowWidth, windowHeight)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	if err := breakout.Init(); err != nil {
		log.Fatalln("Unable to initalize breakout:", err)
	}

	var deltaTime, lastFrame float32
	deltaTime = 0
	lastFrame = 0

	for !window.ShouldClose() {
		currentFrame := float32(glfw.GetTime())
		deltaTime = currentFrame - lastFrame
		lastFrame = currentFrame
		glfw.PollEvents()

		breakout.ProcessInput(deltaTime)

		breakout.Update(deltaTime)

		gl.ClearColor(0, 0, 0, 1)
		gl.Clear(gl.COLOR_BUFFER_BIT)
		breakout.Render()

		window.SwapBuffers()
	}

	resmgr.Clear()
}

func keyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		window.SetShouldClose(true)
	}

	if key >= 0 && key < 1024 {
		if action == glfw.Press {
			breakout.Keys[key] = true
		} else if action == glfw.Release {
			breakout.Keys[key] = false
		}
	}
}

func framebufferSizeCallback(window *glfw.Window, width int, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
}
