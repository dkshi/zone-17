package main

type Ceiling struct {
	top *GameObject
}

func NewCeiling(top *GameObject) *Ceiling {
	return &Ceiling{top: top}
}
