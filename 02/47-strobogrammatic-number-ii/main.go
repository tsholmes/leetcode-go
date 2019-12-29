package main

func main() {

}

func findStrobogrammatic(n int) []string {
	if n == 0 {
		return nil
	}
	work := make([]byte, n)
	var res []string
	var search func(int, int)
	same := []byte(`018`)
	anyl := []byte(`1689`)
	anyr := []byte(`1986`)
	search = func(i, j int) {
		if i > j {
			res = append(res, string(work))
			return
		}
		if i == j {
			for _, c := range same {
				work[i] = c
				search(i+1, j-1)
			}
		} else {
			for k := range anyl {
				work[i] = anyl[k]
				work[j] = anyr[k]
				search(i+1, j-1)
			}
			if i != 0 {
				work[i] = '0'
				work[j] = '0'
				search(i+1, j-1)
			}
		}
	}
	search(0, n-1)
	return res
}
