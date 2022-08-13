package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/eiannone/keyboard"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := keyboard.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := runLoop(); err != nil {
		log.Fatal(err)
	}
}

func runLoop() error {
	game := NewGame(MonsterSlime)

	model := Model{Game: game, Phase: &PhaseStart{}}
	view(model)

	var msg Msg = &SelectCommand{}
	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			return fmt.Errorf("keyboard.GetKey() = : %#v", err)
		}
		if key == keyboard.KeyEsc {
			return nil
		}
		var next func() Msg
		model, next = update(msg, model)
		view(model)

		// next msg
		// switch m := msg.(type) {
		// case *SelectCommand:
		// 	char, key, e := keyboard.GetKey()
		// 	if e != nil {
		// 		return e
		// 	}
		// 	if key == keyboard.KeyEnter {
		// 		msg = &ProcessCommand{subject: statusPlayer}
		// 		continue
		// 	}
		// 	msg = &SelectCommand{moveCursor(m.command, char), statusPlayer}
		// case *ProcessCommand:
		// 	next := m.subject + 1
		// 	if next == statusCount {
		// 		msg = &SelectCommand{}
		// 		continue
		// 	}
		// 	s := game.status(next)
		// 	msg = &ProcessCommand{subject: s.target}
		// }
	}
}

func moveCursor(c Command, char rune) Command {
	switch char {
	case 'w':
		c--
	case 's':
		c++
	default:
		return c
	}
	if c < 0 {
		c = CommandMax - 1
	} else if CommandMax <= c {
		c = 0
	}
	return c
}
