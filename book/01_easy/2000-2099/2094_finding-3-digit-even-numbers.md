# 2094 — Finding 3 Digit Even Numbers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func FindingThreeDigitEvenNumbers(digits []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n^3), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

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
  // HashMap: O(1) lookup
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

  // Alokasi slice
	result := make([]int, 0, len(set))
	for v := range set {
		result = append(result, v)
	}
  // Sort O(n log n)
	sort.Ints(result)
	return result
}
```
