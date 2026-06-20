# 0985 — Sum Of Even Numbers After Queries

## Deskripsi

**Soal:** [0985. Sum Of Even Numbers After Queries](https://leetcode.com/problems/sum-of-even-numbers-after-queries/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n + q) where n = len(nums), q = len(queries)  
**Kompleksitas Ruang:** O(1) excluding output

**Algoritma:** —

> **Ide Kunci:** Maintain running sum of even numbers, update incrementally

## Solusi Go

```go
package main

// LeetCode #985: Sum of Even Numbers After Queries
// https://leetcode.com/problems/sum-of-even-numbers-after-queries/
// Difficulty: Medium
//
// Approach: Maintain running sum of even numbers, update incrementally
// Time: O(n + q) where n = len(nums), q = len(queries)
// Space: O(1) excluding output

import "fmt"

func main() {
	fmt.Println(sumEvenAfterQueries([]int{1, 2, 3, 4}, [][]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}})) // [8,6,2,4]
	fmt.Println(sumEvenAfterQueries([]int{1}, [][]int{{4, 0}}))                                  // [0]
}

func sumEvenAfterQueries(nums []int, queries [][]int) []int {
  // Membuat slice untuk menyimpan hasil
	result := make([]int, len(queries))
	sum := 0

	for _, n := range nums {
		if n%2 == 0 {
			sum += n
		}
	}

	for i, q := range queries {
		val, idx := q[0], q[1]
		if nums[idx]%2 == 0 {
			sum -= nums[idx]
		}
		nums[idx] += val
		if nums[idx]%2 == 0 {
			sum += nums[idx]
		}
		result[i] = sum
	}

	return result
}
```
