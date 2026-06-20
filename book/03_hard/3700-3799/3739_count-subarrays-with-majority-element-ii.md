# 3739 — Count Subarrays With Majority Element Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func countMajoritySubarrays(nums []int, target int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum, Fenwick Tree

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3739: Count Subarrays With Majority Element II
// https://leetcode.com/problems/count-subarrays-with-majority-element-ii/
// Difficulty: Hard
//
// Count subarrays where target appears more than half the length.
//
// Approach: Transform condition to f[i] = 2*cnt[i] - i where cnt[i]
// is prefix count of target. Subarray [l, r] has target majority iff
// f[r+1] > f[l]. Use Fenwick tree to count such pairs.
//
// f[i] ranges from -n to n, offset by n for 0-based indexing.

import "fmt"

func main() {
	// Example 1
	fmt.Println(countMajoritySubarrays([]int{1, 2, 2, 1, 2}, 2))
	// Example 2
	fmt.Println(countMajoritySubarrays([]int{1, 1, 1}, 1))
	// Edge: target not majority anywhere
	fmt.Println(countMajoritySubarrays([]int{1, 2, 3}, 4))
	// Edge: single element
	fmt.Println(countMajoritySubarrays([]int{5}, 5))
}

func countMajoritySubarrays(nums []int, target int) int64 {
	n := len(nums)
	// f[i] = 2*cnt[i] - i where cnt[i] = number of target in first i elements
	// f[0] = 0 (cnt[0]=0, i=0)
	// f ranges from -n to n, so we offset by n

	offset := n
	size := 2*n + 3
  // Alokasi slice
	bit := make([]int, size+1)

	add := func(idx, val int) {
		idx++ // 1-indexed BIT
		for idx <= size {
			bit[idx] += val
			idx += idx & -idx
		}
	}

	sum := func(idx int) int {
		idx++
		s := 0
		for idx > 0 {
			s += bit[idx]
			idx -= idx & -idx
		}
		return s
	}

	var ans int64
	cnt := 0

	// f[0] = 0
	add(0+offset, 1)

	for i := 1; i <= n; i++ {
		if nums[i-1] == target {
			cnt++
		}
		fi := 2*cnt - i
		// Count how many f[l] < fi (l from 0 to i-1)
		// sum(fi-1) = count of indices with f < fi
		ans += int64(sum(fi - 1 + offset))
		add(fi+offset, 1)
	}

	return ans
}
```
