# 1160 — Find Words That Can Be Formed By Characters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func countCharacters(words []string, chars string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * k) where k is max word length  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1160: Find Words That Can Be Formed by Characters
// https://leetcode.com/problems/find-words-that-can-be-formed-by-characters/
// Difficulty: Easy
// Time: O(n * k) where k is max word length | Space: O(1)

import "fmt"

func main() {
	fmt.Println(countCharacters([]string{"cat", "bt", "hat", "tree"}, "atach")) // 6
	fmt.Println(countCharacters([]string{"hello", "world", "leetcode"}, "welldonehoneyr")) // 10
}

// LeetCode submission: countCharacters
func countCharacters(words []string, chars string) int {
  // Alokasi slice integer
	ch := make([]int, 26)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(chars); i++ {
		ch[chars[i]-'a']++
	}
	ans := 0
	for _, w := range words {
  // Alokasi slice integer
		need := make([]int, 26)
		ok := true
  // Loop linear O(n): iterasi setiap elemen
		for i := 0; i < len(w); i++ {
			idx := w[i] - 'a'
			need[idx]++
			if need[idx] > ch[idx] {
				ok = false
				break
			}
		}
		if ok {
			ans += len(w)
		}
	}
	return ans
}
```
