package main

func main() {

}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	cs := map[int][]int{}
	for i, m := range manager {
		cs[m] = append(cs[m], i)
	}

	var search func(i int) int
	search = func(i int) int {
		max := 0
		for _, c := range cs[i] {
			v := search(c)
			if v > max {
				max = v
			}
		}
		return max + informTime[i]
	}
	return search(headID)
}
