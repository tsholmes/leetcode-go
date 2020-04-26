package main

func main() {

}

func minNumberOfFrogs(croakOfFrogs string) int {
	cis := map[byte]int{
		'c': 0,
		'r': 1,
		'o': 2,
		'a': 3,
		'k': 4,
	}

	var counts [4]int
	var max int
	var inflight int

	for _, c := range []byte(croakOfFrogs) {
		idx := cis[c]
		if idx == 0 {
			inflight++
			if inflight > max {
				max = inflight
			}
			counts[0]++
			continue
		}
		if counts[idx-1] == 0 {
			return -1
		}
		counts[idx-1]--
		if idx == 4 {
			inflight--
		} else {
			counts[idx]++
		}
	}
	for i := 0; i < 4; i++ {
		if counts[i] != 0 {
			return -1
		}
	}
	return max
}
