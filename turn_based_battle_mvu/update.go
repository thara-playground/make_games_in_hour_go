package main

import "math/rand"

type Msg interface{ _msg() }

func (SelectCommand) _msg() {}
func (Next) _msg()          {}

type SelectCommand struct {
	subject int
	command Command
}

type Next struct{}

// type ProcessCommand struct {
// 	subject int
// }

func (CmdStart) _msg()             {}
func (CmdWaitSelectCommand) _msg() {}
func (CmdProcessCommand) _msg()    {}

type CmdStart struct{}
type CmdWaitSelectCommand struct{ subject int }
type CmdProcessCommand struct{ subject int }

type Fight struct{}

func update(msg Msg, model Model) (newState Model, next func() Msg) {
	g := model.Game
	switch m := msg.(type) {
	case *CmdStart:
		model.Current = statusPlayer
		model.Phase = &PhaseCommandSelect{}
		return model.clone(), func() Msg {
			return &CmdWaitSelectCommand{}
		}

	case *CmdWaitSelectCommand:
		model.Current = m.subject

		switch m.subject {
		case statusPlayer:
			model.Phase = &PhaseCommandSelect{}
			return model.clone(), nil
		case statusMonster:
			return model.clone(), func() Msg {
				return &SelectCommand{model.Current, CommandFight}
			}
		}

	case *SelectCommand:
		s := &g.s[model.Current]
		s.cmd = m.command

		switch m.subject {
		case statusPlayer:
			return model.clone(), func() Msg {
				return &CmdWaitSelectCommand{model.Game.targetOf(model.Current)}
			}
		case statusMonster:
			return model.clone(), func() Msg {
				return &CmdProcessCommand{s.target}
			}
		}
	case *Next:
		switch p := model.Phase.(type) {
		case *PhaseStart:
			model.Current = statusPlayer
			model.Phase = &PhaseCommandSelect{}
			return model.clone(), func() Msg {
				return &CmdWaitSelectCommand{}
			}
		case *PhaseCommandSelect:
			s := &g.s[m.subject]
			s.cmd = m.command

			switch model.Current {
			case statusPlayer:
				model.Current = model.Game.targetOf(model.Current)
				return model.clone(), func() Msg {
					return &CmdWaitSelectCommand{}
				}
			case statusMonster:
				model.Current = model.Game.targetOf(model.Current)
			}

			model.Phase = &PhaseBattleTurn{}
		case *PhaseBattleTurn:

			// switch p.subject {
			// case statusPlayer:
			// case statusMonster:
			// }
		}

	case *CmdProcessCommand:
		var result CommandResult
		switch g.s[m.subject].cmd {
		case CommandFight:
			result = fight(&g, m.subject)
		case CommandSpell:
			result = spell(&g, m.subject)
		case CommandRun:
			result = run(&g, m.subject)
		}

		model.Phase = &PhaseBattleTurn{result: result}

		switch m.subject {
		case statusPlayer:
			return model.clone(), func() Msg {
				return &CmdProcessCommand{g.targetOf(m.subject)}
			}
		case statusMonster:
			return model.clone(), func() Msg {
				return &CmdWaitSelectCommand{g.targetOf(m.subject)}
			}
		}
	}
	panic("unrecognized msg")
}

func fight(g *Game, idx int) FightCommandResult {
	subject := &g.s[idx]
	target := &g.s[subject.target]

	dmg := 1 + rand.Int()%subject.attack
	target.HP -= dmg
	if target.HP < 0 {
		target.HP = 0
	}

	result := FightContinued
	if target.HP <= 0 {
		switch subject.target {
		case statusMonster:
			target.AA = ""
			result = FightPlayerWin
		case statusPlayer:
			result = FightPlayerLose
		}
	}
	return FightCommandResult{
		Target: subject.target,
		Damage: dmg,
		Result: result,
	}
}

func spell(g *Game, idx int) SpellCommandResult {
	subject := &g.s[idx]

	var result SpellResult
	if subject.MP < spellCost {
		result = SpellFailure
	} else {
		subject.HP = subject.MaxHP
		subject.MP -= spellCost
		result = SpellSuccess
	}
	return SpellCommandResult{
		Subject: idx,
		Result:  result,
	}
}

func run(g *Game, idx int) RunCommandResult {
	return RunCommandResult{
		Subject: idx,
	}
}
