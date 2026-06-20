# 3703 — Remove K Balanced Substrings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func removeKBalancedSubstrings(s string, k int) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Stack

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Stack** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3703: Remove K-Balanced Substrings
// https://leetcode.com/problems/remove-k-balanced-substrings/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

type pair struct {
	ch    byte
	count int
}

func removeKBalancedSubstrings(s string, k int) string {
	var stack []pair

  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if len(stack) > 0 && stack[len(stack)-1].ch == c {
			stack[len(stack)-1].count++
		} else {
			stack = append(stack, pair{ch: c, count: 1})
		}

		// Check for k-balanced pattern
		if c == ')' && stack[len(stack)-1].count == k {
			if len(stack) >= 2 {
				prev := &stack[len(stack)-2]
				if prev.ch == '(' && prev.count >= k {
					// Remove closing part
					stack = stack[:len(stack)-1]
					// Remove or reduce opening part
					if prev.count == k {
						stack = stack[:len(stack)-1]
					} else {
						prev.count -= k
					}
				}
			}
		}
	}

	res := make([]byte, 0, len(s))
	for _, p := range stack {
		for j := 0; j < p.count; j++ {
			res = append(res, p.ch)
		}
	}
	return string(res)
}

func main() {
	fmt.Println(removeKBalancedSubstrings("(())", 1))
	fmt.Println(removeKBalancedSubstrings("(()(", 1))
	fmt.Println(removeKBalancedSubstrings("((()))()()()", 3))
}
```
