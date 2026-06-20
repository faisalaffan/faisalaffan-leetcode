# 0763 — Partition Labels

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func partitionLabels(s string) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
  // Alokasi slice integer
	last := make([]int, 26)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		last[s[i]-'a'] = i
	}

  // Alokasi slice integer
	result := make([]int, 0)
	start, end := 0, 0

  // Loop linear O(n): iterasi setiap elemen
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
