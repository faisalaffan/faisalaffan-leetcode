# 0419 — Battleships In A Board

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func countBattleships(board [][]byte) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m*n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #419: Battleships in a Board
// https://leetcode.com/problems/battleships-in-a-board/
// Difficulty: Medium
// Time: O(m*n) | Space: O(1)

import "fmt"

func countBattleships(board [][]byte) int {
	count := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'X' {
				// Count only if it's the start of a ship (no X above or to the left)
				if (i == 0 || board[i-1][j] != 'X') && (j == 0 || board[i][j-1] != 'X') {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	// Test case 1
	b1 := [][]byte{
		{'X', '.', '.', 'X'},
		{'.', '.', '.', 'X'},
		{'.', '.', '.', 'X'},
	}
	fmt.Println("Test 1:", countBattleships(b1))
	// Expected: 2

	// Test case 2
	b2 := [][]byte{{'X'}}
	fmt.Println("Test 2:", countBattleships(b2))
	// Expected: 1

	// Test case 3: Empty
	b3 := [][]byte{{'.', '.'}, {'.', '.'}}
	fmt.Println("Test 3:", countBattleships(b3))
	// Expected: 0
}
```
