package main

type blockShapeType int

const (
	blockShapeTypeI blockShapeType = iota
	blockShapeTypeL
	blockShapeTypeMax
)

const blockWidthMax, blockHeightMax = 4, 4

type blockShape struct {
	size    int
	pattern [blockWidthMax][blockHeightMax]int
}

var blockShapes = [blockShapeTypeMax]blockShape{
	{
		size: 3,
		pattern: [blockWidthMax][blockHeightMax]int{
			{0, 1, 0, 0},
			{0, 1, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 0, 0},
		},
	},
	{
		size: 3,
		pattern: [blockWidthMax][blockHeightMax]int{
			{0, 1, 0, 0},
			{0, 1, 1, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
	},
}
