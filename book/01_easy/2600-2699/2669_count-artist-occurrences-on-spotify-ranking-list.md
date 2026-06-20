# 2669 — Count Artist Occurrences On Spotify Ranking List

## Deskripsi

**Soal:** [2669. Count Artist Occurrences On Spotify Ranking List](https://leetcode.com/problems/count-artist-occurrences-on-spotify-ranking-list/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** LIS (Longest Increasing Subsequence)

## Solusi Go

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
