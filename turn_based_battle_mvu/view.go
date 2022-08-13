package main

import "fmt"

func view(model Model) {
	fmt.Print("\033[H\033[2J")

	g := model.Game

	p := g.Player()
	fmt.Printf("%s\n", p.Name)
	fmt.Printf("HP: %d/%d MP: %d/%d\n", p.HP, p.MaxHP, p.MP, p.MaxMP)

	m := g.Monster()
	fmt.Printf("\n%s(HP:%d/%d)\n", m.AA, m.HP, m.MaxHP)
	fmt.Println()

	switch p := model.Phase.(type) {
	case *PhaseStart:
		m := g.Monster()
		fmt.Printf("%s があらわれた\n", m.Name)
	case *PhaseCommandSelect:
		switch model.Current {
		case statusPlayer:
			for i, n := range commandNames {
				if i == p.command {
					fmt.Print(">")
				} else {
					fmt.Print(" ")
				}
				fmt.Printf("%s\n", n)
			}
		case statusMonster:
		}
	case *PhaseBattleTurn:
		switch r := p.result.(type) {
		case *FightCommandResult:
			subject, target := g.status(model.Current), g.status(r.Target)
			fmt.Printf("%s のこうげき!\n", subject.Name)
			fmt.Printf("%s に %d のダメージ!\n", subject.Name, r.Damage)

			switch r.Result {
			case FightPlayerWin:
				fmt.Printf("%s をたおした!\n", target.Name)
			case FightPlayerLose:
				fmt.Printf("あなたは まけました\n")
			}
		case *SpellCommandResult:
			subject := g.status(r.Subject)
			fmt.Printf("%sは ヒールを となえた!\n", subject.Name)
			switch r.Result {
			case SpellFailure:
				fmt.Printf("MPが たりない!\n")
			case SpellSuccess:
				fmt.Printf("%sのきずが かいふくした\n", subject.Name)
			}
		case *RunCommandResult:
			subject := g.status(r.Subject)
			fmt.Printf("%sは にげだした!\n", subject.Name)
		}
	}
}

var commandNames = []string{
	"たたかう",
	"じゅもん",
	"にげる",
}
