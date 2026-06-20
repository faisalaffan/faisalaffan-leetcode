# 3846 — Total Distance To Type A String Using One Finger

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func TotalDistanceToTypeAStringUsingOneFinger(s string) int
```

> **💡 Hint:** Precompute keyboard positions, simulate typing from 'a'.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3846: Total Distance to Type a String Using One Finger
// https://leetcode.com/problems/total-distance-to-type-a-string-using-one-finger/
// Difficulty: Medium [Paid]
// Time: O(N) | Space: O(1)
// Approach: Precompute keyboard positions, simulate typing from 'a'.

import "fmt"

func TotalDistanceToTypeAStringUsingOneFinger(s string) int {
	// Keyboard layout (row, col)
	keyboard := []string{
		"qwertyuiop",
		"asdfghjkl",
		"zxcvbnm",
	}

  // Membuat map (HashMap) — pencarian O(1)
	pos := make(map[byte][2]int)
	for r, row := range keyboard {
		for c, ch := range row {
			pos[byte(ch)] = [2]int{r, c}
		}
	}

	total := 0
	cur := pos['a']
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		next := pos[s[i]]
		dist := abs(cur[0]-next[0]) + abs(cur[1]-next[1])
		total += dist
		cur = next
	}

	return total
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Example 1
	fmt.Println(TotalDistanceToTypeAStringUsingOneFinger("hello")) // Expected: 17

	// Example 2
	fmt.Println(TotalDistanceToTypeAStringUsingOneFinger("a")) // Expected: 0

	// Example 3
	fmt.Println(TotalDistanceToTypeAStringUsingOneFinger("qaz")) // q: (0,0), a: (1,0), z: (2,0) = |0-1|+|0-0| + |1-2|+|0-0| = 1+1 = 2
}
```
