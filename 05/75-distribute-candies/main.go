package main

func main() {

}

func distributeCandies(candies []int) int {
	distinct := map[int]bool{}
	for _, c := range candies {
		distinct[c] = true
	}

	res := len(distinct)
	if res > len(candies)/2 {
		res = len(candies) / 2
	}
	return res
}
