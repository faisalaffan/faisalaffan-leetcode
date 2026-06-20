# 1012 — Numbers With Repeated Digits

## Deskripsi

**Soal:** [1012. Numbers With Repeated Digits](https://leetcode.com/problems/numbers-with-repeated-digits/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP)

> **Ide Kunci:** Digit DP.

## Solusi Go

```go
package main

// LeetCode #1012: Numbers With Repeated Digits
// https://leetcode.com/problems/numbers-with-repeated-digits/
// Difficulty: Hard
//
// Approach: Digit DP.
//   Count numbers <= n with at least one repeated digit.
//   = n - count of numbers <= n with all unique digits (including 0 for counting ease).

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(numDupDigitsAtMostN(20))   // 1
	fmt.Println(numDupDigitsAtMostN(100))  // 10
	fmt.Println(numDupDigitsAtMostN(1000)) // 262
}

func numDupDigitsAtMostN(n int) int {
	s := strconv.Itoa(n)
	m := len(s)

	// Count numbers with all-unique digits from 0 to n (inclusive)
	// Use digit DP: tight + mask of used digits
	var countUnique func(pos int, tight bool, started bool, mask int) int
  // Membuat slice untuk menyimpan hasil
	memo := make([][1 << 10][2][2]int, m)
  // Iterasi seluruh elemen
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = [2][2]int{{-1, -1}, {-1, -1}}
		}
	}

	countUnique = func(pos int, tight bool, started bool, mask int) int {
		if pos == m {
			if started {
				return 1
			}
			return 0
		}
		t := 0
		if tight {
			t = 1
		}
		st := 0
		if started {
			st = 1
		}
		if memo[pos][mask][t][st] != -1 {
			return memo[pos][mask][t][st]
		}

		limit := 9
		if tight {
			limit = int(s[pos] - '0')
		}

		total := 0
		for d := 0; d <= limit; d++ {
			nextTight := tight && d == limit
			if !started {
				if d == 0 {
					total += countUnique(pos+1, nextTight, false, 0)
				} else {
					total += countUnique(pos+1, nextTight, true, 1<<d)
				}
			} else {
				if mask&(1<<d) != 0 {
					continue
				}
				total += countUnique(pos+1, nextTight, true, mask|(1<<d))
			}
		}
		memo[pos][mask][t][st] = total
		return total
	}

	unique := countUnique(0, true, false, 0)
	// unique excludes 0 (not started at the end). n - unique = count of numbers 1..n with repeats.
	return n - unique
}
```
