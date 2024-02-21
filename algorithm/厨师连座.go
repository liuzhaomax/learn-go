package algorithm

type Game struct {
	Players []Player
}

type Player struct {
	State State
}

type State struct {
	Evil bool
}

func chef(game Game) int {
	connected := 0
	meetEvil := false
	meetEvilAgain := false
	for i, player := range game.Players {
		if player.State.Evil {
			if meetEvil {
				meetEvilAgain = true
			}
			if meetEvilAgain {
				connected += 1
				meetEvilAgain = false
			}
			if i == len(game.Players)-1 && game.Players[0].State.Evil {
				connected += 1
				break
			}
			meetEvil = true
		} else {
			meetEvil = false
		}
	}
	return connected
}
