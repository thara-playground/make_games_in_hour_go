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

	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

	ch := make(chan rune)
	defer close(ch)
	go func() {
		for {
			char, _ := waitKey()
			select {
			case ch <- char:
			}
		}
	}()

start:
	var g game
	g.init()

	for {
		select {
		case <-t.C:
			g.update()
			draw(&g)

			if g.isGameOver() {
				drawGameOver()
				<-ch
				goto start
			}
		case char := <-ch:
			newPos := g.characters[characterPlayer].currentPos

			switch char {
			case 'w':
				newPos.y--
			case 's':
				newPos.y++
			case 'a':
				newPos.x--
			case 'd':
				newPos.x++
			}

			g.updatePlayerPos(newPos)

			if g.isGameOver() {
				drawGameOver()
				<-ch
				goto start
			}
			if g.isCompleted() {
				drawCompletion()
				<-ch
				goto start
			}

			draw(&g)
		}
	}
}

func draw(g *game) {
	fmt.Print("\033[H\033[2J")

	var screen maze
	copyMaze(&screen, &g.maze)

	for i, character := range g.characters {
		screen[character.currentPos.y][character.currentPos.x] = rune(i)
	}

	for y := 0; y < mazeHeight; y++ {
		for x := 0; x < mazeWidth; x++ {
			switch screen[y][x] {
			case ' ':
				fmt.Print(" ")
			case '#':
				fmt.Print("■")
			case 'o':
				fmt.Print(".")
			case rune(characterPlayer):
				fmt.Print("o")
			case rune(characterRandom):
				fmt.Print("*")
			case rune(characterChase):
				fmt.Print("@")
			case rune(characterAmbush):
				fmt.Print("%")
			case rune(characterSiege):
				fmt.Print("$")
			}
		}
		fmt.Println()
	}
}

func drawGameOver() {
	fmt.Print("\033[H\033[2J")
	for i := 0; i < mazeHeight/2; i++ {
		fmt.Println()
	}
	fmt.Println("GAME OVER")
}

func drawCompletion() {
	fmt.Print("\033[H\033[2J")
	for i := 0; i < mazeHeight/2; i++ {
		fmt.Println()
	}
	fmt.Println("    CONGRATULATION!")
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
