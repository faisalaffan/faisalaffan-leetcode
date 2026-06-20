# 2031 — Count Subarrays With More Ones Than Zeros

## Deskripsi

**Soal:** [2031. Count Subarrays With More Ones Than Zeros](https://leetcode.com/problems/count-subarrays-with-more-ones-than-zeros/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Fenwick Tree (Binary Indexed Tree), Segment Tree (pohon segmen)

**Fungsi Solusi:** `func subarraysWithMoreOnesThanZeros(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #2031: Count Subarrays With More Ones Than Zeros
// https://leetcode.com/problems/count-subarrays-with-more-ones-than-zeros/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func subarraysWithMoreOnesThanZeros(nums []int) int {
	// Treat 0 as -1, 1 as +1
	// prefix[j] - prefix[i] > 0 means subarray (i,j] has more ones
	// prefix[j] > prefix[i] -> count how many previous prefixes are smaller
	prefix := 0
	count := 0
	// Fenwick tree or segment tree for range counts, but simpler:
	// offset by n since prefix ranges from -n to n
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	tree := make([]int, 2*n+2)
	mod := int(1e9 + 7)

	add := func(idx int) {
		idx += n + 1
		for idx < len(tree) {
			tree[idx]++
			idx += idx & -idx
		}
	}

	sum := func(idx int) int {
		idx += n + 1
		res := 0
		for idx > 0 {
			res += tree[idx]
			idx -= idx & -idx
		}
		return res
	}

	add(0) // empty prefix
	for _, v := range nums {
		if v == 1 {
			prefix++
		} else {
			prefix--
		}
		// Count how many previous prefixes are less than current prefix
		// sum(prefix-1) = count of prefixes <= prefix-1
		count = (count + sum(prefix-1)) % mod
		add(prefix)
	}

	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", subarraysWithMoreOnesThanZeros([]int{0, 1, 1, 0, 1}))
	// Expected: 9

	// Test case 2
	fmt.Println("Test 2:", subarraysWithMoreOnesThanZeros([]int{1, 0, 1}))
	// Expected: 3

	// Test case 3
	fmt.Println("Test 3:", subarraysWithMoreOnesThanZeros([]int{1, 1, 1}))
	// Expected: 6
}
```
