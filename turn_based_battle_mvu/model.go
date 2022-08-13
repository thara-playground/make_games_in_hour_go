package main

const spellCost = 3

const (
	statusPlayer int = iota
	statusMonster
	statusCount
)

type Model struct {
	Game    Game
	Phase   Phase
	Current int
}

func (m *Model) clone() Model {
	return Model{Game: m.Game.Clone(), Phase: m.Phase}
}

type Phase interface{ _phase() }

func (p *PhaseStart) _phase()         {}
func (p *PhaseCommandSelect) _phase() {}
func (p *PhaseBattleTurn) _phase()    {}

type PhaseStart struct{}
type PhaseCommandSelect struct {
	command Command
}
type PhaseBattleTurn struct {
	result CommandResult
}

type CommandResult interface{ _commandResult() }

func (FightCommandResult) _commandResult() {}
func (SpellCommandResult) _commandResult() {}
func (RunCommandResult) _commandResult()   {}

type FightCommandResult struct {
	Subject int
	Target  int
	Damage  int
	Result  FightResult
}

type SpellCommandResult struct {
	Subject int
	Result  SpellResult
}

type RunCommandResult struct{ Subject int }

type Status struct {
	character

	cmd    Command
	target int
}

type Command = int

const (
	CommandFight Command = iota
	CommandSpell
	CommandRun
	CommandMax
)

type Game struct {
	s []Status
}

func NewGame(monster Monster) Game {
	var g Game
	g.s = make([]Status, statusCount)
	g.s[statusPlayer] = Status{character: characterConfig[characterPlayer]}

	g.s[statusMonster] = Status{character: characterConfig[monster]}

	g.s[statusPlayer].target = statusMonster
	g.s[statusMonster].target = statusPlayer
	return g
}

func (g *Game) status(i int) Status {
	return g.s[i]
}

func (g *Game) targetOf(i int) int {
	return g.s[statusPlayer].target
}

func (g *Game) Player() Status {
	return g.s[statusPlayer]
}

func (g *Game) Monster() Status {
	return g.s[statusMonster]
}

func (g *Game) Clone() Game {
	dst := make([]Status, len(g.s))
	copy(dst, g.s)
	return Game{dst}
}

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
