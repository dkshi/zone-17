package mainscene

import (
	zone17 "github.com/dkshi/zone-17"
	"github.com/dkshi/zone-17/internal/surroudings"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

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

func InitMainScene(sprites *zone17.Sprites, player *zone17.GameObject, standartFont *font.Face) (*MainScene, error) {
	floor := surroudings.InitFloor(sprites)
	ceiling := surroudings.InitCeiling(sprites)
	mainSceneWorld := NewMainSceneWorld(sprites.Collection["background"], player, floor, ceiling, sprites)
	mainSceneRenderer := NewMainSceneRenderer(mainSceneWorld, standartFont)
	mainSceneController := NewMainSceneController(mainSceneWorld)

	return NewMainScene(mainSceneRenderer, mainSceneController), nil
}

func (sc *MainScene) Draw(screen *ebiten.Image) {
	sc.renderer.RenderBackground(screen)
	sc.renderer.RenderTiles(screen)
	sc.renderer.RenderCeiling(screen)
	sc.renderer.RenderPlayer(screen)
	sc.renderer.RenderPoints(screen)
}

func (sc *MainScene) Update() {
	sc.controller.TickController()
	sc.controller.PlayerController()
	sc.controller.FloorController()
	sc.controller.CeilingController()
}
