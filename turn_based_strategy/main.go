package main

import "fmt"

func main() {

	var g game
	g.init()
	draw(&g)
}

func draw(g *game) {
	// fmt.Print(`
	// 1570ねん　〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜　　　　　〜
	// 　　　　　〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜　0米沢5　〜
	// 〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜1春日5　伊達　〜〜
	// 〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜　〜〜上杉　　　　　〜〜
	// 〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜　〜　　　　　　　　〜〜
	// 〜〜〜〜〜〜〜〜〜〜〜〜〜〜　　　　　2躑躅5　　　〜〜
	// 〜〜〜〜〜〜〜〜〜〜〜〜〜　　　　　　武田　　　〜〜〜
	// 〜〜〜〜〜〜　　　　　　　5岐阜5　　　　　　　　〜〜〜
	// 〜〜〜〜　7吉田5　6二条5　織田　4岡崎5　3小田5　〜〜〜
	// 〜〜〜　　毛利　　足利　　　　　徳川　　北条〜〜〜〜〜
	// 〜〜　〜〜〜〜〜〜〜　　　〜〜〜〜〜〜〜〜〜〜〜〜〜〜
	// 〜　　　〜　8岡豊5〜〜　〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜
	// 〜　　　〜〜長宗〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜
	// 〜9内城5〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜
	// 〜島津〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜
	// 〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜
	// `)

	fmt.Printf("%dねん　〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜　　　　　〜\n", g.year)
	fmt.Printf("　　　　　〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜　%d%.2s%d　〜\n",
		castleYonezawa,
		castles[castleYonezawa].name,
		castles[castleYonezawa].troopCount,
	)
	fmt.Printf("〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜%d%.2s%d　%.2s　〜〜\n",
		castleKasugayama,
		castles[castleKasugayama].name,
		castles[castleKasugayama].troopCount,
		loads[castles[castleYonezawa].owner].familyName,
	)
	fmt.Printf("〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜　〜〜%.2s　　　　　〜〜\n",
		loads[castles[castleKasugayama].owner].familyName,
	)
	fmt.Print("〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜　〜　　　　　　　　〜〜\n")
	fmt.Printf("〜〜〜〜〜〜〜〜〜〜〜〜〜〜　　　　　%d%.2s%d　　　〜〜\n",
		castleTsutsujigasaki,
		castles[castleTsutsujigasaki].name,
		castles[castleTsutsujigasaki].troopCount,
	)
	fmt.Printf("〜〜〜〜〜〜〜〜〜〜〜〜〜　　　　　　%.2s　　　〜〜〜\n",
		loads[castles[castleYonezawa].owner].familyName,
	)
	fmt.Printf("〜〜〜〜〜〜　　　　　　　%d%.2s%d　　　　　　　　〜〜〜\n",
		castleGifu,
		castles[castleGifu].name,
		castles[castleGifu].troopCount,
	)
	fmt.Printf("〜〜〜〜　%d%.2s%d　%d%.2s%d　%.2s　　　　　%d%.2s%d　〜〜〜\n",
		castleYoshidakoriyama,
		castles[castleYoshidakoriyama].name,
		castles[castleYoshidakoriyama].troopCount,
		castleNijo,
		castles[castleNijo].name,
		castles[castleNijo].troopCount,
		loads[castles[castleGifu].owner].familyName,
		castleOdawara,
		castles[castleOdawara].name,
		castles[castleOdawara].troopCount,
	)
	fmt.Printf("〜〜〜　　%.2s　　%.2s　〜　　　%d%.2s%d　%.2s〜〜〜〜〜\n",
		loads[castles[castleYoshidakoriyama].owner].familyName,
		loads[castles[castleNijo].owner].familyName,
		castleOkazaki,
		castles[castleOkazaki].name,
		castles[castleOkazaki].troopCount,
		loads[castles[castleOdawara].owner].familyName,
	)
	fmt.Printf("〜〜　〜〜〜〜〜〜〜　　　　〜〜%.2s〜　〜　〜〜〜〜〜\n",
		loads[castles[castleOkazaki].owner].familyName,
	)
	fmt.Printf("〜　　　〜　%d%.2s%d　〜　　　　〜〜〜〜〜〜〜〜〜〜〜〜\n",
		castleOko,
		castles[castleOko].name,
		castles[castleOko].troopCount,
	)
	fmt.Printf("〜　　　〜　%.2s　〜〜　　〜〜〜〜〜〜〜〜〜〜〜〜〜〜\n",
		loads[castles[castleOko].owner].familyName,
	)
	fmt.Printf("〜%d%.2s%d〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜\n",
		castleUchi,
		castles[castleUchi].name,
		castles[castleUchi].troopCount,
	)
	fmt.Printf("〜%.2s〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜\n",
		loads[castles[castleUchi].owner].familyName,
	)
	fmt.Println("〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜")
	fmt.Println()
}
