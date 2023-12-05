package zone17

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprites struct {
	Collection map[string]*ebiten.Image
}

func NewSprites() (*Sprites, error) {
	newCollection := make(map[string]*ebiten.Image)
	background, _, err := ebitenutil.NewImageFromFile("assets/background.png")
	if err != nil {
		return &Sprites{}, err
	}
	playerModel, _, err := ebitenutil.NewImageFromFile("assets/scp-096.png")
	if err != nil {
		return &Sprites{}, err
	}
	tile, _, err := ebitenutil.NewImageFromFile("assets/tile.png")
	if err != nil {
		return &Sprites{}, err
	}
	ceilingTop, _, err := ebitenutil.NewImageFromFile("assets/ceiling-top.png")
	if err != nil {
		return &Sprites{}, err
	}
	lamp, _, err := ebitenutil.NewImageFromFile("assets/lamp.png")
	if err != nil {
		return &Sprites{}, err
	}

	newCollection["background"] = background
	newCollection["playerModel"] = playerModel
	newCollection["tile"] = tile
	newCollection["ceiling-top"] = ceilingTop
	newCollection["lamp"] = lamp

	return &Sprites{Collection: newCollection}, nil
}
