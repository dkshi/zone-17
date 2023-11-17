package main

import "github.com/hajimehoshi/ebiten/v2"

func (g *Game) DrawBackground(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.25, 0.25)
	screen.DrawImage(g.world.background, op)
}

func (g *Game) DrawPlayer(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.world.player.posX), float64(g.world.player.posY))
	screen.DrawImage(g.world.player.model, op)
}

func (g *Game) DrawTiles(screen *ebiten.Image) {
	for _, tile := range *g.world.tiles {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(tile.posX), float64(tile.posY))
		screen.DrawImage(tile.model, op)
	}
}
