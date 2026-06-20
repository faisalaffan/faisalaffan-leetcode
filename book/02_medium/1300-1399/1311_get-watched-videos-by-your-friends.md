# 1311 — Get Watched Videos By Your Friends

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** O(V + E + F log F) where V = friends count, F = videos count  
**Kompleksitas Ruang:** O(V + F)

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1311: Get Watched Videos by Your Friends
// https://leetcode.com/problems/get-watched-videos-by-your-friends/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// Test case 1
	fmt.Println(watchedVideosByFriends(
		[][]string{{"A", "B"}, {"C"}, {"B", "C"}, {"D"}},
		[][]int{{1, 2}, {0, 3}, {0, 3}, {1, 2}},
		0, 1,
	))
	// ["B","C"]

	// Test case 2
	fmt.Println(watchedVideosByFriends(
		[][]string{{"A", "B"}, {"C"}, {"B", "C"}, {"D"}},
		[][]int{{1, 2}, {0, 3}, {0, 3}, {1, 2}},
		0, 2,
	))
	// ["D"]

	// Test case 3
	fmt.Println(watchedVideosByFriends(
		[][]string{{"a"}, {"a"}, {"a"}},
		[][]int{{1}, {0}, {0}},
		0, 1,
	))
	// ["a"]
}

// Time: O(V + E + F log F) where V = friends count, F = videos count
// Space: O(V + F)
func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	n := len(friends)
	visited := make([]bool, n)
	queue := []int{id}
	visited[id] = true
	depth := 0

	for len(queue) > 0 && depth < level {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]
			for _, f := range friends[curr] {
				if !visited[f] {
					visited[f] = true
					queue = append(queue, f)
				}
			}
		}
		depth++
	}

	if depth < level {
		return []string{}
	}

  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[string]int)
	for _, person := range queue {
		for _, video := range watchedVideos[person] {
			freq[video]++
		}
	}

	videos := make([]string, 0, len(freq))
	for v := range freq {
		videos = append(videos, v)
	}

  // Custom sort dengan comparator
	sort.Slice(videos, func(i, j int) bool {
		if freq[videos[i]] != freq[videos[j]] {
			return freq[videos[i]] < freq[videos[j]]
		}
		return videos[i] < videos[j]
	})

	return videos
}
```
