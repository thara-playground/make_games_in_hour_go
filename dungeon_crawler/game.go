package main

type game struct {
	maze   maze
	player character

	goalX, goalY int
}

func (g *game) init() {
	g.maze.generateMap()

	g.player.pos = vec2{x: 0, y: 0}
	g.player.direction = directionNorth

	g.goalX = mazeWidth - 1
	g.goalY = mazeHeight - 1
}

func (g *game) getTile(y, x int) tile {
	return g.maze.tiles[y][x]
}

func (g *game) isPlayerExist(y, x int) bool {
	return g.player.pos.y == y && g.player.pos.x == x
}

func (g *game) isGoal(y, x int) bool {
	return g.goalX == x && g.goalY == y
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

func (g *game) isReachedGoal() bool {
	return g.player.pos.x == g.goalX && g.player.pos.y == g.goalY
}
