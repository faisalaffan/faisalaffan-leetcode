# 2406 — Divide Intervals Into Minimum Number Of Groups

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func minGroups(intervals [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2406: Divide Intervals Into Minimum Number of Groups
// https://leetcode.com/problems/divide-intervals-into-minimum-number-of-groups/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)
// Sweep line: count concurrent intervals, answer is max concurrency.

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minGroups([][]int{{5, 10}, {6, 8}, {1, 5}, {2, 3}, {1, 10}})) // 3
	fmt.Println(minGroups([][]int{{1, 3}, {5, 6}, {8, 10}, {11, 13}}))        // 1
}

func minGroups(intervals [][]int) int {
  // Alokasi slice integer
	events := make([][2]int, 0, len(intervals)*2)
	for _, iv := range intervals {
		events = append(events, [2]int{iv[0], 1})   // start
		events = append(events, [2]int{iv[1] + 1, -1}) // end (inclusive)
	}

  // Custom sort dengan comparator
	sort.Slice(events, func(i, j int) bool {
		if events[i][0] != events[j][0] {
			return events[i][0] < events[j][0]
		}
		return events[i][1] < events[j][1]
	})

	cur, ans := 0, 0
	for _, e := range events {
		cur += e[1]
		if cur > ans {
			ans = cur
		}
	}
	return ans
}
```
