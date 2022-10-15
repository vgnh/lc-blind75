// https://leetcode.com/problems/longest-palindromic-substring/
// Dynamic Programming
/*
   Odd Palindromes
   Start from Substring(i, i), <- Substring(L, R)
   Keep decreasing L, increasing R
   Store the longest as long as charAt(L) == charAt(R)
   Even Palindromes
   Start from Substring(i, i+1), <- Substring(L, R)
   Keep decreasing L, increasing R
   Store the longest as long as charAt(L) == charAt(R)
*/

package main

import "fmt"

func longestPalindrome(s string) string {
	longest := ""
	for i := range s {
		// Odd palindromes
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > len(longest) {
				longest = s[l : r+1]
			}
			l--
			r++
		}

		// Even palindromes
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > len(longest) {
				longest = s[l : r+1]
			}
			l--
			r++
		}
	}
	return longest
}

func main() {
	fmt.Println(longestPalindrome("racecar"))
}
