# 0028 — Find The Index Of The First Occurrence In A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func StrStr(haystack string, needle string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n*m)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #28: Find the Index of the First Occurrence in a String
// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
// Difficulty: Easy

import "fmt"

// Time: O(n*m) | Space: O(1)
func StrStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(StrStr("sadbutsad", "sad"))
	fmt.Println(StrStr("leetcode", "leeto"))
	fmt.Println(StrStr("hello", "ll"))
}
```
