package main

func main() {

}

func isConvex(points [][]int) bool {
	allPos := true
	allNeg := true

	vec := func(i int) [2]int {
		j := (i + 1) % len(points)
		return [2]int{
			points[j][0] - points[i][0],
			points[j][1] - points[i][1],
		}
	}

	cross := func(a, b [2]int) int {
		return a[0]*b[1] - a[1]*b[0]
	}

	for i := range points {
		a, b := vec(i), vec((i+1)%len(points))
		c := cross(a, b)
		if c == 0 {
			continue
		}
		pos := c > 0
		allPos = allPos && pos
		allNeg = allNeg && !pos
	}

	return allPos || allNeg
}
