# 3706 — Maximum Distance Between Unequal Words In Array Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumDistanceBetweenUnequalWordsInArrayIi(words []string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3706: Maximum Distance Between Unequal Words in Array II
// https://leetcode.com/problems/maximum-distance-between-unequal-words-in-array-ii/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func maximumDistanceBetweenUnequalWordsInArrayIi(words []string) int {
	n := len(words)
	ans := 0
	for i := 0; i < n; i++ {
		if words[i] != words[0] {
			if i+1 > ans {
				ans = i + 1
			}
		}
		if words[i] != words[n-1] {
			if n-i > ans {
				ans = n - i
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumDistanceBetweenUnequalWordsInArrayIi([]string{"leetcode", "leetcode", "codeforces"}))
	fmt.Println(maximumDistanceBetweenUnequalWordsInArrayIi([]string{"a", "b", "c", "a", "a"}))
	fmt.Println(maximumDistanceBetweenUnequalWordsInArrayIi([]string{"z", "z", "z"}))
}
```
