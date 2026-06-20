# 2833 — Furthest Point From Origin

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func FurthestPointFromOrigin(moves string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2833: Furthest Point From Origin
// https://leetcode.com/problems/furthest-point-from-origin/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(FurthestPointFromOrigin("L_RL__R"))
	fmt.Println(FurthestPointFromOrigin("_R__LL_"))
}

func FurthestPointFromOrigin(moves string) int {
	countL, countR, countUnderscore := 0, 0, 0
	for _, c := range moves {
		switch c {
		case 'L':
			countL++
		case 'R':
			countR++
		case '_':
			countUnderscore++
		}
	}
	diff := countL - countR
	if diff < 0 {
		diff = -diff
	}
	return diff + countUnderscore
}
```
