package main

type character struct {
	name string

	hp, maxHP int
	mp, maxMP int

	attack int

	aa string
}

const (
	characterPlayer = iota
)

type monster int

const (
	monsterSlime monster = 1
	monsterBoss  monster = 2
)

var characterConfig = []character{
	{name: "ゆうしゃ", hp: 100, maxHP: 100, mp: 15, maxMP: 15, attack: 30},
	{name: "スライム", hp: 3, maxHP: 3, mp: 0, maxMP: 0, attack: 2,
		aa: "／・Д・＼\n" +
			"~~~~~~~~~",
	},
	{name: "まおう", hp: 255, maxHP: 255, mp: 0, maxMP: 0, attack: 50,
		aa: "   A＠A\n" +
			"ψ (▼皿▼) ψ",
	},
}
