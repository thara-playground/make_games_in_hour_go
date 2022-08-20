package main

func getFirstPosOfShortestPath(c character, target point, maze *maze) point {
	var checkPositions []point
	checkPositions = append(checkPositions, c.currentPos)

	var distances [mazeHeight][mazeWidth]int
	for y := 0; y < mazeHeight; y++ {
		for x := 0; x < mazeWidth; x++ {
			distances[y][x] = -1
		}
	}
	distances[c.currentPos.y][c.currentPos.x] = 0

	var routes [mazeHeight][mazeWidth][]point

	for 0 < len(checkPositions) {
		for _, d := range directions {
			newPos := checkPositions[0].add(d)
			newPos = getLoopPos(newPos)

			newDis := distances[checkPositions[0].y][checkPositions[0].x] + 1

			if (distances[newPos.y][newPos.x] < 0 || newDis < distances[newPos.y][newPos.x]) && !maze.isWall(newPos) {
				distances[newPos.y][newPos.x] = newDis
				checkPositions = append(checkPositions, newPos)
				routes[newPos.y][newPos.x] = routes[checkPositions[0].y][checkPositions[0].x]
				routes[newPos.y][newPos.x] = append(
					routes[newPos.y][newPos.x], newPos)
			}
		}
		checkPositions = checkPositions[1:]
	}

	if 0 < len(routes[target.y][target.x]) && c.lastPos != routes[target.y][target.x][0] {
		return routes[target.y][target.x][0]
	}
	return getRandomPos(c)
}
