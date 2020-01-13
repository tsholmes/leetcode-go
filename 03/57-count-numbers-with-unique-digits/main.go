package main

func main() {

}

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	res := 10
	cur := 9
	for i := 2; i <= n; i++ {
		cur *= 11 - i
		res += cur
	}
	return res
}
