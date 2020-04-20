package main

import (
	"fmt"
	"head-first-go/31/gadget"
)

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func tryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	recorder, ok := player.(gadget.TapeRecoder)
	if ok {
		recorder.Record()

	} else {
		fmt.Println("Player was not a TapeRecorder")
	}
}

func main() {
	var player Player
	player = gadget.TapePlayer{}
	mixtape := []string{"Banh my khong", "Tha vao mua", "Yeu duoc khong"}
	playList(player, mixtape)
	player = gadget.TapeRecoder{}
	playList(player, mixtape)
	tryOut(gadget.TapePlayer{})
	tryOut(gadget.TapeRecoder{})
}
