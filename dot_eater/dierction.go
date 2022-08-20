package main

type direction = int

const (
	directionUp direction = iota
	directionLeft
	directionDown
	directionRight
	directionMax
)

var directions = [directionMax]point{
	{0, -1},
	{-1, 0},
	{0, 1},
	{1, 0},
}
