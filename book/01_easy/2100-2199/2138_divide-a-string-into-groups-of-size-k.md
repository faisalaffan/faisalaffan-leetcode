# 2138 — Divide A String Into Groups Of Size K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func DivideAStringIntoGroupsOfSizeK(s string, k int, fill byte) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2138: Divide a String Into Groups of Size k
// https://leetcode.com/problems/divide-a-string-into-groups-of-size-k/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(DivideAStringIntoGroupsOfSizeK("abcdefghi", 3, 'x')) // ["abc" "def" "ghi"]
	fmt.Println(DivideAStringIntoGroupsOfSizeK("abcdefghij", 3, 'x')) // ["abc" "def" "ghi" "jxx"]
}

// Time: O(n), Space: O(n)
func DivideAStringIntoGroupsOfSizeK(s string, k int, fill byte) []string {
	var result []string
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i += k {
		end := i + k
		if end > len(s) {
			end = len(s)
		}
		group := s[i:end]
		if len(group) < k {
			for len(group) < k {
				group += string(fill)
			}
		}
		result = append(result, group)
	}
	return result
}
```
