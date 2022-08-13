package main

type character struct {
	Name string

	HP, MaxHP int
	MP, MaxMP int

	attack int

	AA string
}

const (
	characterPlayer = iota
)

type Monster int

const (
	MonsterSlime Monster = 1
	MonsterBoss  Monster = 2
)

var characterConfig = []character{
	{Name: "ゆうしゃ", HP: 100, MaxHP: 100, MP: 15, MaxMP: 15, attack: 30},
	{Name: "スライム", HP: 3, MaxHP: 3, MP: 0, MaxMP: 0, attack: 2,
		AA: "／・Д・＼\n" +
			"~~~~~~~~~",
	},
	{Name: "まおう", HP: 255, MaxHP: 255, MP: 0, MaxMP: 0, attack: 50,
		AA: "   A＠A\n" +
			"ψ (▼皿▼) ψ",
	},
}
