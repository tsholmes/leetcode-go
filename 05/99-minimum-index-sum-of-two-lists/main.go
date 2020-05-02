package main

func main() {

}

func findRestaurant(list1 []string, list2 []string) []string {
	i1s, i2s := map[string]int{}, map[string]int{}
	for i, s := range list1 {
		i1s[s] = i
	}
	for i, s := range list2 {
		i2s[s] = i
	}

	minsum := len(list1) + len(list2)
	for s, i1 := range i1s {
		if i2, ok := i2s[s]; ok {
			if i1+i2 < minsum {
				minsum = i1 + i2
			}
		}
	}

	var res []string
	for s, i1 := range i1s {
		if i2, ok := i2s[s]; ok {
			if i1+i2 == minsum {
				res = append(res, s)
			}
		}
	}

	return res
}
