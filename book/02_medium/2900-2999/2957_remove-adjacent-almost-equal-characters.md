# 2957 — Remove Adjacent Almost Equal Characters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func removeAlmostEqualCharacters(word string) (ans int)`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2957: Remove Adjacent Almost-Equal Characters
// https://leetcode.com/problems/remove-adjacent-almost-equal-characters/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(removeAlmostEqualCharacters("aaaaa"))
	fmt.Println(removeAlmostEqualCharacters("abdcd"))
	fmt.Println(removeAlmostEqualCharacters("acb"))
}

func removeAlmostEqualCharacters(word string) (ans int) {
	for i := 1; i < len(word); i++ {
		d := int(word[i]) - int(word[i-1])
		if d < 0 {
			d = -d
		}
		if d < 2 {
			ans++
			i++
		}
	}
	return
}
```
