package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("   -42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("9223372036854775808"))
}

func myAtoi(str string) int {
	for len(str) > 0 && strings.ContainsAny(str[:1], " \t\r\n") {
		str = str[1:]
	}
	if len(str) == 0 {
		return 0
	}

	neg := false
	if str[0] == '-' {
		neg = true
		str = str[1:]
	} else if str[0] == '+' {
		str = str[1:]
	}

	v := 0
	for len(str) > 0 && str[0] >= '0' && str[0] <= '9' {
		v *= 10
		v += int(str[0] - '0')
		if v > math.MaxInt32 {
			if neg {
				return math.MinInt32
			}
			return math.MaxInt32
		}
		str = str[1:]
	}

	if neg {
		v = -v
		if v < math.MinInt32 {
			return math.MinInt32
		}
	} else if v > math.MaxInt32 {
		return math.MaxInt32
	}

	return v
}
