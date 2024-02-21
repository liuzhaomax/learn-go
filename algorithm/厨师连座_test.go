package algorithm

import (
	"fmt"
	"testing"
)

func TestChef(t *testing.T) {
	game := Game{}
	player1 := Player{State: State{Evil: false}}
	player2 := Player{State: State{Evil: false}}
	player3 := Player{State: State{Evil: true}}
	player4 := Player{State: State{Evil: true}}
	game.Players = []Player{player1, player2, player3, player4}
	connected := chef(game)
	fmt.Println(connected)
}
