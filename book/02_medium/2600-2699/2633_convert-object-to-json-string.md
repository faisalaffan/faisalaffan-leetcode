# 2633 — Convert Object To Json String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func convertToJSON(obj any) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2633: Convert Object to JSON String
// https://leetcode.com/problems/convert-object-to-json-string/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import (
	"encoding/json"
	"fmt"
)

func convertToJSON(obj any) string {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return "null"
	}
	return string(bytes)
}

func main() {
	// Test case 1: object
	fmt.Println("Test 1:", convertToJSON(map[string]any{"a": 1, "b": 2}))
	// Expected: {"a":1,"b":2}

	// Test case 2: array
	fmt.Println("Test 2:", convertToJSON([]any{1, "hello", true}))
	// Expected: [1,"hello",true]

	// Test case 3: nested
	fmt.Println("Test 3:", convertToJSON(map[string]any{"x": []any{1, 2, 3}}))
	// Expected: {"x":[1,2,3]}
}
```
