package coord

type Coord struct {
	X, Y int
}

func New(x, y int) *Coord {
	return &Coord {x, y}
}