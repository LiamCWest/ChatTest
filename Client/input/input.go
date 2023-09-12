package input

import (
	utils "github.com/LiamCWest/ChatTest/api/v1/Utils"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Input struct {
	Window *glfw.Window
}

func NewInput(window *glfw.Window) *Input {
	return &Input{Window: window}
}

func KeyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		window.SetShouldClose(true)
	}
}

func MovementKeys(Window *glfw.Window, player *utils.Player) {
	if Window.GetKey(glfw.KeyW) == glfw.Press {
		player.SetVelocity(utils.NewVector2(player.GetVelocity().X, 1).Normalize())
	}
	if Window.GetKey(glfw.KeyW) == glfw.Release && player.GetVelocity().Y > 0 {
		player.SetVelocity(utils.NewVector2(player.GetVelocity().X, 0).Normalize())
	}
	if Window.GetKey(glfw.KeyA) == glfw.Press {
		player.SetVelocity(utils.NewVector2(-1, player.GetVelocity().Y).Normalize())
	}
	if Window.GetKey(glfw.KeyA) == glfw.Release && player.GetVelocity().X < 0 {
		player.SetVelocity(utils.NewVector2(0, player.GetVelocity().Y).Normalize())
	}
	if Window.GetKey(glfw.KeyS) == glfw.Press {
		player.SetVelocity(utils.NewVector2(player.GetVelocity().X, -1).Normalize())
	}
	if Window.GetKey(glfw.KeyS) == glfw.Release && player.GetVelocity().Y < 0 {
		player.SetVelocity(utils.NewVector2(player.GetVelocity().X, 0).Normalize())
	}
	if Window.GetKey(glfw.KeyD) == glfw.Press {
		player.SetVelocity(utils.NewVector2(1, player.GetVelocity().Y).Normalize())
	}
	if Window.GetKey(glfw.KeyD) == glfw.Release && player.GetVelocity().X > 0 {
		player.SetVelocity(utils.NewVector2(0, player.GetVelocity().Y).Normalize())
	}
}
