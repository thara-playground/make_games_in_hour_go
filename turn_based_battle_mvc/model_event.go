package main

type Observer interface {
	NotifyUpdate(Event)
}

type Event interface {
	_event()
}

type OnStart struct {
}

func (OnStart) _event() {}

type OnFightEnd struct {
	Subject Status
	Target  Status
	Damage  int
	Result  FightResult
}

func (OnFightEnd) _event() {}

type OnSpellEnd struct {
	Subject Status
	Result  SpellResult
}

func (OnSpellEnd) _event() {}

type OnRunEnd struct {
	Subject Status
}

func (OnRunEnd) _event() {}

type Command = int

const (
	CommandFight Command = iota
	CommandSpell
	CommandRun
	CommandMax
)

type FightResult int

const (
	_ FightResult = iota
	FightPlayerWin
	FightPlayerLose
	FightContinued
)

type SpellResult int

const (
	_ SpellResult = iota
	SpellSuccess
	SpellFailure
)
