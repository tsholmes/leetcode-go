package main

func main() {

}

func constructRectangle(area int) []int {
	var res [2]int
	for w := 1; w*w <= area; w++ {
		if area%w != 0 {
			continue
		}
		l := area / w
		res = [2]int{l, w}
	}
	return res[:]
}
