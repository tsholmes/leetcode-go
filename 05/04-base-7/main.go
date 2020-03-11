package main

func main() {

}

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	neg := num < 0
	if neg {
		num = -num
	}
	var res []byte
	for num > 0 {
		d := num % 7
		num = num / 7
		res = append(res, byte(d)+'0')
	}
	if neg {
		res = append(res, '-')
	}
	for i := 0; i < len(res)/2; i++ {
		j := len(res) - i - 1
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
