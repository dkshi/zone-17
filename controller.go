package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	jumpForce  = 5
	jumpHeight = 120
	biasY      = 5
	tileSpeedX = 5
)

var gravity = 8
var isTilesMoving = true
var isFalling = false
var isGrounded = false
var lastTileY = 0

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) PlayerController(g *Game) {
	if isFalling {
		g.world.player.posY += gravity
	}

	isGrounded = false

	if g.world.player.isCollidingWithBottom(g.world.ceiling.top) {
		g.world.player.posY = g.world.ceiling.top.posY + g.world.ceiling.top.height - biasY
		isGrounded = true
	}

	for _, tile := range g.world.floor.tiles {
		if g.world.player.isCollidingWithSides(tile) && !g.world.player.isCollidingWithTop(tile) {
			isTilesMoving = false
			break
		}

		if g.world.player.isCollidingWithTop(tile) {
			g.world.player.posY = tile.posY - g.world.player.height + biasY
			lastTileY = screenHeight - tile.posY
			isGrounded = true
		}
	}

	if isGrounded && inpututil.IsKeyJustPressed(ebiten.KeyG) {
		gravity *= -1
	}

	if isGrounded && inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		isGrounded = false
		isFalling = false
	}

	if !isGrounded && !isFalling {
		g.world.player.posY -= jumpForce
	}

	if g.world.player.posY < screenHeight-(lastTileY+jumpHeight) && !isGrounded {
		isFalling = true
	}
}

func (c *Controller) FloorController(g *Game) {
	tiles := g.world.floor.tiles
	if isTilesMoving {
		for _, tile := range tiles {
			tile.posX -= tileSpeedX
		}
	}

	if tiles[0].posX <= -tiles[0].width {
		tilesCount := len(g.world.floor.tiles)
		for i := 0; i < tilesCount-1; i++ {
			tiles[i] = tiles[i+1]
		}
		tiles[tilesCount-1] = CreateTile()
	}
}

func (c *Controller) CeilingController(g *Game) {
	
}
