# 0920 — Number Of Music Playlists

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func numMusicPlaylists(n int, goal int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #920: Number of Music Playlists
// https://leetcode.com/problems/number-of-music-playlists/
// Difficulty: Hard
// DP[i][j] = number of playlists of length j using exactly i distinct songs.
// DP[i][j] = DP[i-1][j-1] * (n-(i-1)) + DP[i][j-1] * max(0, i-k)
// Then DP[n][goal] is answer.

import "fmt"

func numMusicPlaylists(n int, goal int, k int) int {
	mod := int(1e9 + 7)
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, n+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int, goal+1)
	}
	dp[0][0] = 1

	for i := 1; i <= n; i++ {
		for j := i; j <= goal; j++ {
			// use a new song
			dp[i][j] = (dp[i][j] + dp[i-1][j-1]*(n-(i-1))) % mod
			// reuse an old song (need k different songs before repeating)
			if i > k {
				dp[i][j] = (dp[i][j] + dp[i][j-1]*(i-k)) % mod
			}
		}
	}
	return dp[n][goal]
}

func main() {
	fmt.Println(numMusicPlaylists(3, 3, 1)) // Expected: 6
	fmt.Println(numMusicPlaylists(2, 3, 0)) // Expected: 6
	fmt.Println(numMusicPlaylists(2, 3, 1)) // Expected: 2
	fmt.Println(numMusicPlaylists(1, 3, 0)) // Expected: 1
	fmt.Println(numMusicPlaylists(3, 3, 2)) // Expected: 6
}
```
