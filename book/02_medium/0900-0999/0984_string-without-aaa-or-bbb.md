# 0984 — String Without Aaa Or Bbb

## Deskripsi

**Soal:** [0984. String Without Aaa Or Bbb](https://leetcode.com/problems/string-without-aaa-or-bbb/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(a + b)  
**Kompleksitas Ruang:** O(a + b) for output

**Algoritma:** Greedy (pemilihan optimal lokal)

> **Ide Kunci:** Greedy - always append the character with more remaining count,

## Solusi Go

```go
package main

// LeetCode #984: String Without AAA or BBB
// https://leetcode.com/problems/string-without-aaa-or-bbb/
// Difficulty: Medium
//
// Approach: Greedy - always append the character with more remaining count,
//           but skip a triple by using the other character if we just wrote two.
// Time: O(a + b)
// Space: O(a + b) for output

import "fmt"

func main() {
	fmt.Println(strWithout3a3b(1, 2)) // "bba" or "bab"
	fmt.Println(strWithout3a3b(4, 1)) // "aabaa"
	fmt.Println(strWithout3a3b(3, 3)) // "ababa" or similar
}

func strWithout3a3b(a int, b int) string {
  // Membuat slice untuk menyimpan hasil
	result := make([]byte, 0, a+b)

	for a > 0 || b > 0 {
		writeA := false
		if a > b {
			writeA = true
		} else if a < b {
			writeA = false
		} else {
			// equal counts: prefer the one that won't create triple
			if len(result) >= 2 && result[len(result)-1] == 'a' && result[len(result)-2] == 'a' {
				writeA = false
			} else if len(result) >= 2 && result[len(result)-1] == 'b' && result[len(result)-2] == 'b' {
				writeA = true
			} else {
				writeA = true
			}
		}

		if writeA {
			result = append(result, 'a')
			a--
			if a > 0 && a >= b {
				result = append(result, 'a')
				a--
			}
		} else {
			result = append(result, 'b')
			b--
			if b > 0 && b >= a {
				result = append(result, 'b')
				b--
			}
		}
	}

	return string(result)
}
```
