# 3716 — Find Churn Risk Customers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func findChurnRiskCustomers(events []subEvent) []churnResult`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3716: Find Churn Risk Customers
// https://leetcode.com/problems/find-churn-risk-customers/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type subEvent struct {
	userID        int
	eventDate     int
	eventType     string // "start", "cancel", "downgrade"
	planName      string
	monthlyAmount float64
}

type churnResult struct {
	userID               int
	currentPlan          string
	currentMonthlyAmount float64
	maxHistoricalAmount  float64
	daysAsSubscriber     int
}

func findChurnRiskCustomers(events []subEvent) []churnResult {
	// Group events by user
	type userData struct {
		events         []subEvent
		startDate      int
		lastDate       int
		maxAmount      float64
		downgradeCount int
	}

  // HashMap: O(1) lookup
	userMap := make(map[int]*userData)
	for _, e := range events {
		if _, ok := userMap[e.userID]; !ok {
			userMap[e.userID] = &userData{}
		}
		u := userMap[e.userID]
		u.events = append(u.events, e)
		if e.eventType == "start" && (u.startDate == 0 || e.eventDate < u.startDate) {
			u.startDate = e.eventDate
		}
		if e.eventDate > u.lastDate {
			u.lastDate = e.eventDate
		}
		if e.monthlyAmount > u.maxAmount {
			u.maxAmount = e.monthlyAmount
		}
		if e.eventType == "downgrade" {
			u.downgradeCount++
		}
	}

	var results []churnResult
	for uid, u := range userMap {
		// Sort events by date
  // Custom sort
		sort.Slice(u.events, func(i, j int) bool {
			return u.events[i].eventDate < u.events[j].eventDate
		})

		lastEvent := u.events[len(u.events)-1]
		if lastEvent.eventType == "cancel" {
			continue
		}
		if u.downgradeCount == 0 {
			continue
		}
		if lastEvent.monthlyAmount*2 >= u.maxAmount {
			continue
		}
		days := u.lastDate - u.startDate
		if days < 60 {
			continue
		}

		results = append(results, churnResult{
			userID:               uid,
			currentPlan:          lastEvent.planName,
			currentMonthlyAmount: lastEvent.monthlyAmount,
			maxHistoricalAmount:  u.maxAmount,
			daysAsSubscriber:     days,
		})
	}

  // Custom sort
	sort.Slice(results, func(i, j int) bool {
		if results[i].daysAsSubscriber != results[j].daysAsSubscriber {
			return results[i].daysAsSubscriber > results[j].daysAsSubscriber
		}
		return results[i].userID < results[j].userID
	})

	return results
}

func main() {
	events := []subEvent{
		{501, 1, "start", "Basic", 9.99},
		{501, 30, "downgrade", "Lite", 4.99},
		{501, 60, "downgrade", "Free", 0},
		{502, 10, "start", "Pro", 29.99},
		{502, 40, "downgrade", "Basic", 9.99},
		{503, 5, "start", "Pro", 29.99},
		{504, 1, "start", "Basic", 9.99},
		{504, 30, "cancel", "Basic", 0},
		{506, 1, "start", "Basic", 9.99},
	}
	results := findChurnRiskCustomers(events)
	for _, r := range results {
		fmt.Printf("User %d: plan=%s current=%.2f max=%.2f days=%d\n",
			r.userID, r.currentPlan, r.currentMonthlyAmount, r.maxHistoricalAmount, r.daysAsSubscriber)
	}
	if len(results) == 0 {
		fmt.Println("[]")
	}
}
```
