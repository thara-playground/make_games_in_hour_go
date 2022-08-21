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

	chronology string
}

func (g *game) init() {
	g.year = startYear
	g.chronology = ""

	for i := 0; i < int(castleMax); i++ {
		g.turnOrder[i] = castleID(i)
	}
	rand.Shuffle(len(g.turnOrder), func(i, j int) {
		g.turnOrder[i], g.turnOrder[j] = g.turnOrder[j], g.turnOrder[i]
	})

	g.castles = [castleMax]castle{
		{
			name:             "米沢城",
			owner:            lordDate,
			troopCount:       troopBase,
			connectedCastles: []castleID{castleKasugayama, castleOdawara},
		},
		{
			name:             "春日山城",
			owner:            lordUesugi,
			troopCount:       troopBase,
			connectedCastles: []castleID{castleYonezawa, castleTsutsujigasaki, castleGifu},
		},
		{
			name:             "躑躅ヶ崎館",
			owner:            lordTakeda,
			troopCount:       troopBase,
			connectedCastles: []castleID{castleKasugayama, castleOdawara, castleOkazaki},
		},
		{
			name:             "小田原城",
			owner:            lordHojo,
			troopCount:       troopBase,
			connectedCastles: []castleID{castleYonezawa, castleTsutsujigasaki, castleOkazaki},
		},
		{
			name:             "岡崎城",
			owner:            lordTokugawa,
			troopCount:       troopBase,
			connectedCastles: []castleID{castleTsutsujigasaki, castleOdawara, castleGifu},
		},
		{
			name:             "岐阜城",
			owner:            lordOda,
			troopCount:       troopBase,
			connectedCastles: []castleID{castleKasugayama, castleOkazaki, castleNijo},
		},
		{
			name:             "二条城",
			owner:            lordAshikaga,
			troopCount:       troopBase,
			connectedCastles: []castleID{castleGifu, castleYoshidakoriyama, castleOko},
		},
		{
			name:             "吉田郡山城",
			owner:            lordMori,
			troopCount:       troopBase,
			connectedCastles: []castleID{castleNijo, castleOko, castleUchi},
		},
		{
			name:             "岡豊城",
			owner:            lordChosokabe,
			troopCount:       troopBase,
			connectedCastles: []castleID{castleNijo, castleYoshidakoriyama, castleUchi},
		},
		{
			name:             "内城",
			owner:            lordSimazu,
			troopCount:       troopBase,
			connectedCastles: []castleID{castleYoshidakoriyama, castleOko},
		},
	}
}

func (g *game) setPlayerCastle(c castleID) {
	g.playerLord = g.castles[c].owner
}

func (g *game) castle(c castleID) castle {
	return g.castles[c]
}

func (g *game) castleLord(c castleID) lord {
	return lords[g.castles[c].owner]
}

func (g *game) lord(l lordID) lord {
	return lords[l]
}

func (g *game) PlayerLord() lord {
	return lords[g.playerLord]
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
	if g.castles[to].owner == g.playerLord {
		g.castles[to].troopCount += troopCount
	}
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

		if g.getCastleCount(g.castle(target).owner) <= 0 {
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

func (g *game) turnEnd() {
	for i, c := range g.castles {
		if c.troopCount < troopBase {
			g.castles[i].troopCount++
		} else if troopBase < c.troopCount {
			g.castles[i].troopCount--
		}
	}
	g.year++
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
