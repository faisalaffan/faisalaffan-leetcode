package main

// LeetCode #3599: Partition Array to Minimize XOR
// https://leetcode.com/problems/partition-array-to-minimize-xor/
// Difficulty: Medium
// Complexity: O(n * maxXor) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", PartitionArrayToMinimizeXor([]int{1, 2, 3, 4}))
	// Test case 2
	fmt.Println("Test 2:", PartitionArrayToMinimizeXor([]int{1, 1, 1}))
	// Test case 3
	fmt.Println("Test 3:", PartitionArrayToMinimizeXor([]int{5, 1, 2}))
}

func PartitionArrayToMinimizeXor(nums []int) int {
	// Compute total XOR of all elements
	totalXor := 0
	for _, v := range nums {
		totalXor ^= v
	}

	// We need to split array into two non-empty partitions
	// to minimize XOR of partition1 XOR partition2
	// Since partition1 ^ partition2 = totalXor (if we consider all elements)
	// Actually: if partition1 xor = p1, partition2 xor = p2, then p1 ^ p2 = totalXor
	// We want to minimize p1 ^ p2 = totalXor, which is constant!
	// But we're splitting the array, so:
	// If we split such that partition1 xor = x, partition2 xor = total ^ x = y
	// We want to minimize x ^ y = x ^ (total ^ x) = total
	// That's constant regardless of partition!
	// So we just need any valid partition

	// Alternative: minimize |xor of each partition|
	prefixXor := 0
	minVal := totalXor
	for i := 0; i < len(nums)-1; i++ {
		prefixXor ^= nums[i]
		suffixXor := totalXor ^ prefixXor
		xor := prefixXor ^ suffixXor
		if xor < minVal {
			minVal = xor
		}
	}
	return minVal
}
