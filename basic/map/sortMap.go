package _map

/**
 * For a list of votes, return an ordered set of candidate in descending order of their votes.
 */
// a vote 3 places
// 0  3points
// 1  2
// 2  1

// "A", "b", "c"
type vote struct {
	A string
	B string
	C string
}

// List<String> findWinner(List<Vote> votes)
func findWinner(votes []vote) []string {
	m := make(map[string]int)

	for i := 0; i < len(votes); i++ {
		m[votes[i].A] += 3
		m[votes[i].B] += 2
		m[votes[i].C] += 1
	}

	// [{k:v}, {k1:v1}]
	mm := make(map[int]map[int]string)
	count := 0
	for key, val := range m {
		pair := make(map[int]string)
		pair[val] = key
		mm[count] = pair
		count++
	}

	pairs := make([]map[int]string, 0)
	for _, v := range mm {
		pairs = append(pairs, v)
	}

	score1 := 0
	score2 := 0
	for i := 0; i < len(pairs); i++ {
		score1 = 0
		for k := range pairs[i] {
			score1 = k
		}
		for j := i + 1; j < len(pairs); j++ {
			score2 = 0
			for k := range pairs[j] {
				score2 = k
			}
			if score1 < score2 {
				tmp := pairs[i]
				pairs[i] = pairs[j]
				pairs[j] = tmp
			}
		}
	}
	res := make([]string, 0)
	for i := 0; i < len(pairs); i++ {
		for _, v := range pairs[i] {
			res = append(res, v)
		}
	}
	return res
}

//func main() {
//	vote1 := vote{
//		A: "a",
//		B: "b",
//		C: "c",
//	}
//	vote2 := vote{
//		A: "c",
//		B: "a",
//		C: "b",
//	}
//	votes := make([]vote, 0)
//	votes = append(votes, vote1)
//	votes = append(votes, vote2)
//	fmt.Println(findWinner(votes))
//}
