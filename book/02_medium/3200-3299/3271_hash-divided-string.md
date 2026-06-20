# 3271 — Hash Divided String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func stringHash(s string, k int) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n) Space: O(n/k) for result  |  **Ruang:** O(n/k) for result


## 💻 Solusi Go

```go
package main

// LeetCode #3271: Hash Divided String
// https://leetcode.com/problems/hash-divided-string/
// Difficulty: Medium
// Time: O(n) Space: O(n/k) for result

import "fmt"

func main() {
	fmt.Println(stringHash("abcd", 2))                                             // "bf"
	fmt.Println(stringHash("mxz", 3))                                              // "i"
	fmt.Println(stringHash("leetcode", 4))                                         // "ob"
}

func stringHash(s string, k int) string {
	res := make([]byte, 0, len(s)/k)
	sum := 0
	for i, ch := range s {
		sum += int(ch - 'a')
		if (i+1)%k == 0 {
			res = append(res, byte('a'+sum%26))
			sum = 0
		}
	}
	return string(res)
}
```
