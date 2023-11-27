package main

import (
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

var sprites *Sprites

const (
	screenWidth       = 640
	screenHeight      = 360
	playerStartPosX   = 40
	playerStartPosY   = 200
	playerWidth       = 20
	playerHeight      = 31
	tileHeightDelta   = 32
	tileHeight        = 216
	tileWidth         = 354
	spawnTileBiasX    = 10
	startTilesAmount  = 2
	startTilePosX     = 0
	startTilePosY     = 250
	startTileMidBiasX = 150
	lampWidth         = 75
	lampHeight        = 19
	startLampPosY     = 8
	startLampPosX     = 150
	startLampMidBiasX = 130
	spawnLampBiasX    = 0
	startLampsAmount  = 2
)

type Game struct {
	world      *World
	controller *Controller
	renderer   *Renderer
}

func NewGame(world *World) *Game {
	return &Game{
		world:      world,
		controller: NewController(),
		renderer:   NewRenderer(),
	}
}

func (g *Game) Update() error {
	g.controller.PlayerController(g)
	g.controller.FloorController(g)
	g.controller.CeilingController(g)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderer.RenderBackground(g, screen)
	g.renderer.RenderTiles(g, screen)
	g.renderer.RenderCeiling(g, screen)
	g.renderer.RenderPlayer(g, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func CreateTile() *GameObject {
	tilePosY := screenHeight - (rand.Intn(tileHeightDelta) + 15)
	randSpawnBiasX := rand.Intn(spawnTileBiasX+1) + spawnTileBiasX
	newTile := NewGameObject(sprites.collection["tile"], screenWidth+randSpawnBiasX, tilePosY, tileWidth, tileHeight)
	return newTile
}

func CreateLamp() *GameObject {
	randSpawnBiasX := rand.Intn(spawnLampBiasX+1) + spawnLampBiasX
	newLamp := NewGameObject(sprites.collection["lamp"], screenWidth+randSpawnBiasX, startLampPosY, tileWidth, tileHeight)
	return newLamp
}

func InitFloor() *Floor {
	tile1 := NewGameObject(sprites.collection["tile"], startTilePosX, startTilePosY, tileWidth, tileHeight)
	tile2 := NewGameObject(sprites.collection["tile"], tile1.width+startTileMidBiasX, startTilePosY, tileWidth, tileHeight)
	floor := NewFloor([]*GameObject{tile1, tile2})
	return floor
}

func InitCeling() *Ceiling {
	lamps := make([]*GameObject, startLampsAmount)
	spawnLampPosX := startLampPosX
	for i := 0; i < startLampsAmount; i++ {
		newLamp := NewGameObject(sprites.collection["lamp"], spawnLampPosX, startLampPosY, lampWidth, lampHeight)
		lamps[i] = newLamp
		spawnLampPosX += newLamp.width + startLampMidBiasX
	}

	top := NewGameObject(sprites.collection["ceiling-top"], 0, 0, screenWidth, sprites.collection["ceiling-top"].Bounds().Dy())
	ceiling := NewCeiling(top, lamps)
	return ceiling
}

func main() {
	var err error

	sprites, err = NewSprites()
	if err != nil {
		log.Fatal(err)
	}

	player := NewGameObject(sprites.collection["playerModel"], playerStartPosX, playerStartPosY, playerWidth, playerHeight)

	floor := InitFloor()
	ceiling := InitCeling()

	world := NewWorld(sprites.collection["background"], player, floor, ceiling)

	game := NewGame(world)

	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Zone 17")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
