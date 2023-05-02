package main

import (
	"codecav.es/graphboxes/lib"
	"github.com/veandco/go-sdl2/sdl"
)

type Box struct {
	color sdl.Color
	width int
	height int
	x float32
	y float32
}

func (box *Box) display(renderer *sdl.Renderer) {
	var rect sdl.Rect

	rect.X = int32(box.x)
	rect.Y = int32(box.y)
	rect.H = int32(box.height)
	rect.W = int32(box.width)

	renderer.SetDrawColor(
		box.color.R,
		box.color.G,
		box.color.B,
		box.color.A,
	)
	renderer.FillRect(&rect)
}

func initEntities() []Box {
	var ents []Box

	ents = append(ents, Box{
		x: 50,
		y: 50,
		height: 25,
		width: 25,
		color: sdl.Color{R: 255, G: 0, B: 0, A: 255},
	})

	return ents
}

func updateEntities(entities []Box) {
	;
}

func drawEntitites(entities []Box, r *sdl.Renderer) {
	for _, ent := range entities {
		ent.display(r)
	}
}

func clearScreen(r *sdl.Renderer) {
	// Fill the surface with black.
	r.SetDrawColor(0, 0, 0, 255)
	r.Clear()
}

func main() {
	
	// Initialize the SDL2 subsystems.
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	// Create the Window.
	win, err := sdl.CreateWindow(
		"Hello",
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		lib.SCREEN_WIDTH,
		lib.SCREEN_HEIGHT,
		sdl.WINDOW_SHOWN,
	)

	if err != nil {
		panic(err)
	}
	defer win.Destroy()

	// Create an SDL Renderer so we can draw to it.
	renderer, err := sdl.CreateRenderer(win, 0, 0)
	if err != nil {
		panic(err)
	}

	// Event loop.
	running := true

	// Initialize objects.
	ents := initEntities()

	EventLoop:
	for running {

		// clearScreen
		clearScreen(renderer)

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Goodbye!")
				running = false
				break EventLoop
			}
		}

		updateEntities(ents)
		drawEntitites(ents, renderer)

		renderer.Present()
	}
}