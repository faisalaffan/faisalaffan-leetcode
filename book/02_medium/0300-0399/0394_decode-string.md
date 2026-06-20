# 0394 — Decode String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func decodeString(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #394: Decode String
// https://leetcode.com/problems/decode-string/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func decodeString(s string) string {
  // Alokasi slice integer
	numStack := make([]int, 0)
	strStack := make([]string, 0)
	curNum := 0
	curStr := ""

	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			curNum = curNum*10 + int(ch-'0')
		} else if ch == '[' {
			numStack = append(numStack, curNum)
			strStack = append(strStack, curStr)
			curNum = 0
			curStr = ""
		} else if ch == ']' {
			num := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			prevStr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			curStr = prevStr + strings.Repeat(curStr, num)
		} else {
			curStr += string(ch)
		}
	}
	return curStr
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", decodeString("3[a]2[bc]"))
	// Expected: "aaabcbc"

	// Test case 2
	fmt.Println("Test 2:", decodeString("3[a2[c]]"))
	// Expected: "accaccacc"

	// Test case 3
	fmt.Println("Test 3:", decodeString("2[abc]3[cd]ef"))
	// Expected: "abcabccdcdcdef"
}
```
