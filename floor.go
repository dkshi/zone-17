package main

type Floor struct {
	tiles []*GameObject
}

func NewFloor(tiles []*GameObject) *Floor {
	return &Floor{tiles: tiles}
}