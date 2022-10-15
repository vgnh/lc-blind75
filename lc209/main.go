// https://leetcode.com/problems/minimum-size-subarray-sum/
/*

 */

package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	l, sum := 0, 0
	atLeastOne := false
	min := len(nums)
	for i := range nums {
		sum += nums[i]
		for sum >= target {
			if !atLeastOne {
				atLeastOne = true
			}
			if i+1-l < min {
				min = i + 1 - l
			}
			sum -= nums[l]
			l++
		}
	}
	if atLeastOne {
		return min
	} else {
		return 0
	}
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}
