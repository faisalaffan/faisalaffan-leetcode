# 2308 — Arrange Table By Gender

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func arrangeTable(genders []string) []string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #2308: Arrange Table by Gender
// https://leetcode.com/problems/arrange-table-by-gender/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func arrangeTable(genders []string) []string {
	result := make([]string, 0, len(genders))
	// Order: female, other, male
	females, males, others := []string{}, []string{}, []string{}
	for _, g := range genders {
		if g == "female" {
			females = append(females, g)
		} else if g == "male" {
			males = append(males, g)
		} else {
			others = append(others, g)
		}
	}
	result = append(result, females...)
	result = append(result, others...)
	result = append(result, males...)
	return result
}

func main() {
	// Test case 1
	fmt.Println(arrangeTable([]string{"male", "female", "female", "other", "male", "other"}))
	// Expected: ["female", "female", "other", "other", "male", "male"]

	// Test case 2
	fmt.Println(arrangeTable([]string{"female", "male"}))
	// Expected: ["female", "male"]
}
```
