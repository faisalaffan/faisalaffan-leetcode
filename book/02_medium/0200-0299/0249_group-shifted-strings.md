# 0249 — Group Shifted Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func groupStrings(strs []string) [][]string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n * m), Space: O(n * m)  |  **Ruang:** O(n * m)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #249: Group Shifted Strings
// https://leetcode.com/problems/group-shifted-strings/
// Difficulty: Medium [Paid]
// Time: O(n * m), Space: O(n * m)

import (
	"fmt"
	"strings"
)

func groupStrings(strs []string) [][]string {
  // HashMap: O(1) lookup
	groups := make(map[string][]string)

	for _, s := range strs {
		key := getKey(s)
		groups[key] = append(groups[key], s)
	}

  // Matriks 2D
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}

func getKey(s string) string {
	if len(s) == 0 {
		return ""
	}

	shift := s[0] - 'a'
	var sb strings.Builder

  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		diff := (int(s[i]-'a') - int(shift) + 26) % 26
		sb.WriteByte(byte(diff + 'a'))
	}

	return sb.String()
}

func main() {
	fmt.Println(groupStrings([]string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"}))
	fmt.Println(groupStrings([]string{"a"}))
	fmt.Println(groupStrings([]string{"ab", "bc", "cd"}))
}
```
