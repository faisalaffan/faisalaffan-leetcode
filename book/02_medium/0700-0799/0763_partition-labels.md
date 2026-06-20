# 0763 — Partition Labels

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func partitionLabels(s string) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #763: Partition Labels
// https://leetcode.com/problems/partition-labels/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
	fmt.Println(partitionLabels("eccbbbbdec"))
}

func partitionLabels(s string) []int {
  // Alokasi slice
	last := make([]int, 26)
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		last[s[i]-'a'] = i
	}

  // Alokasi slice
	result := make([]int, 0)
	start, end := 0, 0

  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		if last[s[i]-'a'] > end {
			end = last[s[i]-'a']
		}
		if i == end {
			result = append(result, end-start+1)
			start = i + 1
		}
	}

	return result
}
```
