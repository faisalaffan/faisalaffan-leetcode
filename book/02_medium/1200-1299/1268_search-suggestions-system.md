# 1268 — Search Suggestions System

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func suggestedProducts(products []string, searchWord string) [][]string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n log n + m * n) where m = len(searchWord)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1268: Search Suggestions System
// https://leetcode.com/problems/search-suggestions-system/
// Difficulty: Medium

// For each prefix of searchWord, return top 3 lexicographically
// smallest products that match the prefix.

// Time: O(n log n + m * n) where m = len(searchWord)
// Space: O(n)

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)

  // Membuat matriks/slice 2D untuk DP
	result := make([][]string, len(searchWord))

	for i := 1; i <= len(searchWord); i++ {
		prefix := searchWord[:i]
		suggestions := make([]string, 0)

		for _, p := range products {
			if len(p) >= i && p[:i] == prefix {
				suggestions = append(suggestions, p)
				if len(suggestions) == 3 {
					break
				}
			}
		}

		result[i-1] = suggestions
	}

	return result
}

func main() {
	products := []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}
	searchWord := "mouse"
	result := suggestedProducts(products, searchWord)
	for i, r := range result {
		fmt.Printf("prefix %q: %v\n", searchWord[:i+1], r)
	}
}
```
