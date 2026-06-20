# 0049 — Group Anagrams

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func groupAnagrams(strs []string) [][]string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * k)  
**Kompleksitas Ruang:** O(n * k)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #49: Group Anagrams
// https://leetcode.com/problems/group-anagrams/
// Difficulty: Medium

import "fmt"

func groupAnagrams(strs []string) [][]string {
  // Membuat map (HashMap) — pencarian O(1)
	groups := make(map[[26]byte][]string)

	for _, s := range strs {
		var key [26]byte
  // Loop linear O(n): iterasi setiap elemen
		for i := 0; i < len(s); i++ {
			key[s[i]-'a']++
		}
		groups[key] = append(groups[key], s)
	}

  // Membuat matriks/slice 2D untuk DP
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
