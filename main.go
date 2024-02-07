package main

import (
	"fmt"
)

func letterCombinations(digits string) []string {
	l := len(digits)
	if l == 0 {
		return []string{}
	}

	res := make([]string, 0)

	for _, d := range digits {
		base := d - '2'
		var runes []string
		if base < 0 {
			continue
		} else if base >= 5 {
			if base == 5 {
				runes = []string{"p", "q", "r", "s"}
			} else if base == 6 {
				runes = []string{"t", "u", "v"}
			} else {
				runes = []string{"w", "x", "y", "z"}
			}
		} else {
			offset := base * 3
			for i := 0; i < 3; i++ {
				runes = append(runes, string('a'+offset+rune(i)))
			}
		}

		if len(res) == 0 {
			res = append(res, runes...)
		} else {
			var new_res []string
			for _, s := range res {
				for _, r := range runes {
					new_res = append(new_res, s+r)
				}
			}
			res = new_res
		}

	}

	return res
}

func main() {
	fmt.Println(letterCombinations("234"))
}
