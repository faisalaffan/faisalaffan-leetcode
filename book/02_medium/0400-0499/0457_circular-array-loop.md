# 0457 — Circular Array Loop

## Deskripsi

**Soal:** [0457. Circular Array Loop](https://leetcode.com/problems/circular-array-loop/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func circularArrayLoop(nums []int) bool`

## Solusi Go

```go
package main

// LeetCode #457: Circular Array Loop
// https://leetcode.com/problems/circular-array-loop/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func circularArrayLoop(nums []int) bool {
	n := len(nums)

	next := func(i int) int {
		return ((i+nums[i])%n + n) % n
	}

	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			continue
		}

		slow, fast := i, i
		for {
			slow = next(slow)
			fast = next(next(fast))
			if slow == fast {
				break
			}
		}

		// Found cycle, check direction and length
		if next(slow) != slow && nums[slow]*nums[next(slow)] > 0 {
			length := 1
			cur := next(slow)
			for cur != slow {
				length++
				cur = next(cur)
			}
			if length > 1 {
				return true
			}
		}

		// Mark visited
		cur := i
		for nums[cur]*nums[next(cur)] > 0 {
			tmp := next(cur)
			nums[cur] = 0
			cur = tmp
		}
	}
	return false
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", circularArrayLoop([]int{2, -1, 1, 2, 2}))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", circularArrayLoop([]int{-1, -2, -3, -4, -5, 6}))
	// Expected: false

	// Test case 3
	fmt.Println("Test 3:", circularArrayLoop([]int{1, -1, 2}))
	// Expected: true
}
```
