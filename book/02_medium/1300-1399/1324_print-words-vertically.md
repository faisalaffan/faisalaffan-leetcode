# 1324 — Print Words Vertically

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func printVertically(s string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m*n) where m = max word length, n = number of words  
**Kompleksitas Ruang:** O(m*n) for the result

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1324: Print Words Vertically
// https://leetcode.com/problems/print-words-vertically/
// Difficulty: Medium

import "fmt"
import "strings"

func main() {
	// Test case 1
	fmt.Println(printVertically("HOW ARE YOU"))
	// ["HAY","ORO","WEU"]

	// Test case 2
	fmt.Println(printVertically("TO BE OR NOT TO BE"))
	// ["TBONTB","OEROOE","   T"]

	// Test case 3
	fmt.Println(printVertically("CONTEST IS COMING"))
	// ["CIC","OSO","N M","T I","E N","S G","T"]
}

// Time: O(m*n) where m = max word length, n = number of words
// Space: O(m*n) for the result
func printVertically(s string) []string {
	words := strings.Fields(s)
	maxLen := 0
	for _, w := range words {
		if len(w) > maxLen {
			maxLen = len(w)
		}
	}

	result := make([]string, maxLen)
	for i := 0; i < maxLen; i++ {
		var sb strings.Builder
		for _, w := range words {
			if i < len(w) {
				sb.WriteByte(w[i])
			} else {
				sb.WriteByte(' ')
			}
		}
		// Trim trailing spaces
		result[i] = strings.TrimRight(sb.String(), " ")
	}

	return result
}
```
