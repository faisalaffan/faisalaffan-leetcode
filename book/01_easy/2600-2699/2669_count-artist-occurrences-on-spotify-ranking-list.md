# 2669 — Count Artist Occurrences On Spotify Ranking List

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CountArtistOccurrencesOnSpotifyRankingList(songs []struct {
	SongID int
	Artist string
}) map[string]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2669: Count Artist Occurrences On Spotify Ranking List
// https://leetcode.com/problems/count-artist-occurrences-on-spotify-ranking-list/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(n)
// Note: SQL/JS problem adapted to Go.

import "fmt"

func main() {
	songs := []struct {
		SongID int
		Artist string
	}{
		{1, "Drake"},
		{2, "Taylor Swift"},
		{3, "Drake"},
		{4, "Ed Sheeran"},
	}
	fmt.Println(CountArtistOccurrencesOnSpotifyRankingList(songs))
}

func CountArtistOccurrencesOnSpotifyRankingList(songs []struct {
	SongID int
	Artist string
}) map[string]int {
	counts := map[string]int{}
	for _, s := range songs {
		counts[s.Artist]++
	}
	return counts
}
```
