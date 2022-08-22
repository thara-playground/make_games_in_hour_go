package main

import (
	"math/rand"
	"sort"
)

type aiCommand int

const (
	aiCommandNone aiCommand = iota
	aiCommandAdvanceEnemyCastle
	aiCommandAdvanceFrontCastle
)

func processAI(g *game, currentCastle castleID) (command aiCommand, targetCastle castleID, troopCount int) {
	var enemyCastles []castleID
	for _, c := range g.castle(currentCastle).connectedCastles {
		if g.castle(c).owner != g.castle(currentCastle).owner {
			enemyCastles = append(enemyCastles, c)
		}
	}
	if 0 < len(enemyCastles) {
		sort.Slice(enemyCastles, func(i, j int) bool {
			return g.castle(i).troopCount < g.castle(j).troopCount
		})

		for 1 < len(enemyCastles) &&
			g.castle(enemyCastles[0]).troopCount < g.castle(enemyCastles[len(enemyCastles)-1]).troopCount {
			enemyCastles = enemyCastles[:len(enemyCastles)-1]
		}

		targetCastle = enemyCastles[rand.Intn(len(enemyCastles))]

		if troopBase <= g.castle(currentCastle).troopCount ||
			g.castle(targetCastle).troopCount*2 <= g.castle(currentCastle).troopCount-1 {

			troopCount = g.castle(currentCastle).troopCount - 1
			if troopCount < 0 {
				troopCount = 0
			}

			command = aiCommandAdvanceEnemyCastle
			return
		}
	} else {
		var frontCastles []castleID
		for _, n := range g.castle(currentCastle).connectedCastles {
			for _, nn := range g.castle(n).connectedCastles {
				if g.castle(nn).owner == g.castle(n).owner {
					frontCastles = append(frontCastles, n)
				}
			}
		}
		var dest []castleID
		if len(frontCastles) == 0 {
			dest = g.castle(currentCastle).connectedCastles
		} else {
			dest = frontCastles
		}
		sort.Slice(dest, func(i, j int) bool {
			return g.castle(dest[i]).troopCount < g.castle(dest[j]).troopCount
		})
		for 1 < len(dest) &&
			g.castle(dest[0]).troopCount < g.castle(dest[len(dest)-1]).troopCount {
			dest = dest[:len(dest)-1]
		}

		targetCastle = dest[rand.Intn(len(dest))]

		troopCount = troopMax - g.castle(targetCastle).troopCount

		if 0 < len(frontCastles) {
			nTroops := g.castle(currentCastle).troopCount
			if nTroops < troopCount {
				troopCount = nTroops
			}
		} else {
			nTroops := g.castle(currentCastle).troopCount - (troopBase - 1)
			if nTroops < troopCount {
				troopCount = nTroops
			}
		}

		if 0 < troopCount {
			command = aiCommandAdvanceFrontCastle
			return
		}
	}
	return
}
