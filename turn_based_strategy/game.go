package main

import (
	"fmt"
	"math/rand"
)

const startYear = 1570

type game struct {
	year int

	turnOrder [castleMax]castleID

	playerLord lordID

	castles [castleMax]castle
	lords   [lordMax]lord

	chronology string
}

func (g *game) init() {
	g.year = startYear
	g.chronology = ""

	for i := range castles {
		g.turnOrder[i] = castleID(i)
	}
	rand.Shuffle(len(g.turnOrder), func(i, j int) {
		g.turnOrder[i], g.turnOrder[j] = g.turnOrder[j], g.turnOrder[i]
	})
	for i, c := range castles {
		cs := make([]castleID, len(c.connectedCastles))
		copy(cs[:], c.connectedCastles[:])

		g.castles[i] = c
		g.castles[i].connectedCastles = cs
	}
	copy(g.lords[:], lords[:])
}

func (g *game) setPlayerCastle(c castleID) {
	g.playerLord = g.castles[c].owner
}

func (g *game) castle(c castleID) castle {
	return g.castles[c]
}

func (g *game) castleLord(c castleID) lord {
	return g.lords[g.castles[c].owner]
}

func (g *game) lord(l lordID) lord {
	return g.lords[l]
}

func (g *game) PlayerLord() lord {
	return g.lords[g.playerLord]
}

func (g *game) isPlayerCastle(c castleID) bool {
	return g.castles[c].owner == g.playerLord
}

func (g *game) getPlayerTroopMax(currentCastle, targetCastle castleID) int {
	max := g.castles[currentCastle].troopCount
	if g.castles[targetCastle].owner == g.playerLord {
		cap := troopMax - g.castle(targetCastle).troopCount
		if cap < max {
			max = cap
		}
	}
	return max
}

func (g *game) advance(from, to castleID, troopCount int) {
	g.castles[from].troopCount -= troopCount
}

func (g *game) sendTroops(from, to castleID, troopCount int) {
	g.castles[from].troopCount -= troopCount
	g.castles[to].troopCount += troopCount
}

type siegeResult int

const (
	siegeResultNone siegeResult = iota
	siegeResultWin
	siegeResultLose
)

func (g *game) processSiege(offence lordID, target castleID, offensiveTroopCount *int) (bool, siegeResult) {
	if rand.Intn(2) == 0 {
		g.castles[target].troopCount--
	} else {
		*offensiveTroopCount--
	}

	if *offensiveTroopCount <= 0 {
		return true, siegeResultLose
	} else if g.castles[target].troopCount <= 0 {
		defense := g.castle(target).owner

		g.castles[target].owner = offence
		g.castles[target].troopCount = *offensiveTroopCount

		if g.getCastleCount(defense) <= 0 {
			g.chronology += fmt.Sprintf("%dねん　%s%sが　%sで　%s%sを　ほろぼす\n",
				g.year,
				g.lord(offence).familyName,
				g.lord(offence).firstName,
				g.castle(target).name,
				g.lord(defense).familyName,
				g.lord(defense).firstName,
			)
		}
		return true, siegeResultWin
	}

	return false, siegeResultNone
}

type event int

const (
	eventNone event = iota
	eventHonnoujinhHen
)

func (g *game) turnEnd() event {
	for i, c := range g.castles {
		if c.troopCount < troopBase {
			g.castles[i].troopCount++
		} else if troopBase < c.troopCount {
			g.castles[i].troopCount--
		}
	}
	g.year++

	// if g.year == 1582 && g.castles[castleNijo].owner == lordOda {
	if g.year == 1574 && g.castles[castleNijo].owner == lordOda {
		g.honnoujinhHen()
		return eventHonnoujinhHen
	}

	return eventNone
}

func (g *game) honnoujinhHen() {
	g.lords[lordOda].familyName = "羽柴"
	g.lords[lordOda].firstName = "秀吉"
}

func (g *game) getCastleCount(lord lordID) int {
	castleCount := 0
	for _, c := range g.castles {
		if c.owner == lord {
			castleCount++
		}
	}
	return castleCount
}
