# 3816 — Lexicographically Smallest String After Deleting Duplicate Characters

## Deskripsi

**Soal:** [3816. Lexicographically Smallest String After Deleting Duplicate Characters](https://leetcode.com/problems/lexicographically-smallest-string-after-deleting-duplicate-characters/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Stack (tumpukan LIFO), Monotonic Stack (tumpukan monoton)

> **Ide Kunci:** Monotonic stack. Track last occurrence of each char.

## Solusi Go

```go
package main

// LeetCode #3816: Lexicographically Smallest String After Deleting
// Duplicate Characters
// https://leetcode.com/problems/lexicographically-smallest-string-after-deleting-duplicate-characters/
// Difficulty: Hard
//
// Delete duplicate characters from s, keeping each at most once,
// to obtain lexicographically smallest possible string.
//
// Approach: Monotonic stack. Track last occurrence of each char.
// Maintain stack with increasing characters. Pop if a larger char
// appears later.

import "fmt"

func main() {
	// Example 1
	fmt.Println(lexSmallestAfterDeletion("bcabc"))
	// Example 2
	fmt.Println(lexSmallestAfterDeletion("cbacdcbc"))
	// Edge: already unique
	fmt.Println(lexSmallestAfterDeletion("abc"))
	// Edge: reversed
	fmt.Println(lexSmallestAfterDeletion("cba"))
}

func lexSmallestAfterDeletion(s string) string {
  // Membuat slice untuk menyimpan hasil
	lastPos := make([]int, 26)
  // Iterasi seluruh elemen
	for i := range lastPos {
		lastPos[i] = -1
	}
	for i, ch := range s {
		lastPos[ch-'a'] = i
	}

  // Membuat slice untuk menyimpan hasil
	used := make([]bool, 26)
  // Membuat slice untuk menyimpan hasil
	stack := make([]byte, 0, len(s))

  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		c := s[i] - 'a'
		if used[c] {
			continue
		}
		for len(stack) > 0 && stack[len(stack)-1] > c && lastPos[stack[len(stack)-1]] > i {
			used[stack[len(stack)-1]] = false
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, c)
		used[c] = true
	}

  // Membuat slice untuk menyimpan hasil
	res := make([]byte, len(stack))
	for i, v := range stack {
		res[i] = v + 'a'
	}
	return string(res)
}
```
