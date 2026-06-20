# 1521 — Find A Value Of A Mysterious Function Closest To Target

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func closestToTarget(arr []int, target int) int
```

> **💡 Hint:** Track all possible AND values of subarrays ending at each

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1521: Find a Value of a Mysterious Function Closest to Target
// https://leetcode.com/problems/find-a-value-of-a-mysterious-function-closest-to-target/
// Difficulty: Hard
//
// Winston has a mysterious function func(arr, l, r) that returns the
// bitwise AND of all elements in arr[l..r]. Find the minimum absolute
// difference between any func value and target.
//
// Approach: Track all possible AND values of subarrays ending at each
// position. AND values only decrease, so the set of distinct values
// is small (at most 32 per position).

import "fmt"

func main() {
	// Example 1
	fmt.Println(closestToTarget([]int{9, 12, 3, 7, 15}, 5))
	// Example 2
	fmt.Println(closestToTarget([]int{1000000, 1000000, 1000000}, 1))
	// Edge: single element
	fmt.Println(closestToTarget([]int{5}, 5))
}

func closestToTarget(arr []int, target int) int {
	ans := abs(arr[0] - target)
	pre := map[int]bool{arr[0]: true}
	for _, x := range arr {
		cur := map[int]bool{x: true}
		for y := range pre {
			cur[x&y] = true
		}
		for y := range cur {
			ans = min(ans, abs(y-target))
		}
		pre = cur
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
