package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/eiannone/keyboard"
)

var turnNames = []string{"", "黒", "白"}

var colors = []string{
	"_",
	"*",
	"@",
}

var cursorPos point

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

	mode := selectMode()

start:
	r := newReversi(mode)
	cursorPos = point{x: 3, y: 3}

	for {
		if !r.canPlaceAll() {
			r.skipTurn()
			if !r.canPlaceAll() {
				r.finish()
				draw(r)
				waitKey()
				goto start
			}
			continue
		}

		var pos point
		if r.isPlayerTurn() {
			pos = inputPos(r)
		} else {
			var ok bool
			pos, ok = selectPoint(r)
			if !ok {
				continue
			}
			draw(r)
			waitKey()
		}
		if !r.canPlace(pos, true) {
			fmt.Println("そこには置けません")
			waitKey()
			continue
		}
		r.place(pos)
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

func draw(r *reversi) {
	fmt.Print("\033[H\033[2J")

	for y := 0; y < boardHeight; y++ {
		for x := 0; x < boardWidth; x++ {
			fmt.Printf(" %s", colors[r.color(y, x)])
		}

		if r.isPlayerTurn() && y == cursorPos.y {
			fmt.Print(" ←")
		}

		fmt.Println()
	}
	for x := 0; x < boardWidth; x++ {
		if r.isPlayerTurn() && x == cursorPos.x {
			fmt.Print(" ↑")
		} else {
			fmt.Print("  ")
		}
	}
	fmt.Println()

	if r.turn != colorNone {
		fmt.Printf("%sのターンです\n", turnNames[r.turn])
	} else {
		nBlack := r.getCount(colorBlack)
		nWhite := r.getCount(colorWhite)

		var winner color
		if nWhite < nBlack {
			winner = colorBlack
		} else if nBlack < nWhite {
			winner = colorWhite
		} else {
			winner = colorNone
		}

		fmt.Printf("%s%d-%s%d\n",
			turnNames[colorBlack], nBlack,
			turnNames[colorWhite], nWhite)

		if winner == colorNone {
			fmt.Println("引き分け")
		} else {
			fmt.Printf("%sの勝ち\n", turnNames[winner])
		}
	}
}

func inputPos(r *reversi) point {
	for {
		draw(r)

		char, key := waitKey()
		if key == keyboard.KeyEnter {
			return cursorPos
		}
		switch char {
		case 'w':
			cursorPos.y--
			cursorPos.y = (cursorPos.y + boardWidth) % boardHeight
		case 'a':
			cursorPos.x--
			cursorPos.x = (cursorPos.x + boardWidth) % boardWidth
		case 's':
			cursorPos.y++
			cursorPos.y = (cursorPos.y + boardWidth) % boardHeight
		case 'd':
			cursorPos.x++
			cursorPos.x = (cursorPos.x + boardWidth) % boardWidth
		}
	}
}

var modeNames = [gameModeMax]string{
	"1P GAME",
	"2P GAME",
	"WATCH",
}

func selectMode() gameMode {
	var mode gameMode
	for {
		fmt.Print("\033[H\033[2J")
		fmt.Println("モードを選択して下さい")
		fmt.Print("\n\n")

		for i, name := range modeNames {
			if gameMode(i) == mode {
				fmt.Print("> ")
			} else {
				fmt.Print("  ")
			}
			fmt.Printf("%s\n", name)
			fmt.Println()
		}

		char, key := waitKey()
		if key == keyboard.KeyEnter {
			return mode
		}
		mode = moveCursor(mode, char)
	}
}

func moveCursor(m gameMode, char rune) gameMode {
	switch char {
	case 'w':
		m--
	case 's':
		m++
	default:
		return m
	}
	return (m + gameModeMax) % gameModeMax
}
