package main

import (
	"fmt"
)

type View struct {
	model *Model
}

func (v *View) NotifyUpdate(ev Event) {
	v.drawScreen()

	switch e := ev.(type) {
	case *OnFightEnd:
		fmt.Printf("%s のこうげき!\n", e.Subject.Name)
		fmt.Printf("%s に %d のダメージ!\n", e.Target.Name, e.Damage)

		switch e.Result {
		case FightPlayerWin:
			fmt.Printf("%s をたおした!\n", e.Target.Name)
		case FightPlayerLose:
			fmt.Printf("あなたは まけました\n")
		}
	case *OnSpellEnd:
		fmt.Printf("%sは ヒールを となえた!\n", e.Subject.Name)
		switch e.Result {
		case SpellFailure:
			fmt.Printf("MPが たりない!\n")
		case SpellSuccess:
			fmt.Printf("%sのきずが かいふくした\n", e.Subject.Name)
		}
	case *OnRunEnd:
		fmt.Printf("%sは にげだした!\n", e.Subject.Name)
	}
}

func (v *View) drawScreen() {
	fmt.Print("\033[H\033[2J")

	p := v.model.Player()
	fmt.Printf("%s\n", p.Name)
	fmt.Printf("HP: %d/%d MP: %d/%d\n", p.HP, p.MaxHP, p.MP, p.MaxMP)

	m := v.model.Monster()
	fmt.Printf("\n%s(HP:%d/%d)\n", m.AA, m.HP, m.MaxHP)
	fmt.Println()
}

func (v *View) initScreen() {
	v.drawScreen()
	m := v.model.Monster()
	fmt.Printf("%s があらわれた\n", m.Name)
}

var commandNames = []string{
	"たたかう",
	"じゅもん",
	"にげる",
}

func (v *View) drawSelectCommand(cmd Command) {
	v.drawScreen()

	for i, n := range commandNames {
		if i == cmd {
			fmt.Print(">")
		} else {
			fmt.Print(" ")
		}
		fmt.Printf("%s\n", n)
	}
}
