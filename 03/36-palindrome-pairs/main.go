package main

func main() {

}

func palindromePairs(words []string) [][]int {
	var res [][]int
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if j == i {
				continue
			}
			l1, l2 := len(words[i]), len(words[j])
			ok := true
			for m := 0; m < (l1+l2)/2; m++ {
				n := (l1 + l2) - m - 1
				var a, b byte
				if m < l1 {
					a = words[i][m]
				} else {
					a = words[j][m-l1]
				}
				if n < l1 {
					b = words[i][n]
				} else {
					b = words[j][n-l1]
				}
				if a != b {
					ok = false
					break
				}
			}
			if ok {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}
