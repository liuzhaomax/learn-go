package algorithm

import (
	"fmt"
	"testing"
)

func TestEmpath(t *testing.T) {
	game := Game{}
	player0 := Player{State: State{Evil: true, Dead: true}, Index: 0}
	player1 := Player{State: State{Evil: false, Dead: false}, Index: 1}
	player2 := Player{State: State{Evil: false, Dead: true}, Index: 2}
	player3 := Player{State: State{Evil: true, Dead: false}, Index: 3}
	player4 := Player{State: State{Evil: false, Dead: false}, Index: 4}
	game.Players = []Player{player0, player1, player2, player3, player4}
	evilNum := empath(game, 1)
	fmt.Println(evilNum)
}
