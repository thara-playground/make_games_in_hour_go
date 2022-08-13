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

	run()
}

func run() {
	var (
		model Model
		view  View
		ctrl  Controller
	)
	view.model = &model
	ctrl.model = &model
	ctrl.view = &view

	model.Init(MonsterBoss, &view, &ctrl)

	view.initScreen()

	for !ctrl.exit {
		err := ctrl.update()
		if err != nil {
			log.Fatal(err)
		}
	}
}
