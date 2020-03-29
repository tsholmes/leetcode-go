package main

func main() {

}

func sumFourDivisors(nums []int) int {
	var res int
	for _, n := range nums {
		s, c := divs(n)
		if c == 4 {
			res += s
		}
	}
	return res
}

func divs(n int) (int, int) {
	count := 0
	sum := 0
	for i := 1; i*i <= n; i++ {
		if n%i != 0 {
			continue
		}
		j := n / i
		count++
		sum += i
		if i != j {
			count++
			sum += j
		}
	}
	return sum, count
}
