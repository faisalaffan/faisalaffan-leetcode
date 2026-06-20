# 0771 — Jewels And Stones

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func numJewelsInStones(jewels string, stones string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(j + s). Space: O(j).  |  **Ruang:** O(j).

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #771: Jewels and Stones
// https://leetcode.com/problems/jewels-and-stones/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb")) // 3
	fmt.Println(numJewelsInStones("z", "ZZ"))       // 0
	fmt.Println(numJewelsInStones("", "abc"))       // 0
}

// numJewelsInStones counts how many stones are also jewels.
// Time: O(j + s). Space: O(j).
func numJewelsInStones(jewels string, stones string) int {
  // HashMap: O(1) lookup
	jSet := make(map[rune]bool)
	for _, c := range jewels {
		jSet[c] = true
	}
	count := 0
	for _, c := range stones {
		if jSet[c] {
			count++
		}
	}
	return count
}
```
