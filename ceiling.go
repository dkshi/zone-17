package main

type Ceiling struct {
	top   *GameObject
	lamps []*GameObject
}

func NewCeiling(top *GameObject, lamps []*GameObject) *Ceiling {
	return &Ceiling{
		top:   top,
		lamps: lamps,
	}
}
