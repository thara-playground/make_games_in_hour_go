package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/eiannone/keyboard"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
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

	m := monsterBoss
	g.battle(m)
}
