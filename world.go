package main

import "github.com/hajimehoshi/ebiten/v2"

type World struct {
	background *ebiten.Image
	player     *GameObject
	tiles      *[]GameObject
}

func NewWorld(background *ebiten.Image, player *GameObject, tiles *[]GameObject) *World {
	return &World{
		background: background,
		player:     player,
		tiles:      tiles,
	}
}
