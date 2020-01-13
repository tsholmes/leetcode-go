package main

func main() {

}

func superPow(a int, b []int) int {
	res := 1
	for i := len(b) - 1; i >= 0; i-- {
		for j := 0; j < b[i]; j-- {
			res = (res * a) % 1337
		}
		n := 1
		for j := 0; j < 10; j++ {
			n = (n * a) % 1337
		}
		a = n
	}
	return res
}
