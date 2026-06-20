# 2211 — Count Collisions On A Road

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func countCollisions(directions string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2211: Count Collisions on a Road
// https://leetcode.com/problems/count-collisions-on-a-road/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func countCollisions(directions string) int {
	n := len(directions)
	left, right := 0, n-1

	for left < n && directions[left] == 'L' {
		left++
	}
	for right >= 0 && directions[right] == 'R' {
		right--
	}

	count := 0
	for i := left; i <= right; i++ {
		if directions[i] != 'S' {
			count++
		}
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println(countCollisions("RLRSLL"))
	// Expected: 5

	// Test case 2
	fmt.Println(countCollisions("LLRR"))
	// Expected: 0

	// Test case 3
	fmt.Println(countCollisions("SSRSSRLLRSLLRSRSSRLRRRRRRS"))
	// Expected: 20
}
```
