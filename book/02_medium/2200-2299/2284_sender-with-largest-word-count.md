# 2284 — Sender With Largest Word Count

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func largestWordCount(messages []string, senders []string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * m)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2284: Sender With Largest Word Count
// https://leetcode.com/problems/sender-with-largest-word-count/
// Difficulty: Medium
// Time: O(n * m) | Space: O(n)

import (
	"fmt"
	"strings"
)

func largestWordCount(messages []string, senders []string) string {
  // Membuat map (HashMap) — pencarian O(1)
	count := make(map[string]int)
	maxCount := 0
	maxSender := ""

	for i, msg := range messages {
		sender := senders[i]
		words := len(strings.Fields(msg))
		count[sender] += words
		if count[sender] > maxCount || (count[sender] == maxCount && sender > maxSender) {
			maxCount = count[sender]
			maxSender = sender
		}
	}
	return maxSender
}

func main() {
	// Test case 1
	fmt.Println(largestWordCount([]string{"Hello userTwooo", "Hi userThree", "Wonderful day Alice", "Nice day userThree"}, []string{"Alice", "userTwo", "userThree", "Alice"}))
	// Expected: "Alice"

	// Test case 2
	fmt.Println(largestWordCount([]string{"t", "e", "s", "t"}, []string{"a", "b", "c", "d"}))
	// Expected: "d" (first with max when equal, by lexicographical)
}
```
