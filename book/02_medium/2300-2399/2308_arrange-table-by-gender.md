# 2308 — Arrange Table By Gender

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func arrangeTable(genders []string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
