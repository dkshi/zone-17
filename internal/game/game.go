package game

import (
	zone17 "github.com/dkshi/zone-17"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 360
)

type Game struct {
	Sprites       *zone17.Sprites
	SceneOnScreen zone17.GameScene
}

func NewGame(sceneOnScreen zone17.GameScene) *Game {
	return &Game{
		SceneOnScreen: sceneOnScreen,
	}
}

func (g *Game) Update() error {
	g.SceneOnScreen.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.SceneOnScreen.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
