# 2308 — Arrange Table By Gender

## Deskripsi

**Soal:** [2308. Arrange Table By Gender](https://leetcode.com/problems/arrange-table-by-gender/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func arrangeTable(genders []string) []string`

## Solusi Go

```go
package main

// LeetCode #2308: Arrange Table by Gender
// https://leetcode.com/problems/arrange-table-by-gender/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func arrangeTable(genders []string) []string {
  // Membuat slice untuk menyimpan hasil
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
