package main

func main() {

}

func checkIfExist(arr []int) bool {
	for i, n := range arr {
		for j, n2 := range arr {
			if i != j && n == n2*2 {
				return true
			}
		}
	}
	return false
}
