package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/eiannone/keyboard"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := keyboard.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	var g game
	g.init()
	draw(&g)

	selectCastle(&g)
	waitKey()

	for {
		for i := 0; i < int(castleMax); i++ {
			draw(&g)
			for j := 0; j < int(castleMax); j++ {
				if j == i {
					fmt.Print(" >")
				} else {
					fmt.Print("  ")
				}
				fmt.Printf("%.2s", g.castle(g.turnOrder[j]).name)
			}
			fmt.Print("\n\n")

			currentCastle := g.turnOrder[i]

			fmt.Printf("%sけの　%sの　ひょうじょうちゅう…\n",
				g.castleLord(currentCastle).familyName,
				g.castle(currentCastle).name)
			fmt.Println()

			if g.isPlayerCastle(currentCastle) {
				selectPlayerCommand(&g)
			} else {
				processAICommand(&g, currentCastle)
			}

			waitKey()
		}

		g.year++
	}
}

func draw(g *game) {
	fmt.Print("\033[H\033[2J")

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
		g.castle(castleYonezawa).name,
		g.castle(castleYonezawa).troopCount,
	)
	fmt.Printf("〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜%d%.2s%d　%.2s　〜〜\n",
		castleKasugayama,
		g.castle(castleKasugayama).name,
		g.castle(castleKasugayama).troopCount,
		g.castleLord(castleYonezawa).familyName,
	)
	fmt.Printf("〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜　〜〜%.2s　　　　　〜〜\n",
		g.castleLord(castleKasugayama).familyName,
	)
	fmt.Print("〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜　〜　　　　　　　　〜〜\n")
	fmt.Printf("〜〜〜〜〜〜〜〜〜〜〜〜〜〜　　　　　%d%.2s%d　　　〜〜\n",
		castleTsutsujigasaki,
		g.castle(castleTsutsujigasaki).name,
		g.castle(castleTsutsujigasaki).troopCount,
	)
	fmt.Printf("〜〜〜〜〜〜〜〜〜〜〜〜〜　　　　　　%.2s　　　〜〜〜\n",
		g.castleLord(castleYonezawa).familyName,
	)
	fmt.Printf("〜〜〜〜〜〜　　　　　　　%d%.2s%d　　　　　　　　〜〜〜\n",
		castleGifu,
		g.castle(castleGifu).name,
		g.castle(castleGifu).troopCount,
	)
	fmt.Printf("〜〜〜〜　%d%.2s%d　%d%.2s%d　%.2s　　　　　%d%.2s%d　〜〜〜\n",
		castleYoshidakoriyama,
		g.castle(castleYoshidakoriyama).name,
		g.castle(castleYoshidakoriyama).troopCount,
		castleNijo,
		g.castle(castleNijo).name,
		g.castle(castleNijo).troopCount,
		g.castleLord(castleGifu).familyName,
		castleOdawara,
		g.castle(castleOdawara).name,
		g.castle(castleOdawara).troopCount,
	)
	fmt.Printf("〜〜〜　　%.2s　　%.2s　〜　　　%d%.2s%d　%.2s〜〜〜〜〜\n",
		g.castleLord(castleYoshidakoriyama).familyName,
		g.castleLord(castleNijo).familyName,
		castleOkazaki,
		g.castle(castleOkazaki).name,
		g.castle(castleOkazaki).troopCount,
		g.castleLord(castleOdawara).familyName,
	)
	fmt.Printf("〜〜　〜〜〜〜〜〜〜　　　　〜〜%.2s〜　〜　〜〜〜〜〜\n",
		g.castleLord(castleOkazaki).familyName,
	)
	fmt.Printf("〜　　　〜　%d%.2s%d　〜　　　　〜〜〜〜〜〜〜〜〜〜〜〜\n",
		castleOko,
		g.castle(castleOko).name,
		g.castle(castleOko).troopCount,
	)
	fmt.Printf("〜　　　〜　%.2s　〜〜　　〜〜〜〜〜〜〜〜〜〜〜〜〜〜\n",
		g.castleLord(castleOko).familyName,
	)
	fmt.Printf("〜%d%.2s%d〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜\n",
		castleUchi,
		g.castle(castleUchi).name,
		g.castle(castleUchi).troopCount,
	)
	fmt.Printf("〜%.2s〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜\n",
		g.castleLord(castleUchi).familyName,
	)
	fmt.Println("〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜")
	fmt.Println()
}

func selectCastle(g *game) {
	fmt.Printf("おやかたさま、われらがしろは　このちずの\n"+
		"どこに　ありまするか？！（0〜%d）\n",
		castleMax-1)

	for {
		char, _ := waitKey()
		selected, err := strconv.Atoi(string(char))
		if err != nil {
			continue
		} else if selected < 0 || int(castleMax) <= selected {
			continue
		}

		g.setPlayerCastle(selected)

		break
	}

	fmt.Printf("%sさま、%sから　てんかとういつを\nめざしましょうぞ！\n",
		g.PlayerLord().firstName,
		g.PlayerCastle().name,
	)
}

func selectPlayerCommand(g *game) {
	fmt.Printf("%sさま、どこに　しんぐん　しますか？\n",
		g.PlayerLord().firstName)
	for i, c := range g.PlayerCastle().connectedCastles {
		fmt.Printf("%d %s\n", i, g.castle(c).name)
	}
	fmt.Println()

	char, _ := waitKey()
	targetCastle, err := strconv.Atoi(string(char))
	if err != nil {
		return
	}

	isConnected := false
	for i := range g.PlayerCastle().connectedCastles {
		if i == targetCastle {
			isConnected = true
			break
		}
	}
	if !isConnected {
		fmt.Println("しんぐんを　とりやめました")
		return
	}

	troopMax := g.getPlayerTroopMax(targetCastle)

	fmt.Printf("%sに　なんぜんにん　しんぐん　しますか？（0〜%d）\n",
		g.castle(targetCastle).name,
		troopMax,
	)

	troopCount := 0
	for {
		char, _ := waitKey()
		n, err := strconv.Atoi(string(char))
		if err != nil {
			continue
		} else if n < 0 || troopMax < n {
			continue
		}
		troopCount = n
		break
	}

	g.advance(g.playerCastle, targetCastle, troopCount)

	fmt.Println()

	if g.isPlayerCastle(targetCastle) {
		fmt.Printf("%sに　%dにん　いどう　しました",
			g.castle(targetCastle).name,
			troopCount*troopUnit,
		)
	} else {
		fmt.Printf("%sに　%dにんで　しゅつじんじゃ〜！！",
			g.castle(targetCastle).name,
			troopCount*troopUnit,
		)
		waitKey()
		siege(g, g.playerLord, troopCount, targetCastle)
	}
}

func processAICommand(g *game, currentCastle castleID) {
	command, targetCastle, troopCount := processAI(g, currentCastle)
	switch command {
	case aiCommandAdvanceEnemyCastle:
		fmt.Printf("%sの　%s%sが　%sに　せめこみました！\n",
			g.castle(currentCastle).name,
			g.castleLord(currentCastle).familyName,
			g.castleLord(currentCastle).firstName,
			g.castle(targetCastle).name)
		waitKey()

		siege(g, g.castle(currentCastle).owner, troopCount, targetCastle)

	case aiCommandAdvanceFrontCastle:
		g.advance(currentCastle, targetCastle, troopCount)

		fmt.Printf("%sから　%sに　%dにん　いどうしました\n",
			g.castle(currentCastle).name,
			g.castle(targetCastle).name,
			troopCount*troopUnit,
		)
	}
}

func siege(g *game, offence lordID, troopCount int, target castleID) {
	fmt.Print("\033[H\033[2J")

	defense := g.castleLord(target)

	var result siegeResult
	for {
		fmt.Printf("%sぐん（%4dにん）　Ｘ　%sぐん（%4dにん）\n",
			g.lord(offence).familyName,
			troopCount*troopUnit,
			defense.familyName,
			g.castle(target).troopCount*troopUnit,
		)

		var finished bool
		finished, result = g.processSiege(offence, target, &troopCount)
		if finished {
			break
		}

		waitKey()
	}

	switch result {
	case siegeResultWin:
		fmt.Printf("%s　らくじょう！！\n", g.castle(target).name)
		fmt.Printf("%sは　%sけの　ものとなります\n",
			g.castle(target).name,
			g.lord(offence).familyName,
		)
		fmt.Println()
	case siegeResultLose:
		fmt.Printf("%sぐん　かいめつ！！\n"+
			"\n"+
			"%sぐんが　%sを　まもりきりました！\n",
			g.lord(offence).familyName,
			defense.familyName,
			g.castle(target).name,
		)
	}

	fmt.Println()
}

func waitKey() (char rune, key keyboard.Key) {
	var err error
	char, key, err = keyboard.GetKey()
	if err != nil {
		panic(err)
	}
	if key == keyboard.KeyEsc {
		os.Exit(0)
	}
	return char, key
}
