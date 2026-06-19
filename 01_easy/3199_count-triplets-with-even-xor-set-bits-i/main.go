package main

// LeetCode #3199: Count Triplets with Even XOR Set Bits I
// https://leetcode.com/problems/count-triplets-with-even-xor-set-bits-i/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(CountTripletsWithEvenXorSetBitsI([]int{1}, []int{2}, []int{3}))
	fmt.Println(CountTripletsWithEvenXorSetBitsI([]int{1, 2}, []int{3, 4}, []int{5, 6}))
}

// popcount returns the number of set bits in x.
func popcount(x int) int {
	count := 0
	for x > 0 {
		count += x & 1
		x >>= 1
	}
	return count
}

// CountTripletsWithEvenXorSetBitsI counts triplets (i, j, k) where a[i] XOR b[j] XOR c[k] has even number of set bits.
// Time: O(|a| * |b| * |c|). Space: O(1).
func CountTripletsWithEvenXorSetBitsI(a []int, b []int, c []int) int {
	count := 0
	for _, va := range a {
		for _, vb := range b {
			for _, vc := range c {
				if popcount(va^vb^vc)%2 == 0 {
					count++
				}
			}
		}
	}
	return count
}
