package main

type Direction int

const (
	North Direction = iota
	South Direction
	East  Direction
	West  Direction
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

func main() {

}
