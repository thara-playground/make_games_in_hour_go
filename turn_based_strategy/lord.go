package main

type lordID int

const (
	lordDate      lordID = iota // 伊達輝宗
	lordUesugi                  // 上杉謙信
	lordTakeda                  // 武田信玄
	lordHojo                    // 北条氏政
	lordTokugawa                // 徳川家康
	lordOda                     // 織田信長
	lordAshikaga                // 足利義昭
	lordMori                    // 毛利元就
	lordChosokabe               // 長宗我部元親
	lordSimazu                  // 島津義久
	lordMax
)

type lord struct {
	familyName, firstName string
}

var lords = [lordMax]lord{
	{"伊達", "輝宗"},
	{"上杉", "謙信"},
	{"武田", "信玄"},
	{"北条", "氏政"},
	{"徳川", "家康"},
	{"織田", "信長"},
	{"足利", "義昭"},
	{"毛利", "元就"},
	{"長宗我部", "元親"},
	{"島津", "義久"},
}

type castleID = int

const (
	castleYonezawa        castleID = iota // 米沢城
	castleKasugayama                      // 春日山城
	castleTsutsujigasaki                  // 躑躅ヶ崎館
	castleOdawara                         // 小田原城
	castleOkazaki                         // 岡崎城
	castleGifu                            // 岐阜城
	castleNijo                            // 二条城
	castleYoshidakoriyama                 // 吉田郡山城
	castleOko                             // 岡豊城
	castleUchi                            // 内城
	castleMax                             // 種類の数

)

type castle struct {
	name       string
	owner      lordID
	troopCount int

	connectedCastles []castleID
}

var castles = [castleMax]castle{
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

const troopBase = 5
const troopMax = 9000
const troopUnit = 1000
