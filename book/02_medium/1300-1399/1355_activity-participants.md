# 1355 — Activity Participants

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func activityParticipants(activities []struct {
	id   int
	name string
}, friends []struct {
	id         int
	name       string
	activityID int
}) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n) where n = number of friends  
**Kompleksitas Ruang:** O(m) where m = number of activities

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1355: Activity Participants
// https://leetcode.com/problems/activity-participants/
// Difficulty: Medium

import "fmt"

func main() {
	activities := []struct {
		id   int
		name string
	}{
		{1, "Eating"},
		{2, "Singing"},
		{3, "Horse Riding"},
	}
	friends := []struct {
		id         int
		name       string
		activityID int
	}{
		{1, "Jonathan D.", 1},
		{2, "Jade W.", 1},
		{3, "Victor J.", 1},
		{4, "Elvis O.", 2},
		{5, "Daniel A.", 2},
		{6, "Bob B.", 3},
	}

	result := activityParticipants(activities, friends)
	fmt.Println(result) // ["Singing"]
}

// Time: O(n) where n = number of friends
// Space: O(m) where m = number of activities
func activityParticipants(activities []struct {
	id   int
	name string
}, friends []struct {
	id         int
	name       string
	activityID int
}) []string {
  // Membuat map (HashMap) — pencarian O(1)
	counts := make(map[int]int)
	for _, f := range friends {
		counts[f.activityID]++
	}

	if len(counts) == 0 {
		return nil
	}

	// Find min and max counts
	minCount, maxCount := len(friends), 0
	for _, c := range counts {
		if c < minCount {
			minCount = c
		}
		if c > maxCount {
			maxCount = c
		}
	}

	// Find activities with count between min and max (non-inclusive)
	var result []string
	for _, a := range activities {
		c := counts[a.id]
		if c > minCount && c < maxCount {
			result = append(result, a.name)
		}
	}
	return result
}
```
