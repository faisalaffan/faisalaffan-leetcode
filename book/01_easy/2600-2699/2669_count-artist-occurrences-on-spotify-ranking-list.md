# 2669 — Count Artist Occurrences On Spotify Ranking List

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func CountArtistOccurrencesOnSpotifyRankingList(songs []struct { SongID int Artist string }) map[string]int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

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
