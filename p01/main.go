// https://leetcode.com/problems/two-sum
//
/*
 */

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numIndexMap := make(map[int]int)
	for i := range nums {
		complement := target - nums[i]
		if _, ok := numIndexMap[complement]; ok {
			return []int{numIndexMap[complement], i}
		}
		numIndexMap[nums[i]] = i
	}
	panic("No two sum solution")
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
