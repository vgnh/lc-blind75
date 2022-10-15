// https://leetcode.com/problems/container-with-most-water/
// Dynamic Programming
/*
   Start from the ends.
   Move the shortest of the two ends forward. If start < end, start++ else end--.
   Keep calculating most volume.
*/

package main

import (
	"fmt"
	"math"
)

func maxArea(height []int) int {
	start, end := 0, len(height)-1
	most := int(math.Min(float64(height[start]), float64(height[end])) * float64(end-start))
	for start < end {
		if height[start] < height[end] {
			start++
		} else {
			end--
		}
		vol := int(math.Min(float64(height[start]), float64(height[end])) * float64(end-start))
		if vol > most {
			most = vol
		}
	}
	return most
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
