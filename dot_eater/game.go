package main

import (
	"math/rand"
)

type game struct {
	maze       maze
	characters [characterMax]character
}

func (g *game) init() {
	for y, row := range defaultMaze {
		copy(g.maze[y][:], []rune(row[:mazeWidth]))
	}
	for i, defaultPos := range defaultPositions {
		g.characters[i].currentPos = defaultPos
		g.characters[i].lastPos = defaultPos
	}
}

func (g *game) updatePlayerPos(newPos point) {
	newPos = getLoopPos(newPos)

	if !g.maze.isWall(newPos) {
		g.characters[characterPlayer].currentPos = newPos
	}

	p := g.characters[characterPlayer]
	if g.maze[p.currentPos.y][p.currentPos.x] == 'o' {
		g.maze[p.currentPos.y][p.currentPos.x] = ' '
	}
}

func (g *game) update() {
	for i, c := range g.characters {
		g.characters[i].lastPos = c.currentPos

		newPos := c.currentPos
		switch i {
		case characterRandom:
			newPos = getRandomPos(c)
		case characterChase:
			newPos = getFirstPosOfShortestPath(c, g.characters[characterPlayer].currentPos, &g.maze)
		case characterAmbush:
			playerDirection := g.characters[characterPlayer].currentPos.sub(g.characters[characterPlayer].lastPos)
			targetPos := g.characters[characterPlayer].currentPos
			for j := 0; j < 3; j++ {
				targetPos = targetPos.add(playerDirection)
			}
			targetPos = getLoopPos(targetPos)
			newPos = getFirstPosOfShortestPath(c, targetPos, &g.maze)
		case characterSiege:
			chaseToPlayer := g.characters[characterPlayer].currentPos.sub(g.characters[characterChase].currentPos)
			targetPos := g.characters[characterPlayer].currentPos.add(chaseToPlayer)
			targetPos = getLoopPos(targetPos)
			newPos = getFirstPosOfShortestPath(c, targetPos, &g.maze)
		}
		if !g.maze.isWall(newPos) &&
			g.characters[i].lastPos != newPos {
			g.characters[i].currentPos = newPos
		}
	}
}

func (g *game) isCompleted() bool {
	for y := 0; y < mazeHeight; y++ {
		for x := 0; x < mazeWidth; x++ {
			if g.maze[y][x] == 'o' {
				return false
			}
		}
	}
	return true
}

func (g *game) isGameOver() bool {
	for _, c := range g.characters[characterPlayer+1:] {
		if c.currentPos == g.characters[characterPlayer].currentPos {
			return true
		}
	}
	return false
}

func getRandomPos(c character) point {
	var ps []point
	for _, d := range directions {
		newPos := c.currentPos.add(d)
		newPos = getLoopPos(newPos)
		ps = append(ps, newPos)
	}
	return ps[rand.Intn(len(ps))]
}
