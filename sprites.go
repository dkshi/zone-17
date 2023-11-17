package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprites struct {
	collection *map[string]*ebiten.Image
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
	tile1, _, err := ebitenutil.NewImageFromFile("assets/tile1.png")
	if err != nil {
		return &Sprites{}, err
	}

	newCollection["background"] = background
	newCollection["playerModel"] = playerModel
	newCollection["tile1"] = tile1

	return &Sprites{collection: &newCollection}, nil
}
