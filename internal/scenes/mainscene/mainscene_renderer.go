package mainscene

import (
	"image/color"

	"github.com/dkshi/zone-17/pkg/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

const (
	textPosY = 70
)

type MainSceneRenderer struct {
	world    *MainSceneWorld
	textFont *font.Face
}

func NewMainSceneRenderer(world *MainSceneWorld, textFont *font.Face) *MainSceneRenderer {
	return &MainSceneRenderer{
		world:    world,
		textFont: textFont,
	}
}

func (r *MainSceneRenderer) RenderBackground(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)
	screen.DrawImage(r.world.background, op)
}

func (r *MainSceneRenderer) RenderPlayer(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(r.world.player.PosX), float64(r.world.player.PosY))
	screen.DrawImage(r.world.player.Model, op)
}

func (r *MainSceneRenderer) RenderCeiling(screen *ebiten.Image) {
	opTop := &ebiten.DrawImageOptions{}
	opTop.GeoM.Translate(float64(r.world.ceiling.Top.PosX), float64(r.world.ceiling.Top.PosY))
	screen.DrawImage(r.world.ceiling.Top.Model, opTop)

	for _, lamp := range r.world.ceiling.Lamps {
		opLamp := &ebiten.DrawImageOptions{}
		opLamp.GeoM.Translate(float64(lamp.PosX), float64(lamp.PosY))
		screen.DrawImage(lamp.Model, opLamp)
	}

}

func (r *MainSceneRenderer) RenderTiles(screen *ebiten.Image) {
	for _, tile := range r.world.floor.Tiles {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(tile.PosX), float64(tile.PosY))
		screen.DrawImage(tile.Model, op)
	}
}

func (r *MainSceneRenderer) RenderPoints(screen *ebiten.Image) {
	pointsString := utils.UInt64ToString(r.world.points)

	textWidth := font.MeasureString(*r.textFont, pointsString).Ceil()
	x := (screenWidth - textWidth) / 2
	y := textPosY

	text.Draw(screen, pointsString, *r.textFont, int(x), int(y), color.Black)
}
