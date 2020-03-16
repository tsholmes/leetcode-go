package main

func main() {

}

func killProcess(pid []int, ppid []int, kill int) []int {
	cs := map[int][]int{}
	for i := 0; i < len(pid); i++ {
		cs[ppid[i]] = append(cs[ppid[i]], pid[i])
	}
	var res []int
	var walk func(p int)
	walk = func(p int) {
		res = append(res, p)
		for _, pp := range cs[p] {
			walk(pp)
		}
	}
	walk(kill)
	return res
}
