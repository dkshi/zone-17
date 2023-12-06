package mainscene

import (
	"github.com/dkshi/zone-17/internal/surroudings"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth     = 640
	screenHeight    = 360
	jumpForce   = 5
	jumpHeight  = 120
	biasY       = 5
	tileSpeedX  = 5
	lampSpeedX  = 5
	gReverseInc = 5
)

var (
	gravity       = 8
	isTilesMoving = true
	isLampsMoving = true
	isFalling     = false
	isGrounded    = false
	lastTileY     = 0
)

type MainSceneController struct {
	world *MainSceneWorld
}

func NewMainSceneController(world *MainSceneWorld) *MainSceneController {
	return &MainSceneController{
		world: world,
	}
}

func (c *MainSceneController) PlayerController() {
	if isFalling {
		c.world.Player.PosY += gravity
	}

	isGrounded = false

	if c.world.Player.CollidingWithBottom(c.world.Ceiling.Top) {
		c.world.Player.PosY = c.world.Ceiling.Top.PosY + c.world.Ceiling.Top.Height - biasY
		isGrounded = true
	}

	for _, tile := range c.world.Floor.Tiles {
		if c.world.Player.CollidingWithSides(tile) && !c.world.Player.CollidingWithTop(tile) {
			isTilesMoving = false
			isLampsMoving = false
			break
		}

		if c.world.Player.CollidingWithTop(tile) {
			c.world.Player.PosY = tile.PosY - c.world.Player.Height + biasY
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
		c.world.Player.PosY -= jumpForce
	}

	if c.world.Player.PosY < screenHeight-(lastTileY+jumpHeight) && !isGrounded {
		isFalling = true
	}
}

func (c *MainSceneController) FloorController() {
	tiles := c.world.Floor.Tiles
	if isTilesMoving {
		for _, tile := range tiles {
			tile.PosX -= tileSpeedX
		}
	}

	if tiles[0].PosX <= -tiles[0].Width {
		tilesCount := len(c.world.Floor.Tiles)
		for i := 0; i < tilesCount-1; i++ {
			tiles[i] = tiles[i+1]
		}
		tiles[tilesCount-1] = surroudings.CreateTile(c.world.Sprites)
	}
}

func (c *MainSceneController) CeilingController() {
	lamps := c.world.Ceiling.Lamps
	if isLampsMoving {
		for _, lamp := range lamps {
			lamp.PosX -= lampSpeedX
		}
	}

	if lamps[0].PosX <= -lamps[0].Width {
		lampsCount := len(c.world.Ceiling.Lamps)
		for i := 0; i < lampsCount-1; i++ {
			lamps[i] = lamps[i+1]
		}
		lamps[lampsCount-1] = surroudings.CreateLamp(c.world.Sprites)
	}
}
