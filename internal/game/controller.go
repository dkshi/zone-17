package game

import (
	"github.com/dkshi/zone-17/internal/surroudings"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	jumpForce   = 5
	jumpHeight  = 120
	biasY       = 5
	tileSpeedX  = 5
	lampSpeedX  = 5
	gReverseInc = 5
)

var (
	gravity = 8
	isTilesMoving = true
	isLampsMoving = true
	isFalling = false
	isGrounded = false
	lastTileY = 0
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) PlayerController(g *Game) {
	if isFalling {
		g.World.Player.PosY += gravity
	}

	isGrounded = false

	if g.World.Player.CollidingWithBottom(g.World.Ceiling.Top) {
		g.World.Player.PosY = g.World.Ceiling.Top.PosY + g.World.Ceiling.Top.Height - biasY
		isGrounded = true
	}

	for _, tile := range g.World.Floor.Tiles {
		if g.World.Player.CollidingWithSides(tile) && !g.World.Player.CollidingWithTop(tile) {
			isTilesMoving = false
			isLampsMoving = false
			break
		}

		if g.World.Player.CollidingWithTop(tile) {
			g.World.Player.PosY = tile.PosY - g.World.Player.Height + biasY
			lastTileY = screenHeight - tile.PosY
			isGrounded = true
		}
	}

	if isGrounded && inpututil.IsKeyJustPressed(ebiten.KeyG) {
		gravity = (gravity * -1) - gReverseInc
	}

	if isGrounded && inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		isGrounded = false
		isFalling = false
	}

	if !isGrounded && !isFalling {
		g.World.Player.PosY -= jumpForce
	}

	if g.World.Player.PosY < screenHeight-(lastTileY+jumpHeight) && !isGrounded {
		isFalling = true
	}
}

func (c *Controller) FloorController(g *Game) {
	tiles := g.World.Floor.Tiles
	if isTilesMoving {
		for _, tile := range tiles {
			tile.PosX -= tileSpeedX
		}
	}

	if tiles[0].PosX <= -tiles[0].Width {
		tilesCount := len(g.World.Floor.Tiles)
		for i := 0; i < tilesCount-1; i++ {
			tiles[i] = tiles[i+1]
		}
		tiles[tilesCount-1] = surroudings.CreateTile(g.Sprites)
	}
}

func (c *Controller) CeilingController(g *Game) {
	lamps := g.World.Ceiling.Lamps
	if isLampsMoving {
		for _, lamp := range lamps {
			lamp.PosX -= lampSpeedX
		}
	}

	if lamps[0].PosX <= -lamps[0].Width {
		lampsCount := len(g.World.Ceiling.Lamps)
		for i := 0; i < lampsCount-1; i++ {
			lamps[i] = lamps[i+1]
		}
		lamps[lampsCount-1] = surroudings.CreateLamp(g.Sprites)
	}
}
