package main

import (
	"log"

	"github.com/dkshi/zone-17/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game, err := game.InitGame()
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Zone 17")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
