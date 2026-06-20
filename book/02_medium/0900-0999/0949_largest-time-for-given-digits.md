# 0949 — Largest Time For Given Digits

## Deskripsi

**Soal:** [0949. Largest Time For Given Digits](https://leetcode.com/problems/largest-time-for-given-digits/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func largestTimeFromDigits(arr []int) string`

## Solusi Go

```go
package main

// LeetCode #949: Largest Time for Given Digits
// https://leetcode.com/problems/largest-time-for-given-digits/
// Difficulty: Medium

import "fmt"

// Time: O(1) | Space: O(1)
func largestTimeFromDigits(arr []int) string {
	ans := -1
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				if i != j && j != k && i != k {
					l := 6 - i - j - k
					h := arr[i]*10 + arr[j]
					m := arr[k]*10 + arr[l]
					if h < 24 && m < 60 {
						if t := h*60 + m; t > ans {
							ans = t
						}
					}
				}
			}
		}
	}
	if ans < 0 {
		return ""
	}
	return fmt.Sprintf("%02d:%02d", ans/60, ans%60)
}

func main() {
	fmt.Println(largestTimeFromDigits([]int{1, 2, 3, 4}))
	fmt.Println(largestTimeFromDigits([]int{5, 5, 5, 5}))
	fmt.Println(largestTimeFromDigits([]int{2, 0, 6, 6}))
}
```
