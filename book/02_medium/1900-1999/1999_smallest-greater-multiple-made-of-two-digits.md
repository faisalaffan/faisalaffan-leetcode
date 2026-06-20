# 1999 — Smallest Greater Multiple Made Of Two Digits

## Deskripsi

**Soal:** [1999. Smallest Greater Multiple Made Of Two Digits](https://leetcode.com/problems/smallest-greater-multiple-made-of-two-digits/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(2^k * log n) where k is number of digits, Space: O(2^k)  
**Kompleksitas Ruang:** O(2^k)

**Algoritma:** BFS (Breadth-First Search / pencarian lebar)

## Solusi Go

```go
package main

// LeetCode #1999: Smallest Greater Multiple Made of Two Digits
// https://leetcode.com/problems/smallest-greater-multiple-made-of-two-digits/
// Difficulty: Medium [Paid]

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(SmallestGreaterMultipleMadeOfTwoDigits(2, 0, 2))
	fmt.Println(SmallestGreaterMultipleMadeOfTwoDigits(8, 2, 4))
	fmt.Println(SmallestGreaterMultipleMadeOfTwoDigits(2, 3, 9))
}

// Time: O(2^k * log n) where k is number of digits, Space: O(2^k)
func SmallestGreaterMultipleMadeOfTwoDigits(k int, digit1 int, digit2 int) int {
	if digit1 > digit2 {
		digit1, digit2 = digit2, digit1
	}

  // Membuat slice untuk menyimpan hasil
	nums := make([]int, 0)

	// BFS to generate all numbers using only digit1 and digit2
	if digit1 != 0 {
		nums = append(nums, digit1)
	}
	if digit2 != 0 && digit2 != digit1 {
		nums = append(nums, digit2)
	}

	queue := nums[:]
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr > math.MaxInt32/10 {
			continue
		}
		n1 := curr*10 + digit1
		if n1 <= math.MaxInt32 {
			nums = append(nums, n1)
			queue = append(queue, n1)
		}
		if digit2 != digit1 {
			n2 := curr*10 + digit2
			if n2 <= math.MaxInt32 {
				nums = append(nums, n2)
				queue = append(queue, n2)
			}
		}
	}

	sort.Ints(nums)

	for _, num := range nums {
		if num >= k && num%k == 0 {
			return num
		}
	}
	return -1
}
```
