package main

import "github.com/hajimehoshi/ebiten/v2"

type GameObject struct {
	model *ebiten.Image
	posX  int
	posY  int
}

func NewGameObject(model *ebiten.Image, posX, posY int) *GameObject {
	return &GameObject{model: model, posX: posX, posY: posY}
}

func (obj GameObject) isColliding(other *GameObject) bool {
	return obj.posX < obj.posX+other.model.Bounds().Dx() &&
		obj.posX+obj.model.Bounds().Dx() > other.posX &&
		obj.posY < other.posY+other.model.Bounds().Dy() &&
		obj.posY+obj.model.Bounds().Dy() > other.posY
}
