# 1717 — Maximum Score From Removing Substrings

## Deskripsi

**Soal:** [1717. Maximum Score From Removing Substrings](https://leetcode.com/problems/maximum-score-from-removing-substrings/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Stack (tumpukan LIFO)

**Fungsi Solusi:** `func maximumGain(s string, x int, y int) int`

## Solusi Go

```go
package main

// LeetCode #1717: Maximum Score From Removing Substrings
// https://leetcode.com/problems/maximum-score-from-removing-substrings/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

func maximumGain(s string, x int, y int) int {
	// Ensure we always process the higher-scoring pair first
	if y > x {
		s = reverse(s)
		x, y = y, x
	}

	ans := 0

	// First pass: remove "ab" for x points
  // Membuat slice untuk menyimpan hasil
	stack := make([]byte, 0, len(s))
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == 'a' && s[i] == 'b' {
			stack = stack[:len(stack)-1]
			ans += x
		} else {
			stack = append(stack, s[i])
		}
	}

	// Second pass: remove "ba" for y points from remaining
  // Membuat slice untuk menyimpan hasil
	stack2 := make([]byte, 0, len(stack))
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(stack); i++ {
		if len(stack2) > 0 && stack2[len(stack2)-1] == 'b' && stack[i] == 'a' {
			stack2 = stack2[:len(stack2)-1]
			ans += y
		} else {
			stack2 = append(stack2, stack[i])
		}
	}

	return ans
}

func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func main() {
	fmt.Println(maximumGain("cdbcbbaaabab", 4, 5)) // Expected: 19
	fmt.Println(maximumGain("aabbaaxybbaabb", 5, 4)) // Expected: 20
	fmt.Println(maximumGain("ab", 1, 2)) // Expected: 1
}
```
