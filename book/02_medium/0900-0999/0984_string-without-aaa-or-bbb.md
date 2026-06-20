# 0984 — String Without Aaa Or Bbb

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func strWithout3a3b(a int, b int) string
```

> **💡 Hint:** Greedy - always append the character with more remaining count,

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(a + b)  
**Kompleksitas Ruang:** O(a + b) for output

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #984: String Without AAA or BBB
// https://leetcode.com/problems/string-without-aaa-or-bbb/
// Difficulty: Medium
//
// Approach: Greedy - always append the character with more remaining count,
//           but skip a triple by using the other character if we just wrote two.
// Time: O(a + b)
// Space: O(a + b) for output

import "fmt"

func main() {
	fmt.Println(strWithout3a3b(1, 2)) // "bba" or "bab"
	fmt.Println(strWithout3a3b(4, 1)) // "aabaa"
	fmt.Println(strWithout3a3b(3, 3)) // "ababa" or similar
}

func strWithout3a3b(a int, b int) string {
	result := make([]byte, 0, a+b)

	for a > 0 || b > 0 {
		writeA := false
		if a > b {
			writeA = true
		} else if a < b {
			writeA = false
		} else {
			// equal counts: prefer the one that won't create triple
			if len(result) >= 2 && result[len(result)-1] == 'a' && result[len(result)-2] == 'a' {
				writeA = false
			} else if len(result) >= 2 && result[len(result)-1] == 'b' && result[len(result)-2] == 'b' {
				writeA = true
			} else {
				writeA = true
			}
		}

		if writeA {
			result = append(result, 'a')
			a--
			if a > 0 && a >= b {
				result = append(result, 'a')
				a--
			}
		} else {
			result = append(result, 'b')
			b--
			if b > 0 && b >= a {
				result = append(result, 'b')
				b--
			}
		}
	}

	return string(result)
}
```
