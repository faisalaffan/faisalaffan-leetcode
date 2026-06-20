# 3177 — Find The Maximum Length Of A Good Subsequence Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumLength(nums []int, k int) int
```

> **💡 Hint:** DP tracking best[val][k] and overall best[k].

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3177: Find the Maximum Length of a Good Subsequence II
// https://leetcode.com/problems/find-the-maximum-length-of-a-good-subsequence-ii/
// Difficulty: Hard
//
// A subsequence is "good" if at most k adjacent pairs have different values.
// Find the maximum possible length of a good subsequence.
//
// Approach: DP tracking best[val][k] and overall best[k].
//   best[val][k] = max length of good subsequence ending with value val
//                  using at most k diff-pairs.
//   global[k]     = max over all best[val][k].
//
// For each element v, for each kk:
//   len = max(best[v][kk] + 1, (kk>0 ? global[kk-1] + 1 : 1))

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumLength(nums []int, k int) int {
  // Edge case: input kosong — langsung return
	if len(nums) == 0 {
		return 0
	}
	// bestSame[val][kk] = max length ending with val using at most kk diffs.
  // Membuat map (HashMap) — pencarian O(1)
	bestSame := make(map[int][]int)
	// global[kk] = overall max length using at most kk diffs.
  // Alokasi slice integer
	global := make([]int, k+1)

	for _, v := range nums {
		if bestSame[v] == nil {
			bestSame[v] = make([]int, k+1)
		}
		row := bestSame[v]
		// Use temporary slice to avoid using updated values within the same
		// iteration (we need the state before processing this element).
  // Alokasi slice integer
		newBest := make([]int, k+1)
		copy(newBest, row)

		for kk := 0; kk <= k; kk++ {
			cur := 1
			if row[kk] > 0 {
				cur = max(cur, row[kk]+1)
			}
			if kk > 0 && global[kk-1] > 0 {
				cur = max(cur, global[kk-1]+1)
			}
			newBest[kk] = max(newBest[kk], cur)
			global[kk] = max(global[kk], cur)
		}
		bestSame[v] = newBest
	}

	ans := 0
	for kk := 0; kk <= k; kk++ {
		ans = max(ans, global[kk])
	}
	return ans
}

func main() {
	fmt.Println(maximumLength([]int{1, 2, 1, 1, 3}, 2)) // expect 4 (e.g. [1,2,1,1])
	fmt.Println(maximumLength([]int{1, 2, 3, 4, 5}, 1)) // expect 2
}
```
