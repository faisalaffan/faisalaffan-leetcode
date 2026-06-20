# 2059 — Minimum Operations To Convert Number

## Deskripsi

**Soal:** [2059. Minimum Operations To Convert Number](https://leetcode.com/problems/minimum-operations-to-convert-number/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * range)  
**Kompleksitas Ruang:** O(range)

**Algoritma:** Queue (antrian FIFO)

**Fungsi Solusi:** `func minimumOperations(nums []int, start int, goal int) int`

## Solusi Go

```go
package main

// LeetCode #2059: Minimum Operations to Convert Number
// https://leetcode.com/problems/minimum-operations-to-convert-number/
// Difficulty: Medium
// Time: O(n * range) | Space: O(range)

import "fmt"

func minimumOperations(nums []int, start int, goal int) int {
  // Membuat slice untuk menyimpan hasil
	visited := make([]bool, 1001)
	queue := []int{start}
	steps := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[i]
			if curr == goal {
				return steps
			}
			for _, v := range nums {
				for _, next := range []int{curr + v, curr - v, curr ^ v} {
					if next == goal {
						return steps + 1
					}
					if next >= 0 && next <= 1000 && !visited[next] {
						visited[next] = true
						queue = append(queue, next)
					}
				}
			}
		}
		queue = queue[size:]
		steps++
	}

	return -1
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minimumOperations([]int{1, 3}, 6, 4))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", minimumOperations([]int{2, 4, 8}, 3, 10))
	// Expected: 3

	// Test case 3
	fmt.Println("Test 3:", minimumOperations([]int{1}, 0, 1001))
	// Expected: -1 (goal outside range, unreachable)

	// Test case 4
	fmt.Println("Test 4:", minimumOperations([]int{2, 8, 16}, 0, 1))
	// Expected: -1
}
```
