# 1209 — Remove All Adjacent Duplicates In String Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func removeDuplicates(s string, k int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1209: Remove All Adjacent Duplicates in String II
// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string-ii/
// Difficulty: Medium

// Use stack of (char, count). When count reaches k, pop.

// Time: O(n)
// Space: O(n)

func removeDuplicates(s string, k int) string {
	type pair struct {
		char  byte
		count int
	}
	stack := make([]pair, 0)

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1].char == s[i] {
			stack[len(stack)-1].count++
			if stack[len(stack)-1].count == k {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, pair{s[i], 1})
		}
	}

	result := make([]byte, 0)
	for _, p := range stack {
		for j := 0; j < p.count; j++ {
			result = append(result, p.char)
		}
	}
	return string(result)
}

func main() {
	fmt.Printf("%q (expected: %q)\n", removeDuplicates("abcd", 2), "abcd")
	fmt.Printf("%q (expected: %q)\n", removeDuplicates("deeedbbcccbdaa", 3), "aa")
	fmt.Printf("%q (expected: %q)\n", removeDuplicates("pbbcggttciiippooaais", 2), "ps")
}
```
