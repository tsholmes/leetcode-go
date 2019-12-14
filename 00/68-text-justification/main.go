package main

import "fmt"

func main() {
	do := func(maxWidth int, words ...string) {
		res := fullJustify(words, maxWidth)
		for _, line := range res {
			fmt.Println(line + "|")
		}
		fmt.Println()
	}
	do(16, "This", "is", "an", "example", "of", "text", "justification.")
	do(16, "What", "must", "be", "acknowledgment", "shall", "be")
	do(20, "Science", "is", "what", "we", "understand", "well", "enough", "to", "explain",
		"to", "a", "computer.", "Art", "is", "everything", "else", "we", "do")
}

func fullJustify(words []string, maxWidth int) []string {
	line := make([]byte, 0, maxWidth)
	var res []string
	for len(words) > 0 {
		count := 1
		length := len(words[0])
		for len(words) > count && length+1+len(words[count]) <= maxWidth {
			length += 1 + len(words[count])
			count++
		}
		ws := words[:count]
		words = words[count:]

		line = line[:0]
		// left justify
		if len(words) == 0 || count == 1 {
			for i, w := range ws {
				if i != 0 {
					line = append(line, ' ')
				}
				line = append(line, w...)
			}
			for len(line) < maxWidth {
				line = append(line, ' ')
			}
		} else {
			spaces := maxWidth - length + count - 1
			each := spaces / (count - 1)
			for i, w := range ws {
				if i > 0 {
					for j := 0; j < each; j++ {
						line = append(line, ' ')
					}
					if i <= spaces-each*(count-1) {
						line = append(line, ' ')
					}
				}
				line = append(line, w...)
			}
		}
		res = append(res, string(line))
	}

	return res
}
