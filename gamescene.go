package zone17

import "github.com/hajimehoshi/ebiten/v2"

type GameScene interface {
	Update()
	Draw(screen *ebiten.Image)
}
