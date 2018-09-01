// aqrs runs the main backend of Aquarius.
package main

import "log"

type World struct {
	Size      Coordinates
	Blocks    map[Coordinates]Block
	Units     []struct{}
	Buildings []struct{}
}

func (w *World) Debug() string {
	return "TODO"
}

func (w *World) Generate() {
	// Implementation notes:
	// - Cellular automata to generate caves?
	// - Perlin noise to generate material deposits?
}

type Coordinates struct {
	X int
	Y int
}

type Block struct {
	Type BlockType
}

type BlockType int

const (
	BlockUnknown BlockType = iota
	BlockSolid
	BlockEmpty
)

func main() {
	world := World{}
	log.Println(world)
}
