package main

func main() {

}

/**
 * The read4 API is already defined for you.
 *
 *     read4 := func(buf []byte) int
 *
 * // Below is an example of how the read4 API can be called.
 * file := File("abcdefghijk") // File is "abcdefghijk", initially file pointer (fp) points to 'a'
 * buf := make([]byte, 4) // Create buffer with enough space to store characters
 * read4(buf) // read4 returns 4. Now buf = ['a','b','c','d'], fp points to 'e'
 * read4(buf) // read4 returns 4. Now buf = ['e','f','g','h'], fp points to 'i'
 * read4(buf) // read4 returns 3. Now buf = ['i','j','k',...], fp points to end of file
 */
var solution = func(read4 func([]byte) int) func([]byte, int) int {
	var rbuf [4]byte
	var rem []byte
	readRem := func(buf []byte, n int) int {
		cp := len(rem)
		if cp > n {
			cp = n
		}
		copy(buf, rem)
		rem = rem[cp:]
		return cp
	}
	done := false
	fillRem := func() {
		if done {
			return
		}
		if len(rem) > 0 {
			return
		}
		r := read4(rbuf[:])
		rem = rbuf[:r]
		done = r < 4
	}
	// implement read below.
	return func(buf []byte, n int) int {
		read := 0
		for n > 0 && (!done || len(rem) > 0) {
			fillRem()
			r := readRem(buf, n)
			buf = buf[r:]
			n -= r
			read += r
		}
		return read
	}
}
