package world

import (
	zone17 "github.com/dkshi/zone-17"
	"github.com/dkshi/zone-17/internal/surroudings"
	"github.com/hajimehoshi/ebiten/v2"
)

type World struct {
	Background *ebiten.Image
	Player     *zone17.GameObject
	Floor      *surroudings.Floor
	Ceiling    *surroudings.Ceiling
}

func NewWorld(background *ebiten.Image, player *zone17.GameObject, floor *surroudings.Floor, ceiling *surroudings.Ceiling) *World {
	return &World{
		Background: background,
		Player:     player,
		Floor:      floor,
		Ceiling:    ceiling,
	}
}
