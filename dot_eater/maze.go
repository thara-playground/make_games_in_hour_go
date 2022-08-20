package main

const mazeWidth, mazeHeight = 19, 19

type maze [mazeHeight][mazeWidth]rune

func (m *maze) isWall(pos point) bool {
	return m[pos.y][pos.x] == '#'
}

var defaultMaze = [mazeHeight]string{
	"#########o#########",
	"#ooooooo#o#ooooooo#",
	"#o###o#o#o#o#o###o#",
	"#o# #o#ooooo#o# #o#",
	"#o###o###o###o###o#",
	"#ooooooooooooooooo#",
	"#o###o###o###o###o#",
	"#ooo#o#ooooo#o#ooo#",
	"###o#o#o###o#o#o###",
	"oooooooo# #oooooooo",
	"###o#o#o###o#o#o###",
	"#ooo#o#ooooo#o#ooo#",
	"#o###o###o###o###o#",
	"#oooooooo oooooooo#",
	"#o###o###o###o###o#",
	"#o# #o#ooooo#o# #o#",
	"#o###o#o#o#o#o###o#",
	"#ooooooo#o#ooooooo#",
	"#########o#########",
}

func setDefaultMaze(m *maze) {
	for y, row := range defaultMaze {
		copy(m[y][:], []rune(row[:mazeWidth]))
	}
}

func copyMaze(dst, src *maze) {
	for y, row := range src {
		copy(dst[y][:], row[:])
	}
}

func getLoopPos(pos point) point {
	return point{
		x: (mazeWidth + pos.x) % mazeWidth,
		y: (mazeHeight + pos.y) % mazeHeight,
	}
}
