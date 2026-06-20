# 1221 — Split A String In Balanced Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func balancedStringSplit(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1221: Split a String in Balanced Strings
// https://leetcode.com/problems/split-a-string-in-balanced-strings/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(balancedStringSplit("RLRRLLRLRL")) // 4
	fmt.Println(balancedStringSplit("RLLLLRRRLR")) // 3
	fmt.Println(balancedStringSplit("LLLLRRRR"))   // 1
}

// LeetCode submission: balancedStringSplit
func balancedStringSplit(s string) int {
	count, ans := 0, 0
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		if s[i] == 'R' {
			count++
		} else {
			count--
		}
		if count == 0 {
			ans++
		}
	}
	return ans
}
```
