# 2457 — Minimum Addition To Make Integer Beautiful

## Deskripsi

**Soal:** [2457. Minimum Addition To Make Integer Beautiful](https://leetcode.com/problems/minimum-addition-to-make-integer-beautiful/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(log n * log n)  
**Kompleksitas Ruang:** O(log n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2457: Minimum Addition to Make Integer Beautiful
// https://leetcode.com/problems/minimum-addition-to-make-integer-beautiful/
// Difficulty: Medium
// Time: O(log n * log n) | Space: O(log n)
// Try rounding up to higher digit positions until digit sum <= target.

import "fmt"

func main() {
	fmt.Println(makeIntegerBeautiful(16, 6))   // 4 (16+4=20, digit sum 2 <= 6)
	fmt.Println(makeIntegerBeautiful(467, 6))  // 33 (467+33=500, digit sum 5 <= 6)
	fmt.Println(makeIntegerBeautiful(1, 1))    // 0
}

func makeIntegerBeautiful(n int64, target int) int64 {
	if digitSum(n) <= target {
		return 0
	}

	pow10 := int64(10)
	for {
		next := ((n / pow10) + 1) * pow10
		if digitSum(next) <= target {
			return next - n
		}
		pow10 *= 10
	}
}

func digitSum(n int64) int {
	sum := 0
	for n > 0 {
		sum += int(n % 10)
		n /= 10
	}
	return sum
}
```
