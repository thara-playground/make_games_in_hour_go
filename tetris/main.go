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

	var g game
	g.init()
	draw(&g)

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
			g.fallBlock()
			draw(&g)
		case char := <-ch:
			last := g.block

			switch char {
			case 'w':
			case 's':
				g.block.y++
			case 'a':
				g.block.x--
			case 'd':
				g.block.x++
			default:
				g.rotateBlock()
			}

			if g.isBlockCollided() {
				g.block = last
			} else {
				draw(&g)
			}
		}

	}
}

func draw(g *game) {
	fmt.Print("\033[H\033[2J")

	var screen Field
	copyField(&screen, &g.field)
	for y := 0; y < blockHeightMax; y++ {
		for x := 0; x < blockWidthMax; x++ {
			if 0 < g.block.shape.pattern[y][x] {
				screen[g.block.y+y][g.block.x+x] = blockFall
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
