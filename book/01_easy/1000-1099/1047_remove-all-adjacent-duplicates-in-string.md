# 1047 — Remove All Adjacent Duplicates In String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func removeDuplicates(s string) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Stack

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Stack** — sering muncul di interview!

## 💻 Solusi Go

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
	stack := make([]byte, 0, len(s))
  // Linear scan O(n)
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
