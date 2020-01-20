package main

func main() {

}

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	cs := [16]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

	num = num & 0xFFFFFFFF

	var res []byte
	for num != 0 {
		res = append(res, cs[num&0xF])
		num = num >> 4
	}
	for i := 0; i < len(res)/2; i++ {
		j := len(res) - i - 1
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
