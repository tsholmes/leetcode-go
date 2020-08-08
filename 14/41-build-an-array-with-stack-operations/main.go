package main

func main() {

}

func buildArray(target []int, n int) []string {
	var res []string
	next := 1
	for _, n := range target {
		for next != n {
			res = append(res, "Push", "Pop")
			next++
		}
		res = append(res, "Push")
		next++
	}
	return res
}
