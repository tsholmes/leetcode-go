package main

func main() {

}

func fourSumCount(A []int, B []int, C []int, D []int) int {
	c1s := map[int]int{}
	for _, n := range A {
		for _, m := range B {
			c1s[n+m]++
		}
	}
	c2s := map[int]int{}
	for _, n := range C {
		for _, m := range D {
			c2s[n+m]++
		}
	}
	res := 0
	for k := range c1s {
		res += c1s[k] * c2s[-k]
	}
	return res
}
