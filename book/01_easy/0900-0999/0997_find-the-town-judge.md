# 0997 — Find The Town Judge

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func findJudge(n int, trust [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(E). Space: O(V).  
**Kompleksitas Ruang:** O(V).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #997: Find the Town Judge
// https://leetcode.com/problems/find-the-town-judge/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(findJudge(2, [][]int{{1, 2}}))             // 2
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}}))    // 3
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}, {3, 1}})) // -1
	fmt.Println(findJudge(3, [][]int{{1, 2}, {2, 3}}))    // -1
}

// findJudge finds the town judge (trusted by everyone, trusts no one).
// Time: O(E). Space: O(V).
func findJudge(n int, trust [][]int) int {
  // Alokasi slice integer
	inDeg := make([]int, n+1)
  // Alokasi slice integer
	outDeg := make([]int, n+1)
	for _, t := range trust {
		outDeg[t[0]]++
		inDeg[t[1]]++
	}
	for i := 1; i <= n; i++ {
		if inDeg[i] == n-1 && outDeg[i] == 0 {
			return i
		}
	}
	return -1
}
```
