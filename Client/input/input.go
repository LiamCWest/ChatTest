package input

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Input struct {
	Window *glfw.Window
}

func NewInput(window *glfw.Window) *Input {
	return &Input{Window: window}
}

func (input Input) keyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		window.SetShouldClose(true)
	}
}
