package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

type command = int

const (
	commandFight command = iota
	commandSpell
	commandRun
	commandMax
)

const spellCost = 3

var commandNames = []string{
	"たたかう",
	"じゅもん",
	"にげる",
}

func selectCommand(g *game) {
	player := g.player()
	player.cmd = commandFight

	for {
		drawScreen(g)
		for i, n := range commandNames {
			if i == player.cmd {
				fmt.Print(">")
			} else {
				fmt.Print(" ")
			}
			fmt.Printf("%s\n", n)
		}

		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if key == keyboard.KeyEnter {
			return
		}

		player.cmd = moveCursor(player.cmd, char)
	}
}

func moveCursor(c command, char rune) command {
	switch char {
	case 'w':
		c--
	case 's':
		c++
	default:
		return c
	}
	if c < 0 {
		c = commandMax - 1
	} else if commandMax <= c {
		c = 0
	}
	return c
}
