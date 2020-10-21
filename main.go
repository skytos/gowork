package main

import (
	"fmt"
	"github.com/skytos/goplay"
)

func main() {
	player := goplay.Player{Name: "Sky", Score: 36}
	player.Play()
	goplay.Play()
	fmt.Println("Work!")
}
