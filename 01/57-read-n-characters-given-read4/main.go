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
	// implement read below.
	return func(buf []byte, n int) int {
		var rd [4]byte
		var read int
		for n > 0 {
			cr := read4(rd[:])
			cp := cr
			if n < cp {
				cp = n
			}
			copy(buf, rd[:cp])
			buf = buf[cp:]
			read += cp
			if cr < 4 {
				break
			}
			n -= cp
		}
		return read
	}
}
