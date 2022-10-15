package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func plusOne(digits []int) []int {
	b := big.NewInt(0)
	for _, v := range digits {
		b.Mul(b, big.NewInt(10))
		b.Add(b, big.NewInt(int64(v)))
	}
	b.Add(b, big.NewInt(1))
	s := b.String()
	n := make([]int, len(s))
	for i := range n {
		n[i], _ = strconv.Atoi(s[i : i+1])
	}
	return n
}

func main() {
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
}
