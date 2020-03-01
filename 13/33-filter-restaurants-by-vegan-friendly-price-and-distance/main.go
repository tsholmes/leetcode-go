package main

import "sort"

func main() {

}
func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	var valid []int
	for i, r := range restaurants {
		if (veganFriendly == 0 || r[2] == 1) && (r[3] <= maxPrice) && (r[4] <= maxDistance) {
			valid = append(valid, i)
		}
	}
	sort.Slice(valid, func(i, j int) bool {
		ri, rj := restaurants[valid[i]], restaurants[valid[j]]
		if ri[1] > rj[1] {
			return true
		} else if ri[1] < rj[1] {
			return false
		} else if ri[0] > rj[0] {
			return true
		} else {
			return false
		}
	})

	res := make([]int, len(valid))
	for i, v := range valid {
		res[i] = restaurants[v][0]
	}
	return res
}
