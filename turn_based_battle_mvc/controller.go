package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

type Controller struct {
	model *Model
	view  *View
	exit  bool
}

func (c *Controller) update() error {
	_, key, err := keyboard.GetKey()
	if err != nil {
		return fmt.Errorf("controller input: %#v", err)
	}
	if key == keyboard.KeyEsc {
		c.exit = true
		return nil
	}

	cmd, err := c.selectCommand()
	if err != nil {
		return err
	}
	c.model.Update(cmd)

	return nil
}

func (c *Controller) selectCommand() (cmd Command, err error) {
	cmd = CommandFight

	for {
		c.view.drawSelectCommand(cmd)

		char, key, e := keyboard.GetKey()
		if e != nil {
			return cmd, fmt.Errorf("controller selectCommand: %#v", e)
		}
		if key == keyboard.KeyEnter {
			return
		}

		c := cmd
		switch char {
		case 'w':
			c--
		case 's':
			c++
		default:
			return
		}
		if c < 0 {
			c = CommandMax - 1
		} else if CommandMax <= c {
			c = 0
		}
		cmd = c
	}
}

func (c *Controller) NotifyUpdate(ev Event) {
	switch e := ev.(type) {
	case *OnFightEnd:
		switch e.Result {
		case FightPlayerWin, FightPlayerLose:
			c.exit = true
		}
	case *OnSpellEnd:
		// NOP
	case *OnRunEnd:
		c.exit = true
	}
	keyboard.GetKey()
}
