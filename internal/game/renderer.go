package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Renderer struct {
}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (r *Renderer) RenderBackground(g *Game, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)
	screen.DrawImage(g.World.Background, op)
}

func (r *Renderer) RenderPlayer(g *Game, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.World.Player.PosX), float64(g.World.Player.PosY))
	screen.DrawImage(g.World.Player.Model, op)
}

func (r *Renderer) RenderCeiling(g *Game, screen *ebiten.Image) {
	opTop := &ebiten.DrawImageOptions{}
	opTop.GeoM.Translate(float64(g.World.Ceiling.Top.PosX), float64(g.World.Ceiling.Top.PosY))
	screen.DrawImage(g.World.Ceiling.Top.Model, opTop)

	for _, lamp := range g.World.Ceiling.Lamps {
		opLamp := &ebiten.DrawImageOptions{}
		opLamp.GeoM.Translate(float64(lamp.PosX), float64(lamp.PosY))
		screen.DrawImage(lamp.Model, opLamp)
	}

}

func (r *Renderer) RenderTiles(g *Game, screen *ebiten.Image) {
	for _, tile := range g.World.Floor.Tiles {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(tile.PosX), float64(tile.PosY))
		screen.DrawImage(tile.Model, op)
	}
}
