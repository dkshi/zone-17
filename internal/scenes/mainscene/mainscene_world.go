package mainscene

import (
	zone17 "github.com/dkshi/zone-17"
	"github.com/dkshi/zone-17/internal/surroudings"
	"github.com/hajimehoshi/ebiten/v2"
)

type MainSceneWorld struct {
	background *ebiten.Image
	player     *zone17.GameObject
	floor      *surroudings.Floor
	ceiling    *surroudings.Ceiling
	sprites    *zone17.Sprites
	points     uint64
}

func NewMainSceneWorld(background *ebiten.Image, player *zone17.GameObject, floor *surroudings.Floor, ceiling *surroudings.Ceiling, sprites *zone17.Sprites) *MainSceneWorld {
	return &MainSceneWorld{
		background: background,
		player:     player,
		floor:      floor,
		ceiling:    ceiling,
		sprites:    sprites,
	}
}
