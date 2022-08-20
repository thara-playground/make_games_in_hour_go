package main

type point struct {
	x, y int
}

func (p *point) add(p2 point) point {
	return point{x: p.x + p2.x, y: p.y + p2.y}
}

func (p *point) sub(p2 point) point {
	return point{x: p.x - p2.x, y: p.y - p2.y}
}
