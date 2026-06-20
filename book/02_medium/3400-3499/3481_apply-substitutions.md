# 3481 — Apply Substitutions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func ApplySubstitutions(s string, subs map[string]string) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3481: Apply Substitutions
// https://leetcode.com/problems/apply-substitutions/
// Difficulty: Medium [Paid]
// Complexity: O(n * m) time, O(n) space

import (
	"fmt"
	"strings"
)

func main() {
	// Test case 1
	fmt.Println("Test 1:", ApplySubstitutions("Hello %name%", map[string]string{"name": "World"}))
	// Test case 2
	fmt.Println("Test 2:", ApplySubstitutions("%a%%b%", map[string]string{"a": "foo", "b": "bar"}))
	// Test case 3
	fmt.Println("Test 3:", ApplySubstitutions("No placeholders", map[string]string{}))
}

func ApplySubstitutions(s string, subs map[string]string) string {
	result := s
	for key, val := range subs {
		result = strings.ReplaceAll(result, "%"+key+"%", val)
	}
	return result
}
```
