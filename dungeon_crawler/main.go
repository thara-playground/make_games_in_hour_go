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

	var m maze
	m.init()

	for {
		draw(&m)
		waitKey()
	}
}

func draw(m *maze) {
	fmt.Print("\033[H\033[2J")

	for y := 0; y < mazeHeight; y++ {
		for x := 0; x < mazeWidth; x++ {
			if m.tiles[y][x].walls[directionNorth] {
				fmt.Print("+--+")
			} else {
				fmt.Print("+  +")
			}
		}
		fmt.Println()

		for x := 0; x < mazeWidth; x++ {
			floorAA := "  "
			if m.tiles[y][x].walls[directionWest] {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
			fmt.Print(floorAA)
			if m.tiles[y][x].walls[directionEast] {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()

		for x := 0; x < mazeWidth; x++ {
			if m.tiles[y][x].walls[directionSouth] {
				fmt.Print("+--+")
			} else {
				fmt.Print("+  +")
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
