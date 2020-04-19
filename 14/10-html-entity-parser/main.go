package main

import "strings"

func main() {

}

func entityParser(text string) string {
	reps := map[string]string{
		"quot;":  "\"",
		"apos;":  "'",
		"amp;":   "&",
		"gt;":    ">",
		"lt;":    "<",
		"frasl;": "/",
	}

	var res []byte
	for len(text) > 0 {
		c := text[0]
		text = text[1:]
		if c != '&' {
			res = append(res, c)
			continue
		}
		match := false
		for k, v := range reps {
			if strings.HasPrefix(text, k) {
				text = text[len(k):]
				res = append(res, v...)
				match = true
				break
			}
		}
		if !match {
			res = append(res, c)
		}
	}
	return string(res)
}
