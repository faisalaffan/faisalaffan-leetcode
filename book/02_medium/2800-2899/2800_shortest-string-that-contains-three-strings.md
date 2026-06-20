# 2800 — Shortest String That Contains Three Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func merge2(x, y string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2800: Shortest String That Contains Three Strings
// https://leetcode.com/problems/shortest-string-that-contains-three-strings/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func merge2(x, y string) string {
	if len(x) < len(y) {
		return merge2(y, x)
	}
	// Check if x contains y
	for i := 0; i <= len(x)-len(y); i++ {
		if x[i:i+len(y)] == y {
			return x
		}
	}
	// Find longest overlap
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(y); i++ {
		overlap := len(y) - i
		if overlap <= len(x) && x[len(x)-overlap:] == y[:overlap] {
			return x + y[i:]
		}
	}
	return x + y
}

func ShortestStringThatContainsThreeStrings(a string, b string, c string) string {
	strs := []string{a, b, c}
	perms := [][]int{
		{0, 1, 2}, {0, 2, 1}, {1, 0, 2},
		{1, 2, 0}, {2, 0, 1}, {2, 1, 0},
	}

	best := a + b + c
	for _, p := range perms {
		merged := merge2(merge2(strs[p[0]], strs[p[1]]), strs[p[2]])
		if len(merged) < len(best) || (len(merged) == len(best) && merged < best) {
			best = merged
		}
	}

	return best
}

func main() {
	fmt.Println(ShortestStringThatContainsThreeStrings("abc", "bcd", "cde"))
	fmt.Println(ShortestStringThatContainsThreeStrings("a", "ab", "abc"))
}
```
