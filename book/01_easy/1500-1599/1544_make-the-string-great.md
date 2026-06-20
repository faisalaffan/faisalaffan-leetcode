# 1544 — Make The String Great

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func makeGood(s string) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Stack

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Stack** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1544: Make The String Great
// https://leetcode.com/problems/make-the-string-great/
// Difficulty: Easy
//
// LeetCode submission: func makeGood(s string) string

import "fmt"

func main() {
	fmt.Println(MakeTheStringGreat("leEeetcode")) // "leetcode"
	fmt.Println(MakeTheStringGreat("abBAcC"))     // ""
	fmt.Println(MakeTheStringGreat("s"))          // "s"
}

// Time: O(n), Space: O(n)
func MakeTheStringGreat(s string) string {
	stack := make([]byte, 0, len(s))
  // Range loop
	for i := range s {
		stack = append(stack, s[i])
		n := len(stack)
		if n >= 2 {
			diff := int(stack[n-1]) - int(stack[n-2])
			if diff == 32 || diff == -32 {
				stack = stack[:n-2]
			}
		}
	}
	return string(stack)
}
```
