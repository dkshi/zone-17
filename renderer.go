package main

import "github.com/hajimehoshi/ebiten/v2"

type Renderer struct {
	
}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (r *Renderer) DrawBackground(g *Game, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)
	screen.DrawImage(g.world.background, op)
}

func (r *Renderer) DrawPlayer(g *Game, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.world.player.posX), float64(g.world.player.posY))
	screen.DrawImage(g.world.player.model, op)
}

func (r *Renderer) DrawCeiling(g *Game, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.world.ceiling.top.posX), float64(g.world.ceiling.top.posY))
	screen.DrawImage(g.world.ceiling.top.model, op)
}

func (r *Renderer) DrawTiles(g *Game, screen *ebiten.Image) {
	for _, tile := range g.world.floor.tiles {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(tile.posX), float64(tile.posY))
		screen.DrawImage(tile.model, op)
	}
}
