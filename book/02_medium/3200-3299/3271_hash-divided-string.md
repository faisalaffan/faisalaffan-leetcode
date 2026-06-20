# 3271 — Hash Divided String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func stringHash(s string, k int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n) Space: O(n/k) for result  
**Kompleksitas Ruang:** O(n/k) for result

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
