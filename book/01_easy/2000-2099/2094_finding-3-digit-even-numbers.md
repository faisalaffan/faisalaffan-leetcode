# 2094 — Finding 3 Digit Even Numbers

## Deskripsi

**Soal:** [2094. Finding 3 Digit Even Numbers](https://leetcode.com/problems/finding-3-digit-even-numbers/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n^3), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

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
  // Membuat map untuk pencarian O(1): key → value
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

  // Membuat slice untuk menyimpan hasil
	result := make([]int, 0, len(set))
	for v := range set {
		result = append(result, v)
	}
	sort.Ints(result)
	return result
}
```
