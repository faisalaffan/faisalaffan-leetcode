# 1487 — Making File Names Unique

## Deskripsi

**Soal:** [1487. Making File Names Unique](https://leetcode.com/problems/making-file-names-unique/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N) average, Space: O(N)  
**Kompleksitas Ruang:** O(N)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1487: Making File Names Unique
// https://leetcode.com/problems/making-file-names-unique/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(GetFolderNames([]string{"pes", "fifa", "gta", "pes(2019)"}))
	fmt.Println(GetFolderNames([]string{"gta", "gta(1)", "gta", "avalon"}))
	fmt.Println(GetFolderNames([]string{"onepiece", "onepiece(1)", "onepiece(2)", "onepiece", "onepiece(1)"}))
}

func GetFolderNames(names []string) []string {
	// Time: O(N) average, Space: O(N)
  // Membuat map untuk pencarian O(1): key → value
	used := make(map[string]int)
  // Membuat slice untuk menyimpan hasil
	result := make([]string, len(names))

	for i, name := range names {
		if _, exists := used[name]; !exists {
			used[name] = 1
			result[i] = name
			continue
		}

		k := used[name]
		candidate := name + "(" + itoa(k) + ")"
		for {
			if _, exists := used[candidate]; exists {
				k++
				candidate = name + "(" + itoa(k) + ")"
			} else {
				break
			}
		}
		used[name] = k + 1
		used[candidate] = 1
		result[i] = candidate
	}

	return result
}

// Simple int to string for positive ints
func itoa(n int) string {
  // Edge case: input kosong
	if n == 0 {
		return "0"
	}
	var buf [12]byte
	pos := len(buf)
	for n > 0 {
		pos--
		buf[pos] = byte('0' + n%10)
		n /= 10
	}
	return string(buf[pos:])
}
```
