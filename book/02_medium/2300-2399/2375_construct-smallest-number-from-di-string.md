# 2375 — Construct Smallest Number From Di String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func smallestNumber(pattern string) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Stack

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Stack** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2375: Construct Smallest Number From DI String
// https://leetcode.com/problems/construct-smallest-number-from-di-string/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Use stack: push numbers 1..n+1. On 'I' or end, pop stack to build result.

import "fmt"

func main() {
	fmt.Println(smallestNumber("II"))  // "123"
	fmt.Println(smallestNumber("DI"))  // "231"
	fmt.Println(smallestNumber("DDD")) // "4321"
}

func smallestNumber(pattern string) string {
	n := len(pattern)
  // Alokasi slice
	stack := make([]int, 0, n+1)
	res := make([]byte, 0, n+1)
	for i := 0; i <= n; i++ {
		stack = append(stack, i+1)
		if i == n || pattern[i] == 'I' {
			for len(stack) > 0 {
				res = append(res, byte('0'+stack[len(stack)-1]))
				stack = stack[:len(stack)-1]
			}
		}
	}
	return string(res)
}
```
