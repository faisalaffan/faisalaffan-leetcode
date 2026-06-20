# 0067 — Add Binary

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func AddBinary(a string, b string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(max(n,m))  |  **Ruang:** O(max(n,m))


## 💻 Solusi Go

```go
package main

// LeetCode #67: Add Binary
// https://leetcode.com/problems/add-binary/
// Difficulty: Easy

import "fmt"

// Time: O(max(n,m)) | Space: O(max(n,m))
func AddBinary(a string, b string) string {
	i, j, carry := len(a)-1, len(b)-1, 0
	res := make([]byte, 0, max(len(a), len(b))+1)
	for i >= 0 || j >= 0 || carry > 0 {
		if i >= 0 {
			carry += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			carry += int(b[j] - '0')
			j--
		}
		res = append(res, byte('0'+carry%2))
		carry /= 2
	}
	// reverse
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}
	return string(res)
}

func main() {
	fmt.Println(AddBinary("11", "1"))
	fmt.Println(AddBinary("1010", "1011"))
}
```
