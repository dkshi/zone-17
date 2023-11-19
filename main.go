package main

import (
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

var sprites *Sprites

const (
	screenWidth  = 640
	screenHeight = 360
	playerStartPosX = 40
	playerStartPosY = 200
	playerWidth = 20
	playerHeight = 31
	tileHeightDelta = 32
	tileHeight = 216
	tileWidth = 354
	spawnTileBiasX = 10
	startTilesAmount = 2
	startTilePosX = 0
	startTileMidBiasX = 150
	startTilePosY = 250
)

type Game struct {
	world *World
	controller *Controller
	renderer *Renderer  
}

func NewGame(world *World) *Game {
	return &Game{
		world: world,
		controller: NewController(),
		renderer: NewRenderer(),
	}
}

func (g *Game) Update() error {
	g.controller.PlayerController(g)
	g.controller.TilesController(g)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderer.DrawBackground(g, screen)
	g.renderer.DrawTiles(g, screen)
	g.renderer.DrawCeiling(g, screen)
	g.renderer.DrawPlayer(g, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func CreateTile() *GameObject {
	tilePosY := screenHeight - (rand.Intn(tileHeightDelta) + 15)
	randSpawnBiasX := rand.Intn(spawnTileBiasX + 1) + spawnTileBiasX
	newTile := NewGameObject(sprites.collection["tile"], screenWidth + randSpawnBiasX, tilePosY, tileWidth, tileHeight)
	return newTile
}

func CreateStart() []*GameObject {
	tile1 := NewGameObject(sprites.collection["tile"], startTilePosX, startTilePosY, tileWidth, tileHeight)
	tile2 := NewGameObject(sprites.collection["tile"], tile1.width + startTileMidBiasX, startTilePosY, tileWidth, tileHeight)
	return []*GameObject{tile1, tile2}
}


func main() {
	var err error

	sprites, err = NewSprites()
	if err != nil {
		log.Fatal(err)
	}

	player := NewGameObject(sprites.collection["playerModel"], playerStartPosX, playerStartPosY, playerWidth, playerHeight)

	floor := NewFloor(CreateStart())

	top := NewGameObject(sprites.collection["ceiling-top"], 0, 0, screenWidth, sprites.collection["ceiling-top"].Bounds().Dy())
	ceiling := NewCeiling(top)

	world := NewWorld(sprites.collection["background"], player, floor, ceiling)

	game := NewGame(world)

	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Zone 17")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
