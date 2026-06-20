# 1910 — Remove All Occurrences Of A Substring

## Deskripsi

**Soal:** [1910. Remove All Occurrences Of A Substring](https://leetcode.com/problems/remove-all-occurrences-of-a-substring/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n*m) where n = len(s), m = len(part), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Stack (tumpukan LIFO)

## Solusi Go

```go
package main

// LeetCode #1910: Remove All Occurrences of a Substring
// https://leetcode.com/problems/remove-all-occurrences-of-a-substring/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(RemoveOccurrences("daabcbaabcbc", "abc"))
	fmt.Println(RemoveOccurrences("axxxxyyyyb", "xy"))
	fmt.Println(RemoveOccurrences("aabababa", "aba"))
}

// Time: O(n*m) where n = len(s), m = len(part), Space: O(n)
func RemoveOccurrences(s string, part string) string {
  // Membuat slice untuk menyimpan hasil
	stack := make([]byte, 0)
	m := len(part)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		stack = append(stack, s[i])
		if len(stack) >= m && string(stack[len(stack)-m:]) == part {
			stack = stack[:len(stack)-m]
		}
	}
	return string(stack)
}
```
