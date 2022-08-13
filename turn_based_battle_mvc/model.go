package main

import (
	"math/rand"
)

const spellCost = 3

const (
	statusPlayer int = iota
	statusMonster
	statusCount
)

type Status struct {
	character

	cmd    Command
	target int
}

type Model struct {
	s []*Status

	observers []Observer
}

func (m *Model) Init(monster Monster, observers ...Observer) {
	m.s = make([]*Status, statusCount)
	m.s[statusPlayer] = &Status{character: characterConfig[characterPlayer]}

	m.s[statusMonster] = &Status{character: characterConfig[monster]}

	m.s[statusPlayer].target = statusMonster
	m.s[statusMonster].target = statusPlayer

	m.observers = observers
}

func (m *Model) Player() Status {
	return *m.s[statusPlayer]
}

func (m *Model) Monster() Status {
	return *m.s[statusMonster]
}

func (m *Model) Update(playerCommand Command) {
	m.s[statusPlayer].cmd = playerCommand

	for i, c := range m.s {
		switch c.cmd {
		case CommandFight:
			if m.fight(i) {
				return
			}
		case CommandSpell:
			m.spell(i)
		case CommandRun:
			m.run(i)
			return
		}
	}
}

func (m *Model) fight(idx int) bool {
	subject := m.s[idx]
	target := m.s[subject.target]

	dmg := 1 + rand.Int()%subject.attack
	target.HP -= dmg
	if target.HP < 0 {
		target.HP = 0
	}

	exit := false

	result := FightContinued
	if target.HP <= 0 {
		switch subject.target {
		case statusMonster:
			target.AA = ""
			result = FightPlayerWin
		case statusPlayer:
			result = FightPlayerLose
		}
		exit = true
	}
	m.notifyUpdate(&OnFightEnd{
		Subject: *subject,
		Target:  *target,
		Damage:  dmg,
		Result:  result,
	})
	return exit
}

func (m *Model) spell(idx int) {
	subject := m.s[idx]

	var result SpellResult
	if subject.MP < spellCost {
		result = SpellFailure
	} else {
		subject.HP = subject.MaxHP
		subject.MP -= spellCost
		result = SpellSuccess
	}

	m.notifyUpdate(&OnSpellEnd{
		Subject: *subject,
		Result:  result,
	})
}

func (m *Model) run(idx int) {
	subject := m.s[idx]
	m.notifyUpdate(&OnRunEnd{
		Subject: *subject,
	})
}

func (m *Model) notifyUpdate(ev Event) {
	for _, o := range m.observers {
		o.NotifyUpdate(ev)
	}
}
