package main

import "github.com/hajimehoshi/ebiten/v2"

type World struct {
	background *ebiten.Image
	player     *GameObject
	floor      *Floor
	ceiling    *Ceiling
}

func NewWorld(background *ebiten.Image, player *GameObject, floor *Floor, ceiling *Ceiling) *World {
	return &World{
		background: background,
		player:     player,
		floor:      floor,
		ceiling:    ceiling,
	}
}
