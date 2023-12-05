package game

import (
	zone17 "github.com/dkshi/zone-17"
	"github.com/dkshi/zone-17/internal/world"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 360
)

type Game struct {
	World      *world.World
	Controller *Controller
	Renderer   *Renderer
	Sprites    *zone17.Sprites
}

func NewGame(world *world.World, sprites *zone17.Sprites) *Game {
	return &Game{
		World:      world,
		Controller: NewController(),
		Renderer:   NewRenderer(),
		Sprites: sprites,
	}
}

func (g *Game) Update() error {
	g.Controller.PlayerController(g)
	g.Controller.FloorController(g)
	g.Controller.CeilingController(g)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Renderer.RenderBackground(g, screen)
	g.Renderer.RenderTiles(g, screen)
	g.Renderer.RenderCeiling(g, screen)
	g.Renderer.RenderPlayer(g, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
