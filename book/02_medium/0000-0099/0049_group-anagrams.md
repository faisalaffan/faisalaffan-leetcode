# 0049 — Group Anagrams

## Deskripsi

**Soal:** [0049. Group Anagrams](https://leetcode.com/problems/group-anagrams/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * k)  
**Kompleksitas Ruang:** O(n * k)

**Algoritma:** —

**Fungsi Solusi:** `func groupAnagrams(strs []string) [][]string`

## Solusi Go

```go
package main

// LeetCode #49: Group Anagrams
// https://leetcode.com/problems/group-anagrams/
// Difficulty: Medium

import "fmt"

func groupAnagrams(strs []string) [][]string {
  // Membuat map untuk pencarian O(1): key → value
	groups := make(map[[26]byte][]string)

	for _, s := range strs {
		var key [26]byte
  // Loop standar: indeks 0 sampai n-1
		for i := 0; i < len(s); i++ {
			key[s[i]-'a']++
		}
		groups[key] = append(groups[key], s)
	}

  // Membuat slice 2D untuk DP/tabel
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
