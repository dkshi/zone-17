package mainscene

import "github.com/hajimehoshi/ebiten/v2"

type MainScene struct {
	renderer   *MainSceneRenderer
	controller *MainSceneController
}

func NewMainScene(renderer *MainSceneRenderer, controller *MainSceneController) *MainScene {
	return &MainScene{
		renderer:   renderer,
		controller: controller,
	}
}

func (sc *MainScene) Draw(screen *ebiten.Image) {
	sc.renderer.RenderBackground(screen)
	sc.renderer.RenderTiles(screen)
	sc.renderer.RenderCeiling(screen)
	sc.renderer.RenderPlayer(screen)
}

func (sc *MainScene) Update() {
	sc.controller.PlayerController()
	sc.controller.FloorController()
	sc.controller.CeilingController()
}
