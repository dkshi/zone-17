package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	jumpForce = 13
	gravity = 4
	jumpHeight = 30
	bias = 5
	screenWidth = 320
	screenHeight = 180
)

var isGrounded = true
var isFalling = true

type Game struct {
	world *World
}

func NewGame(world *World) *Game {
	return &Game{world: world}
}

func (g *Game) Update() error {
	g.world.player.posY += gravity 

	for _, tile := range *g.world.tiles {
		if tile.isColliding(g.world.player) {
			g.world.player.posY = tile.posY - g.world.player.model.Bounds().Dy() + bias
			isGrounded = true
		}
	}

	if isGrounded && ebiten.IsKeyPressed(ebiten.KeySpace) {
		isGrounded = false
		isFalling = false
	}

	if !isGrounded && !isFalling {
		g.world.player.posY -= jumpForce
	}

	if g.world.player.posY <= jumpHeight {
		isFalling = true
	}

	
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawBackground(screen)
	g.DrawTiles(screen)
	g.DrawPlayer(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	sprites, err := NewSprites()
	if err != nil {
		log.Fatal(err)
	}
	
	tile1 := NewGameObject((*sprites.collection)["tile1"], 0, 128)
	player := NewGameObject((*sprites.collection)["playerModel"], 41, 100)
	world := NewWorld((*sprites.collection)["background"], player, &[]GameObject{*tile1})
	game := NewGame(world)

	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Zone 17")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
