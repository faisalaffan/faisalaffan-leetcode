# 3169 — Count Days Without Meetings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countDays(days int, meetings [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Merge Sort

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(log n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Merge Sort** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3169: Count Days Without Meetings
// https://leetcode.com/problems/count-days-without-meetings/
// Difficulty: Medium
// Time: O(n log n) | Space: O(log n)

import (
	"fmt"
	"sort"
)

func countDays(days int, meetings [][]int) int {
  // Custom sort dengan comparator
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

  // Alokasi slice integer
	merged := make([][2]int, 0)
	for _, m := range meetings {
		if len(merged) > 0 && m[0] <= merged[len(merged)-1][1]+1 {
			if m[1] > merged[len(merged)-1][1] {
				merged[len(merged)-1][1] = m[1]
			}
		} else {
			merged = append(merged, [2]int{m[0], m[1]})
		}
	}

	ans := days
	for _, m := range merged {
		ans -= m[1] - m[0] + 1
	}
	return ans
}

func main() {
	fmt.Println(countDays(10, [][]int{{5, 7}, {1, 3}, {9, 10}})) // Expected: 2
	fmt.Println(countDays(5, [][]int{{2, 4}, {1, 3}}))            // Expected: 1
	fmt.Println(countDays(6, [][]int{{1, 6}}))                    // Expected: 0
}
```
