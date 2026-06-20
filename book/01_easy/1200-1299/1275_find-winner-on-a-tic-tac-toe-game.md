# 1275 — Find Winner On A Tic Tac Toe Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func tictactoe(moves [][]int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1275: Find Winner on a Tic Tac Toe Game
// https://leetcode.com/problems/find-winner-on-a-tic-tac-toe-game/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(tictactoe([][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}})) // "A"
	fmt.Println(tictactoe([][]int{{0, 0}, {1, 1}, {0, 1}, {0, 2}, {1, 0}, {2, 0}})) // "B"
	fmt.Println(tictactoe([][]int{{0, 0}, {1, 1}, {2, 0}, {1, 0}, {1, 2}, {2, 1}, {0, 1}, {0, 2}, {2, 2}})) // "Draw"
}

// LeetCode submission: tictactoe
func tictactoe(moves [][]int) string {
  // Membuat matriks/slice 2D untuk DP
	board := make([][]byte, 3)
  // Range loop: iterasi dengan indeks + nilai
	for i := range board {
		board[i] = make([]byte, 3)
	}
	for i, m := range moves {
		player := byte('A')
		if i%2 == 1 {
			player = 'B'
		}
		board[m[0]][m[1]] = player
	}
	// Check rows and cols
	for i := 0; i < 3; i++ {
		if board[i][0] != 0 && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return string(board[i][0])
		}
		if board[0][i] != 0 && board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return string(board[0][i])
		}
	}
	// Check diagonals
	if board[0][0] != 0 && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return string(board[0][0])
	}
	if board[0][2] != 0 && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return string(board[0][2])
	}
	if len(moves) == 9 {
		return "Draw"
	}
	return "Pending"
}
```
