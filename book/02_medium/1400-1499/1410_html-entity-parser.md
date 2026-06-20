# 1410 — Html Entity Parser

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func entityParser(text string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n) where n = length of text  
**Kompleksitas Ruang:** O(n) for the result

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1410: HTML Entity Parser
// https://leetcode.com/problems/html-entity-parser/
// Difficulty: Medium

import "fmt"
import "strings"

func main() {
	// Test case 1
	fmt.Println(entityParser("&amp; is an HTML entity but &ambassador; is not."))
	// "& is an HTML entity but &ambassador; is not."

	// Test case 2
	fmt.Println(entityParser("and I quote: &quot;...&quot;"))
	// "and I quote: \"...\""

	// Test case 3
	fmt.Println(entityParser("x &gt; y &amp;&amp; x &lt; y"))
	// "x > y && x < y"

	// Test case 4
	fmt.Println(entityParser("leetcode.com&frasl;problemset&frasl;all"))
	// "leetcode.com/problemset/all"
}

// Time: O(n) where n = length of text
// Space: O(n) for the result
func entityParser(text string) string {
	entities := map[string]string{
		"&quot;":  "\"",
		"&apos;":  "'",
		"&amp;":   "&",
		"&gt;":    ">",
		"&lt;":    "<",
		"&frasl;": "/",
	}

	var sb strings.Builder
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(text); i++ {
		if text[i] == '&' {
			matched := false
			for entity, replacement := range entities {
				if i+len(entity) <= len(text) && text[i:i+len(entity)] == entity {
					sb.WriteString(replacement)
					i += len(entity) - 1
					matched = true
					break
				}
			}
			if !matched {
				sb.WriteByte('&')
			}
		} else {
			sb.WriteByte(text[i])
		}
	}

	return sb.String()
}
```
