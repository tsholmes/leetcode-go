package main

func main() {

}

func kLengthApart(nums []int, k int) bool {
	last := -k - 1
	for i, n := range nums {
		if n == 0 {
			continue
		}
		if i-last <= k {
			return false
		}
		last = i
	}
	return true
}
