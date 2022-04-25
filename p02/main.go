// https://leetcode.com/problems/longest-substring-without-repeating-characters
//
/*

 */

package main

import (
	"fmt"
	"math"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	start := 0
	max := 0
	for i := 0; i < len(s)-1; i++ {
		s2 := s[start : i+1]
		if strings.Contains(s2, string(s[i+1])) {
			start += strings.LastIndex(s2, string(s[i+1])) + 1
		} else {
			max = int(math.Max(float64(len(s2)+1), float64(max)))
		}
	}

	return int(math.Max(float64(max), 1))
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
