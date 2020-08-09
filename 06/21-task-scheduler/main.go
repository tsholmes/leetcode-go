package main

func main() {

}

func leastInterval(tasks []byte, n int) int {
	var counts [26]int
	for _, t := range tasks {
		counts[t-'A']++
	}

	var nexts [26]int

	t := 0
	rem := len(tasks)
	for rem > 0 {
		maxc := 0
		maxi := -1
		for i := 0; i < 26; i++ {
			if counts[i] > maxc && nexts[i] <= t {
				maxc = counts[i]
				maxi = i
			}
		}
		t++
		if maxi == -1 {
			continue
		}
		counts[maxi]--
		rem--
		nexts[maxi] = t + n
	}

	return t
}
