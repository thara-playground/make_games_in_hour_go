package main

type vec2 struct {
	x, y int
}

func (v *vec2) add(v2 vec2) vec2 {
	return vec2{x: v.x + v2.x, y: v.y + v2.y}
}
