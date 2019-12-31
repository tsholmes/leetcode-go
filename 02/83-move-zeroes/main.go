package main

func main() {

}

func moveZeroes(nums []int) {
	count := 0
	for _, n := range nums {
		if n != 0 {
			nums[count] = n
			count++
		}
	}
	for i := count; i < len(nums); i++ {
		nums[i] = 0
	}
}
