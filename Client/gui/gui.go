package gui

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl" // OR: github.com/go-gl/gl/v2.1/gl
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	width  = 500
	height = 500
)

type GUI struct {
	window  *glfw.Window
	program uint32
}

func New() GUI {
	gui := GUI{}

	log.Println("Created GUI")

	return gui
}

func (gui GUI) Run() {
	runtime.LockOSThread()

	log.Printf("Running GUI")

	gui.initGlfw()
	defer glfw.Terminate()

	gui.initOpenGL()

	for !gui.window.ShouldClose() {
		gui.draw()
	}
}

func (gui GUI) draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(gui.program)

	glfw.PollEvents()
	gui.window.SwapBuffers()
}

// initGlfw initializes glfw and returns a Window to use.
func (gui GUI) initGlfw() {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4) // OR 2
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "ChatTest", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	gui.window = window
}

// initOpenGL initializes OpenGL and returns an intiialized program.
func (gui GUI) initOpenGL() {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	prog := gl.CreateProgram()
	gl.LinkProgram(prog)
	gui.program = prog
}
