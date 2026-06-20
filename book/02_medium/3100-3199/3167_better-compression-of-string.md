# 3167 — Better Compression Of String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func betterCompression(compressed string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3167: Better Compression of String
// https://leetcode.com/problems/better-compression-of-string/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"strconv"
)

func betterCompression(compressed string) string {
  // Alokasi slice integer
	count := make([]int, 26)
	i := 0
	for i < len(compressed) {
		c := compressed[i] - 'a'
		i++
		freq := 0
		for i < len(compressed) && compressed[i] >= '0' && compressed[i] <= '9' {
			freq = freq*10 + int(compressed[i]-'0')
			i++
		}
		count[c] += freq
	}

	ans := make([]byte, 0)
	for c := 0; c < 26; c++ {
		if count[c] > 0 {
			ans = append(ans, byte('a'+c))
			ans = append(ans, []byte(strconv.Itoa(count[c]))...)
		}
	}
	return string(ans)
}

func main() {
	fmt.Println(betterCompression("a12b3c5"))        // Expected: "a12b3c5"
	fmt.Println(betterCompression("a3b2a2c1"))       // Expected: "a5b2c1"
	fmt.Println(betterCompression("z26y25x24"))      // Expected: "x24y25z26"
}
```
