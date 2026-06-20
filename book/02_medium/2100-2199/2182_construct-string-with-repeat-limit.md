# 2182 — Construct String With Repeat Limit

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func repeatLimitedString(s string, repeatLimit int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2182: Construct String With Repeat Limit
// https://leetcode.com/problems/construct-string-with-repeat-limit/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func repeatLimitedString(s string, repeatLimit int) string {
  // Alokasi slice integer
	count := make([]int, 26)
	for _, ch := range s {
		count[ch-'a']++
	}

	result := make([]byte, 0, len(s))
	for i := 25; i >= 0; {
		if count[i] == 0 {
			i--
			continue
		}

		use := min(count[i], repeatLimit)
		for k := 0; k < use; k++ {
			result = append(result, byte('a'+i))
		}
		count[i] -= use

		if count[i] > 0 {
			j := i - 1
			for j >= 0 && count[j] == 0 {
				j--
			}
			if j < 0 {
				break
			}
			result = append(result, byte('a'+j))
			count[j]--
		} else {
			i--
		}
	}

	return string(result)
}

func main() {
	// Test case 1
	fmt.Println(repeatLimitedString("cczazcc", 3))
	// Expected: "zzcccac"

	// Test case 2
	fmt.Println(repeatLimitedString("aababab", 2))
	// Expected: "bbabaa"

	// Test case 3
	fmt.Println(repeatLimitedString("robnsdvpuxbapuqgopqvxdrchivlifeepy", 2))
	// Expected: "yxxvvuvusrrqqppopponliihgfeeddcba"
}
```
