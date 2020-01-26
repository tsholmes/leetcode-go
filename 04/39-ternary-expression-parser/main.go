package main

import "fmt"

func main() {
	fmt.Println(parseTernary("T?2:3"))
	fmt.Println(parseTernary("F?1:T?4:5"))
	fmt.Println(parseTernary("T?T?F:5:3"))
}

func parseTernary(ex string) string {
	var get func() string
	get = func() string {
		var c string
		c, ex = ex[:1], ex[1:]
		if len(ex) == 0 || ex[0] != '?' {
			return c
		}
		ex = ex[1:] // ?
		lv := get()
		ex = ex[1:] // :
		rv := get()
		if c == "T" {
			return lv
		} else {
			return rv
		}
	}
	return get()
}
