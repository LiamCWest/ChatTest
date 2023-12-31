package gui

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	serverApi "github.com/LiamCWest/ChatTest/Client/api"
	"github.com/LiamCWest/ChatTest/Client/input"
	utils "github.com/LiamCWest/ChatTest/api/v1/Utils"
	"github.com/go-gl/gl/v4.1-core/gl" // OR: github.com/go-gl/gl/v2.1/gl
	"github.com/go-gl/glfw/v3.2/glfw"
)

type GUI struct {
	Window *glfw.Window
	Input  *input.Input
}

const (
	width  = 1000
	height = 600
	speed  = 250

	vertexShaderSource = `
		#version 410
		in vec3 vp;
		void main() {
			gl_Position = vec4(vp, 1.0);
		}
	` + "\x00"

	fragmentShaderSource = `
		#version 410
		out vec4 frag_colour;
		void main() {
			frag_colour = vec4(0, 0.5, 0.5, 1);
		}
	` + "\x00"
)

func NewGUI(player *utils.Player) *GUI {
	runtime.LockOSThread()

	window := initGlfw()
	defer glfw.Terminate()

	gui := &GUI{Window: window, Input: input.NewInput(window)}
	gui.Input.Window.SetKeyCallback(input.KeyCallback)

	program := initOpenGL()

	API := serverApi.New()

	var (
		dt          float64 = 1.0 / 60.0
		accumulator float64 = 0
		currentTime float64 = glfw.GetTime()
	)

	for !window.ShouldClose() {

		newTime := glfw.GetTime()
		frameTime := newTime - currentTime
		currentTime = newTime

		accumulator += frameTime

		for accumulator >= dt {
			// move player
			API.MovePlayer(player.GetID(), player.GetVelocity().MultiplyScalar(float32(dt)*speed))

			accumulator -= dt
		}

		playersMessage := API.GetPlayers()
		players := make([]*utils.Player, 0, len(playersMessage))

		points := make([]float32, 0, len(playersMessage)*12)

		for _, playerMessage := range playersMessage {
			player := utils.NewPlayerFromMessage(playerMessage)
			player.GenGameObject()
			points = append(points, player.GetGameObject().PointsFromQuads()...)
			players = append(players, player)
		}

		glfw.PollEvents()

		input.CheckKeys(gui.Input.Window, player)

		draw(points, window, program)
	}

	API.RemovePlayer(player.GetID())

	return &GUI{Window: window, Input: gui.Input}
}

// initGlfw initializes glfw and returns a Window to use.
func initGlfw() *glfw.Window {
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

	return window
}

// initOpenGL initializes OpenGL and returns an intiialized program.
func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}
	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	return prog
}

func draw(points []float32, window *glfw.Window, program uint32) {
	for i := 0; i < len(points); i += 3 {
		vectorPoint := utils.NewVector2(points[i], points[i+1])
		vectorPoint = vectorPoint.Divide(utils.NewVector2(width, height)).MultiplyScalar(2).Subtract(utils.NewVector2(1, 1))
		points[i] = vectorPoint.X
		points[i+1] = vectorPoint.Y
	}
	vao := makeVao(points)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(points)/3))

	glfw.PollEvents()
	window.SwapBuffers()
}

// makeVao initializes and returns a vertex array from the points provided.
func makeVao(points []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
