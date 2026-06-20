# 0267 — Palindrome Permutation Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func generatePalindromes(s string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Backtracking

**Kompleksitas Waktu:** O((n/2)!), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Backtracking** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #267: Palindrome Permutation II
// https://leetcode.com/problems/palindrome-permutation-ii/
// Difficulty: Medium [Paid]
// Time: O((n/2)!), Space: O(n)

import "fmt"

func generatePalindromes(s string) []string {
	charCount := make([]byte, 128)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		charCount[s[i]]++
	}

	oddChar := byte(0)
	oddCount := 0
  // Range loop: iterasi dengan indeks + nilai
	for i := range charCount {
		if charCount[i]%2 == 1 {
			oddChar = byte(i)
			oddCount++
		}
	}

	if oddCount > 1 {
		return nil
	}

	half := []byte{}
  // Range loop: iterasi dengan indeks + nilai
	for i := range charCount {
		for j := 0; j < int(charCount[i])/2; j++ {
			half = append(half, byte(i))
		}
	}

	result := []string{}
	used := make([]bool, len(half))
	var backtrack func(path []byte)
	backtrack = func(path []byte) {
		if len(path) == len(half) {
			pal := string(path)
			rev := ""
			for i := len(path) - 1; i >= 0; i-- {
				rev += string(path[i])
			}
			if oddChar != 0 {
				pal += string(oddChar)
			}
			result = append(result, pal+rev)
			return
		}

  // Loop linear O(n): iterasi setiap elemen
		for i := 0; i < len(half); i++ {
			if used[i] {
				continue
			}
			if i > 0 && half[i] == half[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			path = append(path, half[i])
			backtrack(path)
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack([]byte{})
	return result
}

func main() {
	fmt.Println(generatePalindromes("aabb"))
	fmt.Println(generatePalindromes("abc"))
	fmt.Println(generatePalindromes("a"))
}
```
