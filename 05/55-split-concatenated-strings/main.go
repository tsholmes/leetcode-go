package main

func main() {

}

func splitLoopedString(strs []string) string {
	var maxc byte
	revs := make([]string, len(strs))
	maxs := make([]string, len(strs))
	for i, s := range strs {
		revs[i] = reverse(s)
		maxs[i] = strs[i]
		if revs[i] > strs[i] {
			maxs[i] = revs[i]
		}
		for j := 0; j < len(s); j++ {
			if s[j] > maxc {
				maxc = s[j]
			}
		}
	}

	var work []byte
	build := func(i int, ci int, rev bool) string {
		if rev {
			ci = len(strs[i]) - ci - 1
		}

		si := strs[i]
		if rev {
			si = revs[i]
		}

		work = work[:0]
		work = append(work, si[ci:]...)
		for j := (i + 1) % len(strs); j != i; j = (j + 1) % len(strs) {
			work = append(work, maxs[j]...)
		}
		work = append(work, si[:ci]...)
		return string(work)
	}

	max := ""
	for i, s := range strs {
		for j := 0; j < len(s); j++ {
			if s[j] == maxc {
				ss := build(i, j, false)
				if ss > max {
					max = ss
				}
				ss = build(i, j, true)
				if ss > max {
					max = ss
				}
			}
		}
	}

	return max
}

func reverse(s string) string {
	res := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		res = append(res, s[len(s)-i-1])
	}
	return string(res)
}
