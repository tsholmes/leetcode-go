package main

func main() {

}

func decodeString(s string) string {
	var i int
	var decode func() (string, int)
	decode = func() (string, int) {
		if s[i] < '0' || s[i] > '9' {
			i++
			return s[i-1 : i], 1
		}
		var c int
		for s[i] >= '0' && s[i] <= '9' {
			c *= 10
			c += int(s[i] - '0')
			i++
		}
		// consume [
		i++
		var r []byte
		for s[i] != ']' {
			w, wc := decode()
			for j := 0; j < wc; j++ {
				r = append(r, w...)
			}
		}
		i++
		return string(r), c
	}
	var res []byte
	for i < len(s) {
		w, c := decode()
		for j := 0; j < c; j++ {
			res = append(res, w...)
		}
	}
	return string(res)
}
