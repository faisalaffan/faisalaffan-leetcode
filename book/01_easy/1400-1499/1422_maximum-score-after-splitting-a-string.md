# 1422 — Maximum Score After Splitting A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxScore(s string) int

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1422: Maximum Score After Splitting a String
// https://leetcode.com/problems/maximum-score-after-splitting-a-string/
// Difficulty: Easy
//
// LeetCode submission: func maxScore(s string) int

import "fmt"

func main() {
	fmt.Println(MaximumScoreAfterSplittingAString("011101")) // 5
	fmt.Println(MaximumScoreAfterSplittingAString("00111"))  // 5
	fmt.Println(MaximumScoreAfterSplittingAString("1111"))   // 3
}

// Time: O(n), Space: O(1)
func MaximumScoreAfterSplittingAString(s string) int {
	ones := 0
	for _, ch := range s {
		if ch == '1' {
			ones++
		}
	}
	zeros, maxScore := 0, 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			zeros++
		} else {
			ones--
		}
		if zeros+ones > maxScore {
			maxScore = zeros + ones
		}
	}
	return maxScore
}
```
