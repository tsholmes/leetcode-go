package main

import "fmt"

func main() {
	c := Codec{}
	fmt.Println(c.Decode(c.Encode([]string{
		"Hello", "World",
	})))
	fmt.Println(c.Decode(c.Encode([]string{
		"01234567ABCDEF01234567ABCDEF01234567ABCDEF01234567ABCDEF01234567ABCDEF01234567ABCDEF01234567ABCDEF01234567ABCDEF",
		"01234567ABCDEF01234567ABCDEF01234567ABCDEF01234567ABCDEF01234567ABCDEF01234567ABCDEF01234567ABCDEF01234567ABCDEF",
		"01234567ABCDEF01234567ABCDEF01234567ABCDEF01234567ABCDEF01234567ABCDEF01234567ABCDEF01234567ABCDEF01234567ABCDEF",
	})))
}

type Codec struct{}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	var res []byte

	for _, s := range strs {
		l := len(s)
		for l > 127 {
			res = append(res, byte(0x80|(l&0x7F)))
			l = l >> 7
		}
		res = append(res, byte(l))
		res = append(res, s...)
	}

	return string(res)
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	var res []string

	for len(strs) > 0 {
		c := strs[0]
		strs = strs[1:]
		l := 0
		var shift uint
		for c > 127 {
			l |= int(c&0x7f) << shift
			shift += 7
			c = strs[0]
			strs = strs[1:]
		}
		l |= int(c) << shift

		res = append(res, strs[:l])
		strs = strs[l:]
	}

	return res
}

// Your Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs));
