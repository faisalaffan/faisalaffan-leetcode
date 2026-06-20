# 1047 — Remove All Adjacent Duplicates In String

## Deskripsi

**Soal:** [1047. Remove All Adjacent Duplicates In String](https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Stack (tumpukan LIFO)

## Solusi Go

```go
package main

// LeetCode #1047: Remove All Adjacent Duplicates In String
// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(removeDuplicates("abbaca")) // "ca"
	fmt.Println(removeDuplicates("azxxzy")) // "ay"
	fmt.Println(removeDuplicates("a"))      // "a"
}

// LeetCode submission: removeDuplicates
func removeDuplicates(s string) string {
  // Membuat slice untuk menyimpan hasil
	stack := make([]byte, 0, len(s))
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
```
