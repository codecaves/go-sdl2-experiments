package lib

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Entity interface {
	Display (renderer *sdl.Renderer)
}

type Box struct {
	Color sdl.Color
	Rect sdl.Rect
}

func (box *Box) Display(renderer *sdl.Renderer) {

	renderer.SetDrawColor(
		box.Color.R,
		box.Color.G,
		box.Color.B,
		box.Color.A,
	)
	renderer.FillRect(&box.Rect)
}