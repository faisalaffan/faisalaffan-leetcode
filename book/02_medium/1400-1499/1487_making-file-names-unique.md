# 1487 — Making File Names Unique

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func GetFolderNames(names []string) []string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(N) average, Space: O(N)  |  **Ruang:** O(N)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

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
  // HashMap: O(1) lookup
	used := make(map[string]int)
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
