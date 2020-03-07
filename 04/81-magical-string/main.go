package main

func main() {

}

func magicalString(n int) int {
	ds := []int{1, 2, 2}
	last := 2
	i := 2
	for len(ds) < n {
		last = 3 - last
		for j := 0; j < ds[i]; j++ {
			ds = append(ds, last)
		}
		i++
	}
	count := 0
	for _, d := range ds[:n] {
		if d == 1 {
			count++
		}
	}
	return count
}
