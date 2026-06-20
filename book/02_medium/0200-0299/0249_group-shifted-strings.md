# 0249 — Group Shifted Strings

## Deskripsi

**Soal:** [0249. Group Shifted Strings](https://leetcode.com/problems/group-shifted-strings/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * m), Space: O(n * m)  
**Kompleksitas Ruang:** O(n * m)

**Algoritma:** —

**Fungsi Solusi:** `func groupStrings(strs []string) [][]string`

## Solusi Go

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
  // Membuat map untuk pencarian O(1): key → value
	groups := make(map[string][]string)

	for _, s := range strs {
		key := getKey(s)
		groups[key] = append(groups[key], s)
	}

  // Membuat slice 2D untuk DP/tabel
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

  // Loop standar: indeks 0 sampai n-1
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
