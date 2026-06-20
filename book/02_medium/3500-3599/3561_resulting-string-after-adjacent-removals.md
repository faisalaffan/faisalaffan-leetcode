# 3561 — Resulting String After Adjacent Removals

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func ResultingStringAfterAdjacentRemovals(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3561: Resulting String After Adjacent Removals
// https://leetcode.com/problems/resulting-string-after-adjacent-removals/
// Difficulty: Medium
// Complexity: O(n) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", ResultingStringAfterAdjacentRemovals("abbaca"))
	// Test case 2
	fmt.Println("Test 2:", ResultingStringAfterAdjacentRemovals("azxxzy"))
	// Test case 3
	fmt.Println("Test 3:", ResultingStringAfterAdjacentRemovals("a"))
}

func ResultingStringAfterAdjacentRemovals(s string) string {
	stack := make([]byte, 0, len(s))
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
```
