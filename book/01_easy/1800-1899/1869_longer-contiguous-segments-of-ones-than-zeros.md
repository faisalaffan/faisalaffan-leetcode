# 1869 — Longer Contiguous Segments Of Ones Than Zeros

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func CheckZeroOnes(s string) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1869: Longer Contiguous Segments of Ones Than Zeros
// https://leetcode.com/problems/longer-contiguous-segments-of-ones-than-zeros/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func CheckZeroOnes(s string) bool {
	maxOnes, maxZeros := 0, 0
	curOnes, curZeros := 0, 0
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			curOnes++
			curZeros = 0
			if curOnes > maxOnes {
				maxOnes = curOnes
			}
		} else {
			curZeros++
			curOnes = 0
			if curZeros > maxZeros {
				maxZeros = curZeros
			}
		}
	}
	return maxOnes > maxZeros
}

func main() {
	fmt.Println(CheckZeroOnes("1101"))
	fmt.Println(CheckZeroOnes("111000"))
	fmt.Println(CheckZeroOnes("110100010"))
}
```
