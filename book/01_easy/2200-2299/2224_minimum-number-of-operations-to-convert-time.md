# 2224 — Minimum Number Of Operations To Convert Time

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func MinimumNumberOfOperationsToConvertTime(current string, correct string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2224: Minimum Number of Operations to Convert Time
// https://leetcode.com/problems/minimum-number-of-operations-to-convert-time/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumNumberOfOperationsToConvertTime("02:30", "04:35")) // 3
	fmt.Println(MinimumNumberOfOperationsToConvertTime("11:00", "11:01")) // 1
}

// Time: O(1), Space: O(1)
func MinimumNumberOfOperationsToConvertTime(current string, correct string) int {
	cur := int(current[0]-'0')*600 + int(current[1]-'0')*60 + int(current[3]-'0')*10 + int(current[4]-'0')
	cor := int(correct[0]-'0')*600 + int(correct[1]-'0')*60 + int(correct[3]-'0')*10 + int(correct[4]-'0')

	diff := cor - cur
	ops := 0

	ops += diff / 60
	diff %= 60
	ops += diff / 15
	diff %= 15
	ops += diff / 5
	diff %= 5
	ops += diff

	return ops
}
```
