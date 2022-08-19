package main

import "fmt"

const (
	boardWidth  = 8
	boardHeight = 8
)

type gameMode int

const (
	gameMode1P gameMode = iota
	gameMode2P
	gameModeWatch
	gameModeMax
)

type color int

const (
	colorNone color = iota
	colorBlack
	colorWhite
	colorMax
)

func (c color) opponent() color {
	return c ^ 0b11
}

type reversi struct {
	mode  gameMode
	board [boardHeight][boardWidth]color

	turn color

	isPlayers map[color]bool
}

func newReversi(mode gameMode) *reversi {
	r := &reversi{}
	r.mode = mode
	r.isPlayers = map[color]bool{}

	r.turn = colorBlack
	r.board[3][4], r.board[4][3] = colorBlack, colorBlack
	r.board[3][3], r.board[4][4] = colorWhite, colorWhite

	switch mode {
	case gameMode1P:
		r.isPlayers[colorBlack] = true
	case gameMode2P:
		r.isPlayers[colorBlack] = true
		r.isPlayers[colorWhite] = true
	case gameModeWatch:
	}
	return r
}

func (r *reversi) isPlayerTurn() bool {
	return r.isPlayers[r.turn]
}

func (r *reversi) place(pos point) {
	r.board[pos.y][pos.x] = r.turn
	r.turn = r.turn.opponent()
}

func (r *reversi) color(y, x int) color {
	return r.board[y][x]
}

func (r *reversi) canPlace(pos point, turnOver bool) bool {
	if r.board[pos.y][pos.x] != colorNone {
		return false
	}
	opponent := r.turn.opponent()

	placed := false

	for _, d := range directions {
		current := pos
		current = current.add(d)
		if !current.valid() {
			continue
		}
		if r.board[current.y][current.x] != opponent {
			continue
		}
		for {
			current = current.add(d)
			if !current.valid() {
				break
			}
			switch r.board[current.y][current.x] {
			case colorNone:
				break
			case r.turn:
				if turnOver {
					rev := pos
					rev = pos.add(d)
					for {
						r.board[rev.y][rev.x] = r.turn
						rev = rev.add(d)
						if r.board[rev.y][rev.x] == r.turn {
							break
						}
					}
				}

				placed = true
			}
		}
	}
	return placed
}

func (r *reversi) canPlaceAll() bool {
	for y := 0; y < boardHeight; y++ {
		for x := 0; x < boardWidth; x++ {
			pos := point{x: x, y: y}
			fmt.Println(pos)
			if r.canPlace(pos, false) {
				return true
			}
		}
	}
	return false
}

func (r *reversi) skipTurn() {
	r.turn = r.turn.opponent()
}

func (r *reversi) finish() {
	r.turn = colorNone
}

func (r *reversi) getCount(c color) int {
	n := 0
	for y := 0; y < boardHeight; y++ {
		for x := 0; x < boardWidth; x++ {
			if r.board[y][x] == c {
				n++
			}
		}
	}
	return n
}

type direction = int

const (
	directionUp direction = iota
	directionUpLeft
	directionLeft
	directionDownLeft
	directionDown
	directionDownRight
	directionRight
	directionUpRight
	directionMax
)

var directions = [directionMax]point{
	{0, -1},
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
	{1, 0},
	{1, -1},
}
