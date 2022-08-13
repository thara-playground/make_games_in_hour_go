package main

import (
	"flag"
	"fmt"
	"time"
)

type cell struct {
	row, col int
}

func (c *cell) neighbors() []cell {
	var ns []cell
	for _, h := range []int{-1, 0, 1} {
		for _, w := range []int{-1, 0, 1} {
			if w == 0 && h == 0 {
				continue
			}
			ns = append(ns, cell{row: c.row + h, col: c.col + w})
		}
	}
	return ns
}

type state map[cell]bool

func advanced(current state) state {
	recalc := []cell{}
	for cell := range current {
		recalc = append(recalc, cell)
		for _, c := range cell.neighbors() {
			recalc = append(recalc, c)
		}
	}

	next := map[cell]bool{}
	for _, cell := range recalc {
		numLived := 0
		for _, c := range cell.neighbors() {
			if current[c] {
				numLived++
			}
		}
		if numLived == 3 || (numLived == 2 && current[cell]) {
			next[cell] = true
		}
	}
	return state(next)
}

const width, height = 8, 8

func draw(s state) {
	fmt.Print("\033[H\033[2J")

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			if s[cell{row: h, col: w}] {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

var pattern int
var duration int

func init() {
	flag.IntVar(&pattern, "p", 0, "pattern no. default: 0")
	flag.IntVar(&duration, "d", 500, "time duration between generations(ms). default: 500")
}

func main() {
	flag.Parse()

	s := state{}
	for _, cell := range patterns[pattern] {
		s[cell] = true
	}
	draw(s)

	t := time.NewTicker(time.Duration(duration) * time.Millisecond)
	for range t.C {
		s = advanced(s)
		draw(s)
	}
}

var patterns = [][]cell{
	// 0. Brinker
	{
		{row: 3, col: 4},
		{row: 3, col: 5},
		{row: 3, col: 6},
	},
	// 1. Toad
	{
		{row: 1, col: 2},
		{row: 1, col: 3},
		{row: 1, col: 4},
		{row: 2, col: 1},
		{row: 2, col: 2},
		{row: 2, col: 3},
	},
	// 2. Beacon
	{
		{row: 1, col: 1},
		{row: 1, col: 2},
		{row: 2, col: 1},
		{row: 3, col: 4},
		{row: 4, col: 3},
		{row: 4, col: 4},
	},
	// 3. Glider
	{
		{row: 0, col: 1},
		{row: 1, col: 2},
		{row: 2, col: 0},
		{row: 2, col: 1},
		{row: 2, col: 2},
	},
	// 4. R-Pentomino
	{
		{row: 2, col: 2},
		{row: 2, col: 3},
		{row: 3, col: 1},
		{row: 3, col: 2},
		{row: 4, col: 2},
	},
}
