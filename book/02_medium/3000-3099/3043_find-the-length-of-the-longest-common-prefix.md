# 3043 — Find The Length Of The Longest Common Prefix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func longestCommonPrefix(arr1 []int, arr2 []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Prefix Sum

**Kompleksitas Waktu:** O(n*logM + m*logM)  
**Kompleksitas Ruang:** O(n*logM)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3043: Find the Length of the Longest Common Prefix
// https://leetcode.com/problems/find-the-length-of-the-longest-common-prefix/
// Difficulty: Medium
// Time: O(n*logM + m*logM) | Space: O(n*logM)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(longestCommonPrefix([]int{1, 10, 100}, []int{1000}))
	fmt.Println(longestCommonPrefix([]int{1, 2, 3}, []int{4, 4, 4}))
}

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	prefixes := map[int]bool{}
	for _, x := range arr1 {
		for x > 0 {
			prefixes[x] = true
			x /= 10
		}
	}
	ans := 0
	for _, x := range arr2 {
		for x > 0 {
			if prefixes[x] {
				if len(strconv.Itoa(x)) > ans {
					ans = len(strconv.Itoa(x))
				}
				break
			}
			x /= 10
		}
	}
	return ans
}
```
