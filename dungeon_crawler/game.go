package main

type game struct {
	maze   maze
	player character
}

func (g *game) init() {
	g.maze.generateMap()

	g.player.pos = vec2{x: 0, y: 0}
	g.player.direction = directionNorth
}

func (g *game) getTile(y, x int) tile {
	return g.maze.tiles[y][x]
}

func (g *game) isPlayerExist(y, x int) bool {
	return g.player.pos.y == y && g.player.pos.x == x
}

func (g *game) movePlayerForward() {
	if !g.maze.tiles[g.player.pos.y][g.player.pos.x].walls[g.player.direction] {
		nextPos := g.player.pos.add(directions[g.player.direction])
		if isInsideMaze(nextPos) {
			g.player.pos = nextPos
		}
	}
}

func (g *game) turnPlayerLeft() {
	g.player.direction++
	g.player.direction = (g.player.direction + directionMax) % directionMax
}

func (g *game) turnPlayerRight() {
	g.player.direction--
	g.player.direction = (g.player.direction + directionMax) % directionMax
}

func (g *game) turnPlayerAround() {
	g.player.direction += 2
	g.player.direction = (g.player.direction + directionMax) % directionMax
}
