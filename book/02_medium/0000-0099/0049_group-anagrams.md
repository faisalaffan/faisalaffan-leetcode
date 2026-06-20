# 0049 — Group Anagrams

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan dua string. Tugasmu adalah menentukan apakah keduanya **anagram** — mengandung huruf yang sama dengan jumlah sama, hanya urutan berbeda.

**Cara berpikir:** Hitung frekuensi huruf string pertama, kurangi dengan string kedua. Semua harus nol.

**Fungsi Solusi:** `func groupAnagrams(strs []string) [][]string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n * k)  |  **Ruang:** O(n * k)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #49: Group Anagrams
// https://leetcode.com/problems/group-anagrams/
// Difficulty: Medium

import "fmt"

func groupAnagrams(strs []string) [][]string {
  // HashMap: O(1) lookup
	groups := make(map[[26]byte][]string)

	for _, s := range strs {
		var key [26]byte
  // Linear scan O(n)
		for i := 0; i < len(s); i++ {
			key[s[i]-'a']++
		}
		groups[key] = append(groups[key], s)
	}

  // Matriks 2D
	result := make([][]string, 0, len(groups))
	for _, v := range groups {
		result = append(result, v)
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	// [["bat"],["nat","tan"],["ate","eat","tea"]]

	// Test case 2
	fmt.Println(groupAnagrams([]string{""})) // [[""]]

	// Test case 3
	fmt.Println(groupAnagrams([]string{"a"})) // [["a"]]
}

// Time: O(n * k) | Space: O(n * k)
```
