# 0439 — Ternary Expression Parser

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func parseTernary(expression string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Stack

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #439: Ternary Expression Parser
// https://leetcode.com/problems/ternary-expression-parser/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func parseTernary(expression string) string {
	stack := make([]byte, 0)
	// Process from right to left
	for i := len(expression) - 1; i >= 0; i-- {
		ch := expression[i]
		if ch >= '0' && ch <= '9' || ch == 'T' || ch == 'F' {
			stack = append(stack, ch)
		} else if ch == '?' {
			// Evaluate: preceding char is condition
			trueVal := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			falseVal := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			cond := expression[i-1]
			i-- // skip the condition character (already consumed)
			if cond == 'T' {
				stack = append(stack, trueVal)
			} else {
				stack = append(stack, falseVal)
			}
		}
		// Skip ':'
	}
	return string(stack[0])
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", parseTernary("T?2:3"))
	// Expected: "2"

	// Test case 2
	fmt.Println("Test 2:", parseTernary("F?1:T?4:5"))
	// Expected: "4"

	// Test case 3
	fmt.Println("Test 3:", parseTernary("T?T?F:5:3"))
	// Expected: "F"
}
```
