# 1545 — Find Kth Bit In Nth Binary String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func FindKthBit(n int, k int) byte`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(N), Space: O(N)  |  **Ruang:** O(N)


## 💻 Solusi Go

```go
package main

// LeetCode #1545: Find Kth Bit in Nth Binary String
// https://leetcode.com/problems/find-kth-bit-in-nth-binary-string/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(FindKthBit(3, 1))
	fmt.Println(FindKthBit(4, 11))
	fmt.Println(FindKthBit(1, 1))
}

func FindKthBit(n int, k int) byte {
	// Time: O(N), Space: O(N)
	// S1 = "0"
	// Si = Si-1 + "1" + reverse(invert(Si-1))
	// The length of Sn is 2^n - 1
	// We can recursively find the bit without constructing the string

	if n == 1 {
		return '0'
	}

	mid := 1 << (n - 1) // 2^(n-1), the middle position (1-indexed)

	if k == mid {
		return '1'
	} else if k < mid {
		return FindKthBit(n-1, k)
	} else {
		// k > mid: mirror position
		mirrorK := mid*2 - k
		bit := FindKthBit(n-1, mirrorK)
		// invert: '0' <-> '1'
		if bit == '0' {
			return '1'
		}
		return '0'
	}
}
```
