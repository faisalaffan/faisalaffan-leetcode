# 1410 — Html Entity Parser

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func entityParser(text string) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n) where n = length of text  |  **Ruang:** O(n) for the result

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

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
  // Linear scan O(n)
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
