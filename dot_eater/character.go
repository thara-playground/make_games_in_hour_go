package main

type characterType = int

const (
	characterPlayer characterType = iota
	characterRandom
	characterChase
	characterAmbush
	characterSiege
	characterMax
)

type character struct {
	currentPos, defaultPos, lastPos point
}

var defaultPositions = [characterMax]point{
	{x: 9, y: 13},
	{x: 1, y: 1},
	{x: 17, y: 1},
	{x: 1, y: 7},
	{x: 17, y: 17},
}
