package surroudings

import (
	"math/rand"

	zone17 "github.com/dkshi/zone-17"
)

var (
	screenWidth       = 640
	screenHeight      = 360
	lampWidth         = 75
	lampHeight        = 19
	startLampPosY     = 8
	startLampPosX     = 150
	startLampMidBiasX = 130
	spawnLampBiasX    = 0
	startLampsAmount  = 2
)

type Ceiling struct {
	Top   *zone17.GameObject
	Lamps []*zone17.GameObject
}

func NewCeiling(top *zone17.GameObject, lamps []*zone17.GameObject) *Ceiling {
	return &Ceiling{
		Top:   top,
		Lamps: lamps,
	}
}

func InitCeiling(sprites *zone17.Sprites) *Ceiling {
	lamps := make([]*zone17.GameObject, startLampsAmount)
	spawnLampPosX := startLampPosX
	for i := 0; i < startLampsAmount; i++ {
		newLamp := zone17.NewGameObject(sprites.Collection["lamp"], spawnLampPosX, startLampPosY, lampWidth, lampHeight)
		lamps[i] = newLamp
		spawnLampPosX += newLamp.Width + startLampMidBiasX
	}

	top := zone17.NewGameObject(sprites.Collection["ceiling-top"], 0, 0, screenWidth, sprites.Collection["ceiling-top"].Bounds().Dy())
	ceiling := NewCeiling(top, lamps)
	return ceiling
}

func CreateLamp(sprites *zone17.Sprites) *zone17.GameObject {
	randSpawnBiasX := rand.Intn(spawnLampBiasX+1) + spawnLampBiasX
	newLamp := zone17.NewGameObject(sprites.Collection["lamp"], screenWidth+randSpawnBiasX, startLampPosY, lampWidth, lampHeight)
	return newLamp
}
