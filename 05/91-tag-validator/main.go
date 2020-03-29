package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(isValid("<DIV>This is the first line <![CDATA[<div>]]></DIV>"))
	fmt.Println(isValid("<![CDATA[wahaha]]]><![CDATA[]> wahaha]]>"))
}

func isValid(code string) bool {
	if len(code) == 0 {
		return false
	}

	var stack []string

	take := func() byte {
		v := code[0]
		code = code[1:]
		return v
	}

	parseTag := func() bool {
		if take() != '<' {
			return false
		}
		end := strings.IndexByte(code, '>')
		if end == -1 {
			return false
		}
		name := code[:end]
		code = code[end+1:]

		if len(name) < 1 {
			return false
		}

		closet := false
		if name[0] == '/' {
			closet = true
			name = name[1:]
		}
		if len(name) < 1 || len(name) > 9 {
			return false
		}
		for _, c := range name {
			if c < 'A' || c > 'Z' {
				return false
			}
		}

		if closet {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != name {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, name)
		}

		return true
	}

	if !parseTag() {
		return false
	}

	for len(code) > 0 {
		if code[0] != '<' {
			code = code[1:]
			continue
		}
		if strings.HasPrefix(code, "<![CDATA[") {
			code = code[len("<![CDATA["):]
			endIdx := strings.Index(code, "]]>")
			if endIdx == -1 {
				return false
			}
			code = code[endIdx+len("]]>"):]
		} else {
			if !parseTag() {
				return false
			}
			if len(stack) == 0 && len(code) > 0 {
				return false
			}
		}
	}

	return len(stack) == 0
}
