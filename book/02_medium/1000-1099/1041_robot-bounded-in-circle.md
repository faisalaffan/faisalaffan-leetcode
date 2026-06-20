# 1041 — Robot Bounded In Circle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func isRobotBounded(instructions string) bool
```

> **💡 Hint:** Simulate robot movement. Robot is bounded in circle iff

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1041: Robot Bounded In Circle
// https://leetcode.com/problems/robot-bounded-in-circle/
// Difficulty: Medium
//
// Approach: Simulate robot movement. Robot is bounded in circle iff
//           final position is (0,0) OR direction != North.
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(isRobotBounded("GGLLGG")) // true
	fmt.Println(isRobotBounded("GG"))     // false
	fmt.Println(isRobotBounded("GL"))     // true
}

func isRobotBounded(instructions string) bool {
	x, y := 0, 0
	dirX, dirY := 0, 1 // facing north

	for _, c := range instructions {
		switch c {
		case 'G':
			x += dirX
			y += dirY
		case 'L':
			dirX, dirY = -dirY, dirX
		case 'R':
			dirX, dirY = dirY, -dirX
		}
	}

	return (x == 0 && y == 0) || !(dirX == 0 && dirY == 1)
}
```
