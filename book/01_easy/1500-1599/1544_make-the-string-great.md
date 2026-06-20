# 1544 — Make The String Great

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func makeGood(s string) string

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1544: Make The String Great
// https://leetcode.com/problems/make-the-string-great/
// Difficulty: Easy
//
// LeetCode submission: func makeGood(s string) string

import "fmt"

func main() {
	fmt.Println(MakeTheStringGreat("leEeetcode")) // "leetcode"
	fmt.Println(MakeTheStringGreat("abBAcC"))     // ""
	fmt.Println(MakeTheStringGreat("s"))          // "s"
}

// Time: O(n), Space: O(n)
func MakeTheStringGreat(s string) string {
	stack := make([]byte, 0, len(s))
  // Range loop: iterasi dengan indeks + nilai
	for i := range s {
		stack = append(stack, s[i])
		n := len(stack)
		if n >= 2 {
			diff := int(stack[n-1]) - int(stack[n-2])
			if diff == 32 || diff == -32 {
				stack = stack[:n-2]
			}
		}
	}
	return string(stack)
}
```
