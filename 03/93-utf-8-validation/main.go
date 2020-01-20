package main

func main() {

}

func validUtf8(data []int) bool {
	for len(data) > 0 {
		ec := 0
		if data[0]&0x80 == 0 {
			// noop
		} else if data[0]>>5 == 6 {
			ec = 1
		} else if data[0]>>4 == 14 {
			ec = 2
		} else if data[0]>>3 == 30 {
			ec = 3
		} else {
			return false
		}
		data = data[1:]
		if len(data) < ec {
			return false
		}
		for _, b := range data[:ec] {
			if b>>6 != 2 {
				return false
			}
		}
		data = data[ec:]
	}
	return true
}
