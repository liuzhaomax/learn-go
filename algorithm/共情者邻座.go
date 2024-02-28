package algorithm

type Game struct {
	Players []Player
}

type Player struct {
	State State
	Index int
}

type State struct {
	Evil bool
	Dead bool
}

func empath(game Game, empathIndex int) int {
	evilQuantity := 0
	var player Player
	for i := range game.Players {
		if i == empathIndex {
			player = game.Players[i]
			break
		}
	}
	// 算法开始
	var left int
	var right int
	if player.Index == 0 {
		left = len(game.Players) - 1
		right = player.Index + 1
	} else if player.Index == len(game.Players)-1 {
		left = player.Index - 1
		right = 0
	} else {
		left = player.Index - 1
		right = player.Index + 1
	}
	var leftPrev = -1
	var rightPrev = -1
	for {
		if leftPrev != left {
			leftPrev = left
			if game.Players[left].State.Dead {
				left--
				if left < 0 {
					left = len(game.Players) - 1
				}
			} else {
				if game.Players[left].State.Evil {
					evilQuantity += 1
				}
			}
		}
		if rightPrev != right {
			rightPrev = right
			if game.Players[right].State.Dead {
				right++
				if right > len(game.Players)-1 {
					right = 0
				}
			} else {
				if game.Players[right].State.Evil {
					evilQuantity += 1
				}
			}
		}
		if leftPrev == left && rightPrev == right {
			break
		}
	}
	return evilQuantity
}
