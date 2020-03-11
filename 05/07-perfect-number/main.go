package main

func main() {

}

func checkPerfectNumber(num int) bool {
	if num < 2 {
		return false
	}
	sum := 1
	for i := 2; i*i <= num; i++ {
		if num%i != 0 {
			continue
		}
		j := num / i
		sum += i
		if i != j {
			sum += j
		}
	}
	return sum == num
}
