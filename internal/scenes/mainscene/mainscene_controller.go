package mainscene

import (
	"github.com/dkshi/zone-17/internal/surroudings"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 640
	screenHeight = 360
	jumpForce    = 5
	jumpHeight   = 120
	biasY        = 5
	tileSpeedX   = 5
	lampSpeedX   = 5
	gReverseInc  = 5
	playerSpeed  = 10
)

var (
	gravity     = 8
	isMoving    = true
	isFalling   = false
	isGrounded  = false
	lastTileY   = 0
	tickCounter = 0
)

type MainSceneController struct {
	world *MainSceneWorld
}

func NewMainSceneController(world *MainSceneWorld) *MainSceneController {
	return &MainSceneController{
		world: world,
	}
}

func (c *MainSceneController) TickController() {
	tickCounter++
}

func (c *MainSceneController) PlayerController() {
	if isMoving && tickCounter%playerSpeed == 0 {
		c.world.points++
		
	}

	if isFalling {
		c.world.player.PosY += gravity
	}

	isGrounded = false

	if c.world.player.CollidingWithBottom(c.world.ceiling.Top) {
		c.world.player.PosY = c.world.ceiling.Top.PosY + c.world.ceiling.Top.Height - biasY
		isGrounded = true
	}

	for _, tile := range c.world.floor.Tiles {
		if c.world.player.CollidingWithSides(tile) && !c.world.player.CollidingWithTop(tile) {
			isMoving = false
			break
		}

		if c.world.player.CollidingWithTop(tile) {
			c.world.player.PosY = tile.PosY - c.world.player.Height + biasY
			lastTileY = screenHeight - tile.PosY
			isGrounded = true
		}
	}

	for _, lamp := range c.world.ceiling.Lamps {
		if c.world.player.CollidingWithBottom(lamp) || c.world.player.CollidingWithSides(lamp) {
			isMoving = false
			break
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
		c.world.player.PosY -= jumpForce
	}

	if c.world.player.PosY < screenHeight-(lastTileY+jumpHeight) && !isGrounded {
		isFalling = true
	}
}

func (c *MainSceneController) FloorController() {
	tiles := c.world.floor.Tiles
	if isMoving {
		for _, tile := range tiles {
			tile.PosX -= tileSpeedX
		}
	}

	if tiles[0].PosX <= -tiles[0].Width {
		tilesCount := len(c.world.floor.Tiles)
		for i := 0; i < tilesCount-1; i++ {
			tiles[i] = tiles[i+1]
		}
		tiles[tilesCount-1] = surroudings.CreateTile(c.world.sprites)
	}
}

func (c *MainSceneController) CeilingController() {
	lamps := c.world.ceiling.Lamps
	if isMoving {
		for _, lamp := range lamps {
			lamp.PosX -= lampSpeedX
		}
	}

	if lamps[0].PosX <= -lamps[0].Width {
		lampsCount := len(c.world.ceiling.Lamps)
		for i := 0; i < lampsCount-1; i++ {
			lamps[i] = lamps[i+1]
		}
		lamps[lampsCount-1] = surroudings.CreateLamp(c.world.sprites)
	}
}
