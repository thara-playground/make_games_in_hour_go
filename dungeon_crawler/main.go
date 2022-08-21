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

	var g game
	g.init()

	for {
		draw(&g)
		char, _ := waitKey()
		switch char {
		case 'w':
			g.movePlayerForward()
		case 's':
			g.turnPlayerAround()
		case 'a':
			g.turnPlayerLeft()
		case 'd':
			g.turnPlayerRight()
		}
	}
}

func draw(g *game) {
	fmt.Print("\033[H\033[2J")

	for y := 0; y < mazeHeight; y++ {
		for x := 0; x < mazeWidth; x++ {
			if g.getTile(y, x).walls[directionNorth] {
				fmt.Print("+--+")
			} else {
				fmt.Print("+  +")
			}
		}
		fmt.Println()

		for x := 0; x < mazeWidth; x++ {
			floorAA := "  "
			if g.getTile(y, x).walls[directionWest] {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}

			if g.isPlayerExist(y, x) {
				floorAA = playerDirections[g.player.direction]
			}

			fmt.Print(floorAA)
			if g.getTile(y, x).walls[directionEast] {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()

		for x := 0; x < mazeWidth; x++ {
			if g.getTile(y, x).walls[directionSouth] {
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

var playerDirections = [directionMax]string{
	"^^",
	"<<",
	"vv",
	">>",
}
