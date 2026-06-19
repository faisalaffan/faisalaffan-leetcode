package main

// LeetCode #1: Two Sum
// https://leetcode.com/problems/two-sum/
// Difficulty: Easy
//
// Approaches (paste ONE twoSum function to LeetCode):
//   v1 HashMap:   Time O(n)   | Space O(n) | Mem 5-6MB | Runtime 0-4ms
//   v2 BruteForce: Time O(n²) | Space O(1) | Mem 4-5MB | Runtime 20-50ms
//   v3 UltraLow:  Time O(n²)  | Space O(1) | Mem 3-4MB | Runtime 15-30ms
//
// Catatan: Go runtime baseline ~2-3MB. Gak mungkin di bawah 2MB.
// LeetCode ukur RSS (resident set size) bukan heap aja.

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [0,1]
	fmt.Println(twoSum([]int{3, 2, 4}, 6))      // [1,2]
	fmt.Println(twoSum([]int{3, 3}, 6))         // [0,1]
}

// --- v1 HashMap: O(n) time, O(n) space ---
func twoSumHashMap(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for i, n := range nums {
		if j, ok := seen[target-n]; ok {
			return []int{j, i}
		}
		seen[n] = i
	}
	return nil
}

// --- v2 Brute Force: O(n²) time, O(1) space ---
func twoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// --- v3 Ultra Low Memory ---
// Hindari return nil (extra branch), minimalkan live vars, gunakan len(nums)
// dan akses indeks langsung tanpa variabel lokal tambahan.
func twoSum(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		vi := nums[i]
		want := target - vi
		for j := i + 1; j < n; j++ {
			if nums[j] == want {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
