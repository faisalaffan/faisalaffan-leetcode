# 2062 — Count Vowel Substrings Of A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountVowelSubstringsOfAString(word string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2062: Count Vowel Substrings of a String
// https://leetcode.com/problems/count-vowel-substrings-of-a-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountVowelSubstringsOfAString("aeiouu"))    // 2
	fmt.Println(CountVowelSubstringsOfAString("unicornarihan")) // 0
	fmt.Println(CountVowelSubstringsOfAString("cuaieuouac"))    // 7
}

// Time: O(n^2), Space: O(1)
func CountVowelSubstringsOfAString(word string) int {
	isVowel := func(ch byte) bool {
		return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
	}

	count := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(word); i++ {
  // Membuat map (HashMap) — pencarian O(1)
		vowelSet := make(map[byte]bool)
		for j := i; j < len(word); j++ {
			if !isVowel(word[j]) {
				break
			}
			vowelSet[word[j]] = true
			if len(vowelSet) == 5 {
				count++
			}
		}
	}
	return count
}
```
