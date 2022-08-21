package main

type loadID int

const (
	loadDate      loadID = iota // 伊達輝宗
	loadUesugi                  // 上杉謙信
	loadTakeda                  // 武田信玄
	loadHojo                    // 北条氏政
	loadTokugawa                // 徳川家康
	loadOda                     // 織田信長
	loadAshikaga                // 足利義昭
	loadMori                    // 毛利元就
	loadChosokabe               // 長宗我部元親
	loadSimazu                  // 島津義久
	loadMax
)

type load struct {
	familyName, firstName string
}

var loads = [loadMax]load{
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

type castleID int

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
	owner      loadID
	troopCount int
}

const troopBase = 5

var castles = [castleMax]castle{

	{
		name:       "米沢城",
		owner:      loadDate,
		troopCount: troopBase,
	},
	{
		name:       "春日山城",
		owner:      loadUesugi,
		troopCount: troopBase,
	},
	{
		name:       "躑躅ヶ崎館",
		owner:      loadTakeda,
		troopCount: troopBase,
	},
	{
		name:       "小田原城",
		owner:      loadHojo,
		troopCount: troopBase,
	},
	{
		name:       "岡崎城",
		owner:      loadTokugawa,
		troopCount: troopBase,
	},
	{
		name:       "岐阜城",
		owner:      loadOda,
		troopCount: troopBase,
	},
	{
		name:       "二条城",
		owner:      loadAshikaga,
		troopCount: troopBase,
	},
	{
		name:       "吉田郡山城",
		owner:      loadMori,
		troopCount: troopBase,
	},
	{
		name:       "岡豊城",
		owner:      loadChosokabe,
		troopCount: troopBase,
	},
	{
		name:       "内城",
		owner:      loadSimazu,
		troopCount: troopBase,
	},
}
