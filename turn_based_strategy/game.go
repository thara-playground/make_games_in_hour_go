package main

import "math/rand"

const startYear = 1570

type game struct {
	year int

	turnOrder [castleMax]castleID

	playerCastle castleID
	playerLord   lordID

	castles [castleMax]castle
}

func (g *game) init() {
	g.year = startYear

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
	g.playerCastle = c
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

func (g *game) PlayerCastle() castle {
	return g.castles[g.playerCastle]
}

func (g *game) isPlayerCastle(c castleID) bool {
	return g.castles[c].owner == g.playerLord
}

func (g *game) getPlayerTroopMax(targetCastle castleID) int {
	troopMax := g.PlayerCastle().troopCount
	if g.isPlayerCastle(targetCastle) {
		cap := troopMax - g.castle(targetCastle).troopCount
		if cap < troopMax {
			troopMax = cap
		}
	}
	return troopMax
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
		g.castles[target].owner = offence
		g.castles[target].troopCount = *offensiveTroopCount
		return true, siegeResultWin
	}
	return false, siegeResultNone
}
