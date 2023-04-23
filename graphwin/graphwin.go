package main

import "github.com/veandco/go-sdl2/sdl"

func main() {
	const (
		SCREEN_WIDTH int32 = 320
		SCREEN_HEIGHT int32 = 240
	)
	
	// Initialize the SDL2 subsystems.
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	// Create the Window.
	win, err := sdl.CreateWindow(
		"Hello",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		SCREEN_WIDTH,
		SCREEN_HEIGHT,
		sdl.WINDOW_SHOWN,
	)

	if err != nil {
		panic(err)
	}
	defer win.Destroy()

	// Get the surface so we can draw to it.
	surface, err := win.GetSurface()
	if err != nil {
		panic(err)
	}

	// Fill the surface with black.
	surface.FillRect(nil, 0)

	// Event loop.
	running := true

	EventLoop:
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Goodbye!")
				running = false
				break EventLoop
			}
		}
	}
}