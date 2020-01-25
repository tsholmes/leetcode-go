package main

func main() {

}

func originalDigits(s string) string {
	counts := map[byte]int{}
	for i := 0; i < len(s); i++ {
		counts[s[i]]++
	}

	handle := func(c byte, full string) int {
		count := counts[c]
		for i := 0; i < len(full); i++ {
			counts[full[i]] -= count
		}
		return count
	}

	var digits [10]int

	digits[8] = handle('g', "eight")
	digits[3] = handle('h', "three")
	digits[2] = handle('t', "two")
	digits[4] = handle('u', "four")
	digits[5] = handle('f', "five")
	digits[0] = handle('z', "zero")
	digits[1] = handle('o', "one")
	digits[7] = handle('v', "seven")
	digits[9] = handle('e', "nine")
	digits[6] = handle('x', "six")

	var res []byte
	for i := 0; i < len(digits); i++ {
		for j := 0; j < digits[i]; j++ {
			res = append(res, byte(i)+'0')
		}
	}

	return string(res)
}
