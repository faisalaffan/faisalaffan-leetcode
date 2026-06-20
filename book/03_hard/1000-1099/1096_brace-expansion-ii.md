# 1096 — Brace Expansion Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func braceExpansionII(expression string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Union-Find (DSU)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Union-Find (DSU)** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1096: Brace Expansion II
// https://leetcode.com/problems/brace-expansion-ii/
// Difficulty: Hard
//
// Recursive descent parser for grammar:
//   expr   → term (',' term)*
//   term   → factor (factor)*   (concatenation = product)
//   factor → '{' expr '}' | letter+
// Returns sorted unique strings.

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(braceExpansionII("{a,b}{c,{d,e}}"))
}

func braceExpansionII(expression string) []string {
	type parser func() []string

	idx := 0

	var parseExpr, parseTerm, parseFactor parser

	parseExpr = func() []string {
		result := parseTerm()
		for idx < len(expression) && expression[idx] == ',' {
			idx++
			result = union(result, parseTerm())
		}
		return result
	}

	parseTerm = func() []string {
		result := []string{""}
		for idx < len(expression) && (expression[idx] == '{' || isLetter(expression[idx])) {
			result = product(result, parseFactor())
		}
		return result
	}

	parseFactor = func() []string {
		if expression[idx] == '{' {
			idx++ // skip '{'
			result := parseExpr()
			idx++ // skip '}'
			return result
		}
		// read consecutive letters
		j := idx
		for j < len(expression) && isLetter(expression[j]) {
			j++
		}
		result := []string{expression[idx:j]}
		idx = j
		return result
	}

	result := parseExpr()
	sort.Strings(result)
	return result
}

func isLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z'
}

func union(a, b []string) []string {
  // Membuat map (HashMap) — pencarian O(1)
	set := make(map[string]bool)
	for _, s := range a {
		set[s] = true
	}
	for _, s := range b {
		set[s] = true
	}
	res := make([]string, 0, len(set))
	for s := range set {
		res = append(res, s)
	}
	return res
}

func product(a, b []string) []string {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}
  // Membuat map (HashMap) — pencarian O(1)
	set := make(map[string]bool)
	for _, s1 := range a {
		for _, s2 := range b {
			set[s1+s2] = true
		}
	}
	res := make([]string, 0, len(set))
	for s := range set {
		res = append(res, s)
	}
	return res
}
```
