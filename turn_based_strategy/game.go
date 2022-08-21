package main

const startYear = 1570

type game struct {
	year int
}

func (g *game) init() {
	g.year = startYear
}
