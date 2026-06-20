# 2094 — Finding 3 Digit Even Numbers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindingThreeDigitEvenNumbers(digits []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n^3), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2094: Finding 3-Digit Even Numbers
// https://leetcode.com/problems/finding-3-digit-even-numbers/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(FindingThreeDigitEvenNumbers([]int{2, 1, 3, 0}))       // [102 120 130 132 210 230 302 310 312 320]
	fmt.Println(FindingThreeDigitEvenNumbers([]int{2, 2, 8, 8, 2}))    // [222 228 282 288 822 828 882]
	fmt.Println(FindingThreeDigitEvenNumbers([]int{0, 0, 0}))          // []
}

// Time: O(n^3), Space: O(1)
func FindingThreeDigitEvenNumbers(digits []int) []int {
  // Membuat map (HashMap) — pencarian O(1)
	set := make(map[int]bool)
	n := len(digits)

	for i := 0; i < n; i++ {
		if digits[i] == 0 {
			continue
		}
		for j := 0; j < n; j++ {
			if j == i {
				continue
			}
			for k := 0; k < n; k++ {
				if k == i || k == j {
					continue
				}
				num := digits[i]*100 + digits[j]*10 + digits[k]
				if num%2 == 0 {
					set[num] = true
				}
			}
		}
	}

  // Alokasi slice integer
	result := make([]int, 0, len(set))
	for v := range set {
		result = append(result, v)
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(result)
	return result
}
```
