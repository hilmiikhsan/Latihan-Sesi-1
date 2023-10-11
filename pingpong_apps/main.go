package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Player struct {
	Name    string
	Hit     int
	IsLoser bool
}

func play(playerName string, player *Player, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		if player.IsLoser {
			return
		}

		counter := rand.Intn(100-1) + 1

		player.Hit++

		fmt.Printf("%s = Hit %d // counter %d\n", playerName, player.Hit, counter)

		if counter%11 == 0 {
			player.IsLoser = true
			fmt.Printf("%s kalah, total hit: %d, kalah di nomor %d\n", playerName, player.Hit, player.Hit)
			return
		}

		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	wg := sync.WaitGroup{}

	playerA := Player{
		Name: "Player A",
		Hit:  0,
	}

	playerB := Player{
		Name: "Player B",
		Hit:  0,
	}

	wg.Add(2)

	go play(playerA.Name, &playerA, &wg)
	go play(playerB.Name, &playerB, &wg)

	wg.Wait()
}
