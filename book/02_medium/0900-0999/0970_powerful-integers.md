# 0970 — Powerful Integers

## Deskripsi

**Soal:** [0970. Powerful Integers](https://leetcode.com/problems/powerful-integers/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(log_x(bound) * log_y(bound))  
**Kompleksitas Ruang:** O(log_x(bound) * log_y(bound))

**Algoritma:** —

**Fungsi Solusi:** `func powerfulIntegers(x int, y int, bound int) []int`

## Solusi Go

```go
package main

// LeetCode #970: Powerful Integers
// https://leetcode.com/problems/powerful-integers/
// Difficulty: Medium

import "fmt"

// Time: O(log_x(bound) * log_y(bound)) | Space: O(log_x(bound) * log_y(bound))
func powerfulIntegers(x int, y int, bound int) []int {
  // Membuat map untuk pencarian O(1): key → value
	seen := make(map[int]bool)

	for a := 1; a <= bound; a *= x {
		for b := 1; a+b <= bound; b *= y {
			seen[a+b] = true
			if y == 1 {
				break
			}
		}
		if x == 1 {
			break
		}
	}

  // Membuat slice untuk menyimpan hasil
	ans := make([]int, 0, len(seen))
	for v := range seen {
		ans = append(ans, v)
	}
	return ans
}

func main() {
	fmt.Println(powerfulIntegers(2, 3, 10))
	fmt.Println(powerfulIntegers(3, 5, 15))
	fmt.Println(powerfulIntegers(2, 1, 10))
}
```
