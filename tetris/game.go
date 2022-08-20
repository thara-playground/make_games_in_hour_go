package main

import "math/rand"

type blockType int

const (
	blockNone blockType = iota
	blockHard
	blockSoft
	blockFall
	blockMax
)

type game struct {
	field Field

	block struct {
		x, y  int
		shape blockShape
	}
}

func (g *game) init() {
	g.initBlock()

	copyField(&g.field, &defaultField)
}

func (g *game) initBlock() {
	g.block.shape = blockShapes[rand.Intn(len(blockShapes))]
	g.block.x = fieldWidth/2 - g.block.shape.size/2
	g.block.y = 0

	rotateCount := rand.Intn(4)
	for i := 0; i < rotateCount; i++ {
		g.rotateBlock()
	}
}

func (g *game) rotateBlock() {
	rotated := g.block

	for y := 0; y < g.block.shape.size; y++ {
		for x := 0; x < g.block.shape.size; x++ {
			rotated.shape.pattern[g.block.shape.size-1-x][y] = g.block.shape.pattern[y][x]
		}
	}

	g.block = rotated
}

func (g *game) fallBlock() {
	last := g.block
	g.block.y++
	if g.isBlockCollided() {
		g.block = last
		for y := 0; y < blockHeightMax; y++ {
			for x := 0; x < blockWidthMax; x++ {
				if 0 < g.block.shape.pattern[y][x] {
					g.field[g.block.y+y][g.block.x+x] = blockSoft
				}
			}
		}
		g.eraceLine()

		g.initBlock()
		if g.isBlockCollided() {
			g.init()
		}
	}
}

func (g *game) isBlockCollided() bool {
	for y := 0; y < g.block.shape.size; y++ {
		for x := 0; x < g.block.shape.size; x++ {
			if 0 < g.block.shape.pattern[y][x] {
				globalX := g.block.x + x
				globalY := g.block.y + y
				if globalX < 0 || fieldWidth <= globalX ||
					globalY < 0 || fieldHeight <= globalY ||
					g.field[globalY][globalX] != blockNone {
					return true
				}
			}
		}
	}
	return false
}

func (g *game) eraceLine() {
	for y := 0; y < fieldHeight; y++ {
		completed := true
		for x := 0; x < fieldWidth; x++ {
			if g.field[y][x] == blockNone {
				completed = false
				break
			}
		}
		if completed {
			for x := 0; x < fieldWidth; x++ {
				if g.field[y][x] == blockSoft {
					g.field[y][x] = blockNone
				}
			}
			for x := 0; x < fieldWidth; x++ {
				for y2 := y; 0 <= y2; y2-- {
					if g.field[y2][x] == blockHard {
						break
					}
					if y2 == 0 {
						g.field[y2][x] = blockNone
					} else if g.field[y2-1][x] != blockHard {
						g.field[y2][x] = g.field[y2-1][x]
					}
				}
			}
		}
	}
}
