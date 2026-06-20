# 1796 — Second Largest Digit In A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func SecondHighest(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1796: Second Largest Digit in a String
// https://leetcode.com/problems/second-largest-digit-in-a-string/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func SecondHighest(s string) int {
	largest := -1
	second := -1
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			digit := int(s[i] - '0')
			if digit > largest {
				second = largest
				largest = digit
			} else if digit < largest && digit > second {
				second = digit
			}
		}
	}
	return second
}

func main() {
	fmt.Println(SecondHighest("dfa12321afd"))
	fmt.Println(SecondHighest("abc1111"))
	fmt.Println(SecondHighest("ck077"))
}
```
