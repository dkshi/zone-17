package surroudings

import (
	"math/rand"

	zone17 "github.com/dkshi/zone-17"
)

var (
	tileHeightDelta   = 32
	tileHeight        = 500
	tileWidth         = 354
	spawnTileBiasX    = 10
	startTilesAmount  = 2
	startTilePosX     = 0
	startTilePosY     = 250
	startTileMidBiasX = 150
)

type Floor struct {
	Tiles []*zone17.GameObject
}

func NewFloor(tiles []*zone17.GameObject) *Floor {
	return &Floor{Tiles: tiles}
}

func CreateTile(sprites *zone17.Sprites) *zone17.GameObject {
	tilePosY := screenHeight - (rand.Intn(tileHeightDelta) + 15)
	randSpawnBiasX := rand.Intn(spawnTileBiasX+1) + spawnTileBiasX
	newTile := zone17.NewGameObject(sprites.Collection["tile"], screenWidth+randSpawnBiasX, tilePosY, tileWidth, tileHeight)
	return newTile
}

func InitFloor(sprites *zone17.Sprites) *Floor {
	tile1 := zone17.NewGameObject(sprites.Collection["tile"], startTilePosX, startTilePosY, tileWidth, tileHeight)
	tile2 := zone17.NewGameObject(sprites.Collection["tile"], tile1.Width+startTileMidBiasX, startTilePosY, tileWidth, tileHeight)
	floor := NewFloor([]*zone17.GameObject{tile1, tile2})
	return floor
}
