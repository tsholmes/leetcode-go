package main

func main() {

}
func xorQueries(arr []int, queries [][]int) []int {
	upTo := make([]int, len(arr))
	upTo[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		upTo[i] = upTo[i-1] ^ arr[i]
	}

	var res []int
	for _, q := range queries {
		if q[0] == 0 {
			res = append(res, upTo[q[1]])
		} else {
			res = append(res, upTo[q[0]-1]^upTo[q[1]])
		}
	}
	return res
}
