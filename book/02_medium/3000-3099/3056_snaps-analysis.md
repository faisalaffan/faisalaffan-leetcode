# 3056 — Snaps Analysis

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func snapsAnalysis(activities []Activity, ages []Age) []AgeBucketAnalysis
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3056: Snaps Analysis
// https://leetcode.com/problems/snaps-analysis/
// Difficulty: Medium (SQL problem — implemented in Go)
// Time: O(n) | Space: O(n)

import (
	"fmt"
)

type Activity struct {
	ActivityID int
	UserID     int
	ActivityType string // 'send' or 'open'
	TimeSpent  float64
}

type Age struct {
	UserID     int
	AgeBucket  string // '21-25', '26-30', '31-35'
}

type AgeBucketAnalysis struct {
	AgeBucket string
	SendPerc  float64
	OpenPerc  float64
}

func snapsAnalysis(activities []Activity, ages []Age) []AgeBucketAnalysis {
	// Build user -> age_bucket map
  // Membuat map (HashMap) — pencarian O(1)
	userAge := make(map[int]string)
	for _, a := range ages {
		userAge[a.UserID] = a.AgeBucket
	}

	// Group activities by age_bucket
	type bucketSums struct {
		totalSend float64
		totalOpen float64
	}
  // Membuat map (HashMap) — pencarian O(1)
	buckets := make(map[string]*bucketSums)

	for _, act := range activities {
		bucket, ok := userAge[act.UserID]
		if !ok {
			continue
		}
		if buckets[bucket] == nil {
			buckets[bucket] = &bucketSums{}
		}
		if act.ActivityType == "send" {
			buckets[bucket].totalSend += act.TimeSpent
		} else if act.ActivityType == "open" {
			buckets[bucket].totalOpen += act.TimeSpent
		}
	}

	var results []AgeBucketAnalysis
	for bucket, sums := range buckets {
		total := sums.totalSend + sums.totalOpen
		sendPerc := 0.0
		openPerc := 0.0
		if total > 0 {
			sendPerc = (sums.totalSend / total) * 100
			openPerc = (sums.totalOpen / total) * 100
		}
		// Round to 2 decimal places
		sendPerc = float64(int(sendPerc*100+0.5)) / 100
		openPerc = float64(int(openPerc*100+0.5)) / 100

		results = append(results, AgeBucketAnalysis{
			AgeBucket: bucket,
			SendPerc:  sendPerc,
			OpenPerc:  openPerc,
		})
	}

	return results
}

func main() {
	activities := []Activity{
		{ActivityID: 1, UserID: 1, ActivityType: "send", TimeSpent: 12.5},
		{ActivityID: 2, UserID: 1, ActivityType: "open", TimeSpent: 7.5},
		{ActivityID: 3, UserID: 2, ActivityType: "send", TimeSpent: 20.0},
		{ActivityID: 4, UserID: 2, ActivityType: "open", TimeSpent: 5.0},
		{ActivityID: 5, UserID: 3, ActivityType: "send", TimeSpent: 8.0},
		{ActivityID: 6, UserID: 3, ActivityType: "open", TimeSpent: 12.0},
		{ActivityID: 7, UserID: 1, ActivityType: "send", TimeSpent: 5.0},
		{ActivityID: 8, UserID: 2, ActivityType: "open", TimeSpent: 10.0},
	}

	ages := []Age{
		{UserID: 1, AgeBucket: "21-25"},
		{UserID: 2, AgeBucket: "21-25"},
		{UserID: 3, AgeBucket: "26-30"},
	}

	fmt.Println("Snaps Analysis")
	fmt.Println("==============")
	fmt.Printf("%-10s %-10s %-10s\n", "AgeBucket", "Send%", "Open%")
	fmt.Println("-----------------------------")

	results := snapsAnalysis(activities, ages)
	for _, r := range results {
		fmt.Printf("%-10s %-10.2f %-10.2f\n", r.AgeBucket, r.SendPerc, r.OpenPerc)
	}
}
```
