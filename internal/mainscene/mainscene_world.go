package mainscene

import (
	zone17 "github.com/dkshi/zone-17"
	"github.com/dkshi/zone-17/internal/surroudings"
	"github.com/hajimehoshi/ebiten/v2"
)

type MainSceneWorld struct {
	Background *ebiten.Image
	Player     *zone17.GameObject
	Floor      *surroudings.Floor
	Ceiling    *surroudings.Ceiling
	Sprites    *zone17.Sprites
}

func NewMainSceneWorld(background *ebiten.Image, player *zone17.GameObject, floor *surroudings.Floor, ceiling *surroudings.Ceiling, sprites *zone17.Sprites) *MainSceneWorld {
	return &MainSceneWorld{
		Background: background,
		Player:     player,
		Floor:      floor,
		Ceiling:    ceiling,
		Sprites:    sprites,
	}
}
