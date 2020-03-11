package main

import "math"

func main() {

}

func fib(N int) int {
	s5 := math.Sqrt(5)

	a := (1 + s5) / 2
	b := (1 - s5) / 2

	fn := (math.Pow(a, float64(N)) - math.Pow(b, float64(N))) / s5
	return int(math.Round(fn))
}
