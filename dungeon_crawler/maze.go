package main

import (
	"math/rand"
)

const mazeWidth, mazeHeight = 8, 8

type direction = int

const (
	directionNorth direction = iota
	directionWest
	directionSouth
	directionEast
	directionMax
)

var directions = [directionMax]vec2{
	{x: 0, y: -1},
	{x: -1, y: 0},
	{x: 0, y: 1},
	{x: 1, y: 0},
}

type tile struct {
	walls [directionMax]bool
}

type maze struct {
	tiles [mazeHeight][mazeWidth]tile
}

func (m *maze) init() {
	m.generateMap()
}

func (m *maze) generateMap() {
	for y := 0; y < mazeHeight; y++ {
		for x := 0; x < mazeWidth; x++ {
			for i := 0; i < directionMax; i++ {
				m.tiles[y][x].walls[i] = true
			}
		}
	}

	var currentPos vec2
	var toDigTiles []vec2
	toDigTiles = append(toDigTiles, currentPos)

	for {
		var canDigDirections []direction
		for i := 0; i < directionMax; i++ {
			if m.canDigWall(currentPos, i) {
				canDigDirections = append(canDigDirections, i)
			}
		}

		if 0 < len(canDigDirections) {
			digDirection := canDigDirections[rand.Intn(len(canDigDirections))]
			m.digWall(currentPos, digDirection)

			currentPos = currentPos.add(directions[digDirection])
			toDigTiles = append(toDigTiles, currentPos)
		} else {
			toDigTiles = toDigTiles[1:]
			if len(toDigTiles) == 0 {
				return
			}
			currentPos = toDigTiles[0]
		}
	}
}

func (m *maze) digWall(pos vec2, d direction) {
	if !isInsideMaze(pos) {
		return
	}
	m.tiles[pos.y][pos.x].walls[d] = false

	nextPos := pos.add(directions[d])
	if isInsideMaze(nextPos) {
		inverseDirection := (d + 2) % directionMax
		m.tiles[nextPos.y][nextPos.x].walls[inverseDirection] = false
	}
}

func (m *maze) canDigWall(pos vec2, d direction) bool {
	nextPos := pos.add(directions[d])
	if !isInsideMaze(nextPos) {
		return false
	}

	for _, walled := range m.tiles[nextPos.y][nextPos.x].walls {
		if !walled {
			return false
		}
	}
	return true
}

func isInsideMaze(pos vec2) bool {
	return 0 <= pos.x && pos.x < mazeWidth && 0 <= pos.y && pos.y < mazeHeight
}
