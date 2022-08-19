package main

import "math/rand"

func selectPoint(r *reversi) (p point, found bool) {
	var ps []point

	for y := 0; y < boardHeight; y++ {
		for x := 0; x < boardWidth; x++ {
			pos := point{x: x, y: y}
			if r.canPlace(pos, false) {
				ps = append(ps, pos)
			}
		}
	}
	switch len(ps) {
	case 0:
		return p, false
	case 1:
		return ps[0], true
	default:
		return ps[rand.Intn(len(ps)-1)], true
	}
}
