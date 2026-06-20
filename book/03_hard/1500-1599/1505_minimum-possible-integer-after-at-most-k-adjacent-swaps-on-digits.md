# 1505 — Minimum Possible Integer After At Most K Adjacent Swaps On Digits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func newFenwick(n int) *fenwick
```

> **💡 Hint:** Fenwick Tree (BIT) + Greedy

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, BFS, Fenwick Tree (BIT)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1505: Minimum Possible Integer After at Most K Adjacent Swaps On Digits
// https://leetcode.com/problems/minimum-possible-integer-after-at-most-k-adjacent-swaps-on-digits/
// Difficulty: Hard
//
// Approach: Fenwick Tree (BIT) + Greedy
// We process positions left-to-right. For each position, we want the smallest
// possible digit that can be moved here within k swaps.
// Use a Fenwick tree to track how many positions have been removed (shifted left).
// Maintain queues of positions for each digit 0-9.
// For each position i:
//   - Try digits 0-9 in order.
//   - For each digit with available positions, compute the cost to bring it
//     to position i: cost = position - i + (number of removed positions before it).
//   - If cost <= k, take it, update k, mark position as removed, break.

import (
	"fmt"
	"strings"
)

func main() {
	// Example 1
	fmt.Println(minInteger("4321", 4))
	// Expected: "1342"

	// Example 2
	fmt.Println(minInteger("100", 1))
	// Expected: "010"

	// Example 3
	fmt.Println(minInteger("36789", 3))
	// Expected: "36789"
}

type fenwick struct {
	tree []int
	n    int
}

func newFenwick(n int) *fenwick {
	return &fenwick{tree: make([]int, n+2), n: n}
}

func (f *fenwick) add(idx int, delta int) {
	idx++
	for idx <= f.n {
		f.tree[idx] += delta
		idx += idx & -idx
	}
}

func (f *fenwick) sum(idx int) int {
	idx++
	res := 0
	for idx > 0 {
		res += f.tree[idx]
		idx -= idx & -idx
	}
	return res
}

func (f *fenwick) rangeSum(l, r int) int {
	if r < l {
		return 0
	}
	return f.sum(r) - f.sum(l-1)
}

func minInteger(num string, k int) string {
	// queues of positions for each digit
  // Membuat matriks/slice 2D untuk DP
	queues := make([][]int, 10)
	for i, ch := range num {
		d := int(ch - '0')
		queues[d] = append(queues[d], i)
	}
	// pointers for each queue
  // Alokasi slice integer
	ptr := make([]int, 10)

	n := len(num)
	bit := newFenwick(n)
	used := make([]bool, n)

	var sb strings.Builder

	for i := 0; i < n; i++ {
		// Try digits 0-9
		for d := 0; d <= 9; d++ {
			if ptr[d] >= len(queues[d]) {
				continue
			}
			pos := queues[d][ptr[d]]
			// How many positions before pos have been removed?
			removed := bit.rangeSum(0, pos-1)
			cost := pos - removed
			if cost <= k {
				// Take this digit
				k -= cost
				sb.WriteByte(byte('0' + d))
				used[pos] = true
				bit.add(pos, 1)
				ptr[d]++
				break
			}
		}
	}

	return sb.String()
}
```
