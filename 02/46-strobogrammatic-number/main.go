package main

func main() {

}

func isStrobogrammatic(num string) bool {
	rev := map[byte]byte{
		'0': '0',
		'1': '1',
		'6': '9',
		'8': '8',
		'9': '6',
	}
	for i := 0; i < (len(num)+1)/2; i++ {
		j := len(num) - i - 1
		if num[i] != rev[num[j]] {
			return false
		}
	}
	return true
}
