# 1411 — Number Of Ways To Paint N 3 Grid

## Deskripsi

**Soal:** [1411. Number Of Ways To Paint N 3 Grid](https://leetcode.com/problems/number-of-ways-to-paint-n-3-grid/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func numOfWays(n int) int`

## Solusi Go

```go
package main

// LeetCode #1411: Number of Ways to Paint N × 3 Grid
// https://leetcode.com/problems/number-of-ways-to-paint-n-3-grid/
// Difficulty: Hard

import "fmt"

const mod1411 = 1_000_000_007

func numOfWays(n int) int {
	// Two pattern types for a 3-column row:
	// Pattern "ABA": 3 colors, first and third same (6 ways: 3*2)
	// Pattern "ABC": 3 colors, all different (6 ways: 3*2*1)
	aba, abc := 6, 6
	for i := 2; i <= n; i++ {
		// ABA can transition to:
		//   ABA: 3 ways (middle different from both ends)
		//   ABC: 2 ways (middle same as first, third different)
		// ABC can transition to:
		//   ABA: 2 ways (first and third same, middle different)
		//   ABC: 2 ways (all different, no color repeats position)
		newAba := (3*aba + 2*abc) % mod1411
		newAbc := (2*aba + 2*abc) % mod1411
		aba, abc = newAba, newAbc
	}
	return (aba + abc) % mod1411
}

func main() {
	// Example: n=1 -> 12
	fmt.Println(numOfWays(1))
}
```
