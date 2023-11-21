package main

import "github.com/hajimehoshi/ebiten/v2"

type GameObject struct {
	model  *ebiten.Image
	posX   int
	posY   int
	width  int
	height int
}

func NewGameObject(model *ebiten.Image, posX, posY, width, height int) *GameObject {
	return &GameObject{
		model:  model,
		posX:   posX,
		posY:   posY,
		width:  width,
		height: height,
	}
}

func (obj GameObject) isCollidingWithTop(other *GameObject) bool {
	return obj.posX < other.posX+other.width &&
		obj.posX+obj.width > other.posX &&
		obj.posY < other.posY &&
		obj.posY+obj.height > other.posY
}

func (obj GameObject) isCollidingWithSides(other *GameObject) bool {
	return obj.posX < other.posX+other.width &&
		obj.posX+obj.width > other.posX &&
		obj.posY < other.posY+other.height &&
		obj.posY+obj.height > other.posY
}

func (obj GameObject) isCollidingWithBottom(other *GameObject) bool {
	return obj.posX < other.posX+other.width &&
		obj.posX+obj.width > other.posX &&
		obj.posY < other.posY+other.height &&
		obj.posY+obj.height > other.posY+other.height
}