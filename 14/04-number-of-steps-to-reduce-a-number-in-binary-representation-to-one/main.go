package main

func main() {

}

func numSteps(s string) int {
	bs := []byte(s)
	count := 0
	for len(bs) > 1 {
		count++
		if bs[len(bs)-1] == '0' {
			bs = bs[:len(bs)-1]
			continue
		}
		end := 1
		for end < len(bs) && bs[len(bs)-end] == '1' {
			end++
		}
		if end == len(bs) {
			bs = make([]byte, len(bs)+1)
			bs[0] = '1'
			for i := 1; i < len(bs); i++ {
				bs[i] = '0'
			}
		} else {
			bs[len(bs)-end] = '1'
			for i := len(bs) - end + 1; i < len(bs); i++ {
				bs[i] = '0'
			}
		}
	}
	return count
}
