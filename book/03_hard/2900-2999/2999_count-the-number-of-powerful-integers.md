# 2999 — Count The Number Of Powerful Integers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func numberOfPowerfulInt(start, finish int64, limit int, s string) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, DP, Prefix Sum

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2999: Count the Number of Powerful Integers
// https://leetcode.com/problems/count-the-number-of-powerful-integers/
// Difficulty: Hard
//
// Count integers in [start, finish] such that:
//   1. Every digit in the integer is <= limit.
//   2. The integer ends with the suffix s.
//
// Approach: Digit DP
//   countLe(X) counts numbers in [1, X] satisfying the conditions.
//   Answer = countLe(finish) - countLe(start-1).

import (
	"fmt"
	"strconv"
	"strings"
)

func numberOfPowerfulInt(start, finish int64, limit int, s string) int64 {
	suffixVal, _ := strconv.ParseInt(s, 10, 64)

	countLe := func(x int64) int64 {
		if x < suffixVal {
			return 0
		}
		sx := strconv.FormatInt(x, 10)
		n := len(sx)
		m := len(s)

		// If x has fewer digits than suffix, impossible
		if n < m {
			return 0
		}

		// If same number of digits, just compare strings
		if n == m {
			if sx >= s {
				return 1
			}
			return 0
		}

		// Prefix digits count = n - m
		preLen := n - m

		// Memoized DFS over the prefix positions
		type state struct {
			pos   int
			tight bool
		}
  // HashMap: O(1) lookup
		memo := make(map[state]int64)

		var dfs func(pos int, tight bool) int64
		dfs = func(pos int, tight bool) int64 {
			if pos == preLen {
				if !tight {
					return 1
				}
				// Compare the suffix portion of x with s
				if sx[preLen:] >= s {
					return 1
				}
				return 0
			}
			key := state{pos, tight}
			if v, ok := memo[key]; ok {
				return v
			}
			var res int64
			up := 9
			if tight {
				up = int(sx[pos] - '0')
			}
			if up > limit {
				up = limit
			}
			for d := 0; d <= up; d++ {
				res += dfs(pos+1, tight && d == int(sx[pos]-'0'))
			}
			memo[key] = res
			return res
		}

		ans := dfs(0, true)

		// Count numbers with fewer than n digits (but at least m+1 digits)
		// For a number with `length` digits: first digit 1..limit, rest 0..limit
		for length := m + 1; length < n; length++ {
			preLen2 := length - m
			ways := int64(limit)
			if limit > 9 {
				ways = 9
			}
			if limit >= 1 {
				for i := 1; i < preLen2; i++ {
					ways *= int64(limit + 1)
				}
				ans += ways
			}
		}

		// Handle leading zeros in prefix for numbers with exactly n digits.
		// The DFS above allowed leading zeros in the prefix which is incorrect
		// for numbers with n digits. We need to subtract the case where the
		// prefix is all zeros (which would make the number have m digits).
		// Actually for the DP, leading zeros in the prefix are fine because
		// we are counting numbers with exactly n digits: the first digit
		// of the prefix can be zero in the DP, but we should exclude the
		// all-zero prefix case (since that would be an m-digit number).
		// Subtract 1 if the prefix can be all zeros and is counted.
		preStr := sx[:preLen]
		if preStr > strings.Repeat("0", preLen) {
			// The all-zero prefix case was counted, subtract it.
			// Actually, if all-zero prefix would give a valid number (its suffix >= s),
			// then it's an m-digit number already counted separately.
			// We need to check: was the all-zero prefix counted in DFS?
			// Yes, if tight is false at that point. Let's just subtract.
			// The all-zero prefix case corresponds to: prefix = "0...0", tight becomes false
			// after the first non-zero digit. So it IS counted in the DP for n-digit numbers.
			// Subtract it.
			if s[0] != '0' {
				// The all-zero prefix means the number is just the suffix (length = m).
				// This was already counted by the length==m case above if suffix <= x.
				if sx[preLen:] >= s {
					ans--
				}
			}
		}

		return ans
	}

	return countLe(finish) - countLe(start-1)
}

func main() {
	// Example 1: [1,6000], limit=4, suffix="124"
	// Powerful integers <= 6000 ending with 124, each digit <= 4:
	// 124, 2124, 3124, 4124 -> 4 (6000 is excluded because suffix "6000" != "124")
	fmt.Println("Test 1:", numberOfPowerfulInt(1, 6000, 4, "124"))

	// Example 2: [15,215], limit=6, suffix="10"
	// 110, 210 -> 2
	fmt.Println("Test 2:", numberOfPowerfulInt(15, 215, 6, "10"))

	// Single digit range
	fmt.Println("Test 3:", numberOfPowerfulInt(1, 10, 9, "5"))

	// Range with small suffix
	fmt.Println("Test 4:", numberOfPowerfulInt(1, 100, 5, "0"))

	// Large range test
	fmt.Println("Test 5:", numberOfPowerfulInt(1, 1000000, 7, "77"))

	// Start > 1
	fmt.Println("Test 6:", numberOfPowerfulInt(100, 200, 9, "99"))

	// limit < 9
	fmt.Println("Test 7:", numberOfPowerfulInt(1, 500, 2, "1"))
}
```
