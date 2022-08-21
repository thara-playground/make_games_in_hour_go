package main

type lord int

const (
	lordDate      lord = iota // 伊達輝宗
	lordUesugi                // 上杉謙信
	lordTakeda                // 武田信玄
	lordHojo                  // 北条氏政
	lordTokugawa              // 徳川家康
	lordOda                   // 織田信長
	lordAshikaga              // 足利義昭
	lordMori                  // 毛利元就
	lordChosokabe             // 長宗我部元親
	lordSimazu                // 島津義久
	lordMax
)

type load struct {
	familyName, firstName string
}

var lords = [lordMax]load{
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
	owner      lord
	troopCount int

	connectedCastles []castleID
}

const troopBase = 5
const troopMax = 9000
const troopUnit = 1000
