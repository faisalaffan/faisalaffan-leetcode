# 0930 — Binary Subarrays With Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func numSubarraysWithSum(nums []int, goal int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Sliding Window

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #930: Binary Subarrays With Sum
// https://leetcode.com/problems/binary-subarrays-with-sum/
// Difficulty: Medium

import "fmt"

// Time: O(n) | Space: O(1)
func numSubarraysWithSum(nums []int, goal int) int {
	// sliding window for sum <= goal, then subtract sum < goal
	atMost := func(g int) int {
		if g < 0 {
			return 0
		}
		sum, cnt, left := 0, 0, 0
		for right, v := range nums {
			sum += v
			for sum > g {
				sum -= nums[left]
				left++
			}
			cnt += right - left + 1
		}
		return cnt
	}
	return atMost(goal) - atMost(goal-1)
}

func main() {
	fmt.Println(numSubarraysWithSum([]int{1, 0, 1, 0, 1}, 2))
	fmt.Println(numSubarraysWithSum([]int{0, 0, 0, 0, 0}, 0))
}
```
