# 2167 — Minimum Time To Remove All Cars Containing Illegal Goods

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func minimumTime(s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2167: Minimum Time to Remove All Cars Containing Illegal Goods
// https://leetcode.com/problems/minimum-time-to-remove-all-cars-containing-illegal-goods/
// Difficulty: Hard
//
// 3-state DP: processing left to right:
//   0 = removing from left end (cost 1 per car)
//   1 = middle section (cost 2 per '1', 0 per '0')
//   2 = removing from right end (cost 1 per car)
// Transitions: 0 -> 1 -> 2 (or 0 -> 2 directly).

import "fmt"

func main() {
	fmt.Println(minimumTime("1100101")) // 5
	fmt.Println(minimumTime("0010"))    // 2
	fmt.Println(minimumTime("010"))     // 2
	fmt.Println(minimumTime("111"))     // 3
	fmt.Println(minimumTime("0"))       // 0
}

func minimumTime(s string) int {
	const inf = 1 << 60
	l, m, r := 0, inf, inf
	for _, ch := range s {
		c := int(ch - '0')
		nl := l + 1
		nm := m + 2*c
		if l+2*c < nm {
			nm = l + 2*c
		}
		nr := r + 1
		if m+1 < nr {
			nr = m + 1
		}
		if l+1 < nr {
			nr = l + 1
		}
		l, m, r = nl, nm, nr
	}
	ans := l
	if m < ans {
		ans = m
	}
	if r < ans {
		ans = r
	}
	return ans
}

func MinimumTimeToRemoveAllCarsContainingIllegalGoods() any {
	return minimumTime("1100101")
}
```
