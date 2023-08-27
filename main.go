package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	window, err := sdl.CreateWindow("Testing", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Printf("could not create window: %v", err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Printf("could not create renderer: %v", err)
	}
	defer renderer.Destroy()

}
