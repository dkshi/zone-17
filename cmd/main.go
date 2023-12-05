package main

import (
	"log"

	zone17 "github.com/dkshi/zone-17"
	"github.com/dkshi/zone-17/internal/game"
	"github.com/dkshi/zone-17/internal/surroudings"
	"github.com/dkshi/zone-17/internal/world"
	"github.com/hajimehoshi/ebiten/v2"
)

var sprites *zone17.Sprites

const (
	screenWidth     = 640
	screenHeight    = 360
	playerStartPosX = 40
	playerStartPosY = 200
	playerWidth     = 20
	playerHeight    = 31
)

func main() {
	var err error

	sprites, err = zone17.NewSprites()
	if err != nil {
		log.Fatal(err)
	}

	player := zone17.NewGameObject(sprites.Collection["playerModel"], playerStartPosX, playerStartPosY, playerWidth, playerHeight)

	floor := surroudings.InitFloor(sprites)
	ceiling := surroudings.InitCeiling(sprites)

	world := world.NewWorld(sprites.Collection["background"], player, floor, ceiling)

	game := game.NewGame(world, sprites)

	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Zone 17")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
