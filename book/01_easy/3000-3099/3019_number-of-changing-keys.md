# 3019 — Number Of Changing Keys

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func NumberOfChangingKeys(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3019: Number of Changing Keys
// https://leetcode.com/problems/number-of-changing-keys/
// Difficulty: Easy

import "fmt"
import "unicode"

func main() {
	// LeetCode name: countKeyChanges
	fmt.Println(NumberOfChangingKeys("aAbBcC")) // 2
	fmt.Println(NumberOfChangingKeys("AaAaAaaA")) // 0
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: countKeyChanges
func NumberOfChangingKeys(s string) int {
	count := 0
	for i := 1; i < len(s); i++ {
		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[i-1])) {
			count++
		}
	}
	return count
}
```
