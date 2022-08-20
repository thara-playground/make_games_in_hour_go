package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/eiannone/keyboard"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := keyboard.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	t := time.NewTicker(1 * time.Second)
	defer t.Stop()

	initGame()

	ch := make(chan rune)
	go func() {
		for {
			char, _ := waitKey()
			select {
			case ch <- char:
			}
		}
	}()

	for {
		select {
		case <-t.C:
			fallBlock()
		case char := <-ch:
			last := block

			switch char {
			case 'w':
			case 's':
				block.y++
			case 'a':
				block.x--
			case 'd':
				block.x++
			default:
				rotateBlock()
			}

			if isCollided() {
				block = last
			} else {
				draw()
			}
		}

	}
}

const fieldWidth, fieldHeight = 12, 18

type blockType int

const (
	blockNone blockType = iota
	blockHard
	blockSoft
	blockFall
	blockMax
)

type Field [fieldHeight][fieldWidth]blockType

var field Field

var defaultField = Field{
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
}

func initGame() {
	initBlock()

	copyField(&field, &defaultField)
	draw()
}

func copyField(dst, src *Field) {
	for i, row := range src {
		copy(dst[i][:], row[:])
	}
}

func draw() {
	fmt.Print("\033[H\033[2J")

	var screen Field
	copyField(&screen, &field)
	for y := 0; y < blockHeightMax; y++ {
		for x := 0; x < blockWidthMax; x++ {
			if 0 < block.shape.pattern[y][x] {
				screen[block.y+y][block.x+x] = blockFall
			}
		}
	}

	for _, row := range screen {
		for _, block := range row {
			switch block {
			case blockNone:
				fmt.Print("  ")
			case blockHard:
				fmt.Print("* ")
			case blockSoft:
				fmt.Print("+ ")
			case blockFall:
				fmt.Print("@ ")
			}
		}
		fmt.Println()
	}
}

type blockShapeType int

const (
	blockShapeTypeI blockShapeType = iota
	blockShapeTypeL
	blockShapeTypeMax
)

const blockWidthMax, blockHeightMax = 4, 4

type blockShape struct {
	size    int
	pattern [blockWidthMax][blockHeightMax]int
}

var blockShapes = [blockShapeTypeMax]blockShape{
	{
		size: 3,
		pattern: [blockWidthMax][blockHeightMax]int{
			{0, 1, 0, 0},
			{0, 1, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 0, 0},
		},
	},
	{
		size: 3,
		pattern: [blockWidthMax][blockHeightMax]int{
			{0, 1, 0, 0},
			{0, 1, 1, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
	},
}

type fallingBlock struct {
	x, y  int
	shape blockShape
}

var block fallingBlock

func initBlock() {
	block.shape = blockShapes[rand.Intn(len(blockShapes))]
	block.x = fieldWidth/2 - block.shape.size/2
	block.y = 0

	rotateCount := rand.Intn(4)
	for i := 0; i < rotateCount; i++ {
		rotateBlock()
	}
}

func waitKey() (char rune, key keyboard.Key) {
	var err error
	char, key, err = keyboard.GetKey()
	if err != nil {
		panic(err)
	}
	if key == keyboard.KeyEsc {
		os.Exit(0)
	}
	return char, key
}

func rotateBlock() {
	rotated := block

	for y := 0; y < block.shape.size; y++ {
		for x := 0; x < block.shape.size; x++ {
			rotated.shape.pattern[block.shape.size-1-x][y] = block.shape.pattern[y][x]
		}
	}

	block = rotated
}

func fallBlock() {
	last := block
	block.y++
	if isCollided() {
		block = last
		for y := 0; y < blockHeightMax; y++ {
			for x := 0; x < blockWidthMax; x++ {
				if 0 < block.shape.pattern[y][x] {
					field[block.y+y][block.x+x] = blockSoft
				}
			}
		}
		eraceLine()

		initBlock()
		if isCollided() {
			initGame()
		}
	}
	draw()
}

func isCollided() bool {
	for y := 0; y < block.shape.size; y++ {
		for x := 0; x < block.shape.size; x++ {
			if 0 < block.shape.pattern[y][x] {
				globalX := block.x + x
				globalY := block.y + y
				if globalX < 0 || fieldWidth <= globalX ||
					globalY < 0 || fieldHeight <= globalY ||
					field[globalY][globalX] != blockNone {
					return true
				}
			}
		}
	}
	return false
}

func eraceLine() {
	for y := 0; y < fieldHeight; y++ {
		completed := true
		for x := 0; x < fieldWidth; x++ {
			if field[y][x] == blockNone {
				completed = false
				break
			}
		}
		if completed {
			for x := 0; x < fieldWidth; x++ {
				if field[y][x] == blockSoft {
					field[y][x] = blockNone
				}
			}
			for x := 0; x < fieldWidth; x++ {
				for y2 := y; 0 <= y2; y2-- {
					if field[y2][x] == blockHard {
						break
					}
					if y2 == 0 {
						field[y2][x] = blockNone
					} else if field[y2-1][x] != blockHard {
						field[y2][x] = field[y2-1][x]
					}
				}
			}
		}
	}
}
