# 1576 — Replace All S To Avoid Consecutive Repeating Characters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func modifyString(s string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #1576: Replace All ?'s to Avoid Consecutive Repeating Characters
// https://leetcode.com/problems/replace-all-s-to-avoid-consecutive-repeating-characters/
// Difficulty: Easy
//
// LeetCode submission: func modifyString(s string) string

import "fmt"

func main() {
	fmt.Println(ReplaceAllSToAvoidConsecutiveRepeatingCharacters("?zs")) // "azs"
	fmt.Println(ReplaceAllSToAvoidConsecutiveRepeatingCharacters("ubv?w")) // "ubvaw"
	fmt.Println(ReplaceAllSToAvoidConsecutiveRepeatingCharacters("??yw?ipkj?")) // "abywcipkja"
}

// Time: O(n), Space: O(n)
func ReplaceAllSToAvoidConsecutiveRepeatingCharacters(s string) string {
	res := []byte(s)
	for i, ch := range res {
		if ch == '?' {
			for c := byte('a'); c <= 'z'; c++ {
				if (i == 0 || res[i-1] != c) && (i == len(res)-1 || res[i+1] != c) {
					res[i] = c
					break
				}
			}
		}
	}
	return string(res)
}
```
