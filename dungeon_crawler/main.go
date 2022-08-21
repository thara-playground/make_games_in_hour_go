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
		fmt.Print("\033[H\033[2J")

		if g.isReachedGoal() {
			drawEnding(&g)
			waitKey()
			g.init()
			continue
		}

		draw3d(&g)
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

func draw3d(g *game) {
	var screen = `
         
         
         
         
         
         
         
         `
	screenBuffer := []rune(screen)

	for i := 0; i < locationMax; i++ {
		pos := g.player.pos.add(locations[g.player.direction][i])
		if !isInsideMaze(pos) {
			continue
		}
		for j := 0; j < directionMax; j++ {
			relatedDirection := (directionMax + j - g.player.direction) % directionMax
			if !g.maze.tiles[pos.y][pos.x].walls[j] {
				continue
			}
			if len(aaTable[i][relatedDirection]) == 0 {
				continue
			}
			for k := range screenBuffer {
				runes := []rune(aaTable[i][relatedDirection])
				if runes[k] != ' ' && runes[k] != '\n' {
					screenBuffer[k] = runes[k]
				}
			}
		}
	}

	for _, c := range screenBuffer {
		switch c {
		case ' ':
			fmt.Print("　")
		case '#':
			fmt.Print("　")
		case '_':
			fmt.Print("＿")
		case '|':
			fmt.Print("｜")
		case '/':
			fmt.Print("／")
		case 'L':
			fmt.Print("＼")
		default:
			fmt.Print(string(c))
		}
	}
	fmt.Println()
	fmt.Println()
}

func draw(g *game) {

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
			} else if g.isGoal(y, x) {
				floorAA = "G "
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

func drawEnding(g *game) {
	fmt.Printf("　＊　＊　ＣＯＮＧＲＡＴＵＬＡＴＩＯＮＳ　＊　＊\n" +
		"\n" +
		"　あなたはついに　でんせつのまよけを　てにいれた！\n" +
		"\n" +
		"　しかし、くらくをともにした　「なかま」という\n" +
		"かけがえのない　たからをてにした　あなたにとって、\n" +
		"まよけのかがやきも　いろあせて　みえるのであった…\n" +
		"\n" +
		"　　　　　　　〜　ＴＨＥ　ＥＮＤ　〜\n")
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
