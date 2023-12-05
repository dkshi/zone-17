package zone17

import "github.com/hajimehoshi/ebiten/v2"

type GameObject struct {
	Model  *ebiten.Image
	PosX   int
	PosY   int
	Width  int
	Height int
}

func NewGameObject(model *ebiten.Image, posX, posY, width, height int) *GameObject {
	return &GameObject{
		Model:  model,
		PosX:   posX,
		PosY:   posY,
		Width:  width,
		Height: height,
	}
}

func (obj GameObject) CollidingWithTop(other *GameObject) bool {
	return obj.PosX < other.PosX+other.Width &&
		obj.PosX+obj.Width > other.PosX &&
		obj.PosY < other.PosY &&
		obj.PosY+obj.Height > other.PosY
}

func (obj GameObject) CollidingWithSides(other *GameObject) bool {
	return obj.PosX < other.PosX+other.Width &&
		obj.PosX+obj.Width > other.PosX &&
		obj.PosY < other.PosY+other.Height &&
		obj.PosY+obj.Height > other.PosY
}

func (obj GameObject) CollidingWithBottom(other *GameObject) bool {
	return obj.PosX < other.PosX+other.Width &&
		obj.PosX+obj.Width > other.PosX &&
		obj.PosY < other.PosY+other.Height &&
		obj.PosY+obj.Height > other.PosY+other.Height
}
