package debounce

//100  10s
// 10   1s

var cnt = make(map[int]int)

func check(timeLimit int, reqLimit int, customerID int, timeStamp int) bool {
	// after timeD  cnt -> 0
	//timer = setTimeout(callback, 10)
	now := 15
	if now-timeStamp > timeLimit {
		cnt[customerID] = 0
	}

	cnt[customerID] += 1

	// calc cnt > reqLimit -> false
	if cnt[customerID] > reqLimit {
		return false
	} else {
		return true
	}
}

func req(cnt int) int {
	return cnt + 1
}

//func main() {
//	checked1 := check(10, 4, 1, 1)
//	checked1 = check(10, 4, 1, 2)
//	checked1 = check(10, 4, 1, 3)
//	checked1 = check(10, 4, 1, 4)
//	// 10
//	checked1 = check(10, 4, 1, 5)
//	checked2 := check(10, 4, 2, 6)
//	fmt.Println(checked1)
//	fmt.Println(checked2)
//}
