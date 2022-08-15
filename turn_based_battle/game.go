package main

import (
	"fmt"
	"math/rand"

	"github.com/eiannone/keyboard"
)

type game struct {
	s []*status
}

func (g *game) init() {
	g.s = make([]*status, statusCount)
	g.s[statusPlayer] = &status{character: characterConfig[characterPlayer]}
}

func (g *game) battle(m monster) {
	g.s[statusMonster] = &status{character: characterConfig[m]}

	g.s[statusPlayer].target = statusMonster
	g.s[statusMonster].target = statusPlayer

	drawScreen(g)
	fmt.Printf("%s があらわれた\n", g.s[1].name)

	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if key == keyboard.KeyEsc {
			return
		}
		selectCommand(g)
		for _, c := range g.s {
			target := g.s[c.target]

			drawScreen(g)
			switch c.cmd {
			case commandFight:
				dmg := 1 + rand.Int()%c.attack
				target.hp -= dmg
				if target.hp < 0 {
					target.hp = 0
				}
				fmt.Printf("%s のこうげき!\n", c.name)
				fmt.Printf("%s に %d のダメージ!\n", target.name, dmg)

				if target.hp <= 0 {
					switch c.target {
					case statusMonster:
						target.aa = ""
						drawScreen(g)
						fmt.Printf("%s をたおした!\n", target.name)
						return
					case statusPlayer:
						fmt.Printf("あなたは まけました\n")
						keyboard.GetKey()
						return
					}
				}

				keyboard.GetKey()
				break
			case commandSpell:
				fmt.Printf("%sは ヒールを となえた!\n", c.name)
				if c.mp < spellCost {
					fmt.Printf("MPが たりない!\n")
					keyboard.GetKey()
					break
				}
				c.hp = c.maxHP
				c.mp -= spellCost
				fmt.Printf("%sのきずが かいふくした\n", c.name)
				keyboard.GetKey()
				drawScreen(g)
				break
			case commandRun:
				fmt.Printf("%sは にげだした!\n", c.name)
				keyboard.GetKey()
				return
			}
		}
	}
}

func (g *game) player() *status {
	return g.s[statusPlayer]
}

func drawScreen(g *game) {
	fmt.Print("\033[H\033[2J")

	p := g.s[statusPlayer]
	fmt.Printf("%s\n", p.name)
	fmt.Printf("HP: %d/%d MP: %d/%d\n", p.hp, p.maxHP, p.mp, p.maxMP)

	m := g.s[statusMonster]
	fmt.Printf("\n%s(HP:%d/%d)\n", m.aa, m.hp, m.maxHP)
	fmt.Println()
}

const (
	statusPlayer int = iota
	statusMonster
	statusCount
)

type status struct {
	character

	cmd    command
	target int
}
