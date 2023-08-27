package main

import "github.com/veandco/go-sdl2/sdl"

type vector struct {
	x, y float64
}

type component interface {
	onTick() error
	onDraw(renderer *sdl.Renderer) error
	// onCollision(*element)
}

type element struct {
	position vector
	velocity vector
	rotation float64
	active   bool

	components []component
}

func (e *element) addComponent(c component) error {

	e.components = append(e.components, c)
	return nil

}
