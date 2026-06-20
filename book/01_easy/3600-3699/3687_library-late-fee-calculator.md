# 3687 — Library Late Fee Calculator

## Deskripsi

**Soal:** [3687. Library Late Fee Calculator](https://leetcode.com/problems/library-late-fee-calculator/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3687: Library Late Fee Calculator
// https://leetcode.com/problems/library-late-fee-calculator/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(LibraryLateFeeCalculator([]int{5, 1, 7}))
	fmt.Println(LibraryLateFeeCalculator([]int{1, 1}))
}

// Time: O(n)
// Space: O(1)
func LibraryLateFeeCalculator(daysLate []int) int {
	ans := 0
	for _, x := range daysLate {
		if x == 1 {
			ans += 1
		} else if x > 5 {
			ans += 3 * x
		} else {
			ans += 2 * x
		}
	}
	return ans
}
```
