package main

import (
	"codecav.es/graphboxes/lib"
	"github.com/veandco/go-sdl2/sdl"
)

func initEntities() []lib.Box {
	var ents []lib.Box

	ents = append(ents, 
		lib.Box{
			Rect: sdl.Rect { 
				X: 50,
				Y: 50,
				H: 25,
				W: 25,
			},
			Color: sdl.Color{
				R: 255, 
				G: 0, 
				B: 0, 
				A: 255},
	})

	return ents
}

func updateEntities(entities []lib.Box) {
	;
}

func drawEntitites(entities []lib.Box, r *sdl.Renderer) {
	for _, ent := range entities {
		ent.Display(r)
	}
}

func clearScreen(r *sdl.Renderer) {
	// Fill the surface with black
	r.SetDrawColor(0, 0, 0, 255)
	r.Clear()
}

func main() {
	
	// Initialize the SDL2 subsystems
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	// Create the Window
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

	// Create an SDL Renderer so we can draw to it
	renderer, err := sdl.CreateRenderer(win, 0, 0)
	if err != nil {
		panic(err)
	}

	// Event loop
	running := true

	// Initialize objects
	ents := initEntities()

	EventLoop:
	for running {

		// Blank the screen
		clearScreen(renderer)

		// Poll for events
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Goodbye!")
				running = false
				break EventLoop
			}
		}

		// Update positions and states of all entities, and draw them to the renderer object
		updateEntities(ents)
		drawEntitites(ents, renderer)

		renderer.Present()
	}
}