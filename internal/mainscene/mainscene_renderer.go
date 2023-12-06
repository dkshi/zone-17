package mainscene

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type MainSceneRenderer struct {
	world *MainSceneWorld
}

func NewMainSceneRenderer(world *MainSceneWorld) *MainSceneRenderer {
	return &MainSceneRenderer{
		world: world,
	}
}

func (r *MainSceneRenderer) RenderBackground(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)
	screen.DrawImage(r.world.Background, op)
}

func (r *MainSceneRenderer) RenderPlayer(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(r.world.Player.PosX), float64(r.world.Player.PosY))
	screen.DrawImage(r.world.Player.Model, op)
}

func (r *MainSceneRenderer) RenderCeiling(screen *ebiten.Image) {
	opTop := &ebiten.DrawImageOptions{}
	opTop.GeoM.Translate(float64(r.world.Ceiling.Top.PosX), float64(r.world.Ceiling.Top.PosY))
	screen.DrawImage(r.world.Ceiling.Top.Model, opTop)

	for _, lamp := range r.world.Ceiling.Lamps {
		opLamp := &ebiten.DrawImageOptions{}
		opLamp.GeoM.Translate(float64(lamp.PosX), float64(lamp.PosY))
		screen.DrawImage(lamp.Model, opLamp)
	}

}

func (r *MainSceneRenderer) RenderTiles(screen *ebiten.Image) {
	for _, tile := range r.world.Floor.Tiles {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(tile.PosX), float64(tile.PosY))
		screen.DrawImage(tile.Model, op)
	}
}
