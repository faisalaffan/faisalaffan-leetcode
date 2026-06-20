# 3412 — Find Mirror Score Of A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func calculateScore(s string) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(n) Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3412: Find Mirror Score of a String
// https://leetcode.com/problems/find-mirror-score-of-a-string/
// Difficulty: Medium
// Time: O(n) Space: O(n)

import "fmt"

func calculateScore(s string) int64 {
  // Membuat matriks/slice 2D untuk DP
	stacks := make([][]int, 26)
	var ans int64
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		ch := int(s[i] - 'a')
		mirror := 25 - ch
		if len(stacks[mirror]) > 0 {
			j := stacks[mirror][len(stacks[mirror])-1]
			stacks[mirror] = stacks[mirror][:len(stacks[mirror])-1]
			ans += int64(i - j)
		} else {
			stacks[ch] = append(stacks[ch], i)
		}
	}
	return ans
}

func main() {
	fmt.Println(calculateScore("aczzx")) // 5
	fmt.Println(calculateScore("abcdef")) // 0
}
```
