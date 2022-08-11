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
	g.player().cmd = commandFight

	for {
		draw(g)
		for i, n := range commandNames {
			if i == g.player().cmd {
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

		c := g.player().cmd
		switch char {
		case 'w':
			c--
		case 's':
			c++
		default:
			return
		}
		if c < 0 {
			c = commandMax - 1
		} else if commandMax <= c {
			c = 0
		}
		g.player().cmd = c
	}
}
