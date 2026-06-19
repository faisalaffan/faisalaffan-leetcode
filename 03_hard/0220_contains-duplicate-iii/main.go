package main

// LeetCode #220: Contains Duplicate III
// https://leetcode.com/problems/contains-duplicate-iii/
// Difficulty: Hard

import "fmt"

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	if valueDiff < 0 || indexDiff <= 0 {
		return false
	}

	buckets := make(map[int]int)

	for i, num := range nums {
		bucketID := num / (valueDiff + 1)
		if num < 0 {
			bucketID--
		}

		if _, exists := buckets[bucketID]; exists {
			return true
		}
		if val, exists := buckets[bucketID-1]; exists && num-val <= valueDiff {
			return true
		}
		if val, exists := buckets[bucketID+1]; exists && val-num <= valueDiff {
			return true
		}

		buckets[bucketID] = num

		if i >= indexDiff {
			oldNum := nums[i-indexDiff]
			oldBucket := oldNum / (valueDiff + 1)
			if oldNum < 0 {
				oldBucket--
			}
			delete(buckets, oldBucket)
		}
	}

	return false
}

func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0))
}
