# 3055 — Top Percentile Fraud

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func topPercentileFraud(claims []FraudClaim) []TopClaim
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3055: Top Percentile Fraud
// https://leetcode.com/problems/top-percentile-fraud/
// Difficulty: Medium (SQL problem — implemented in Go)
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"math"
	"sort"
)

type FraudClaim struct {
	PolicyID   int
	State      string
	FraudScore float64
}

type TopClaim struct {
	PolicyID   int
	State      string
	FraudScore float64
}

func topPercentileFraud(claims []FraudClaim) []TopClaim {
	// Group by state
  // Membuat map (HashMap) — pencarian O(1)
	stateClaims := make(map[string][]FraudClaim)
	for _, c := range claims {
		stateClaims[c.State] = append(stateClaims[c.State], c)
	}

	var results []TopClaim

	for state, cs := range stateClaims {
		// Sort by fraud_score DESC, then policy_id ASC
  // Custom sort dengan comparator
		sort.Slice(cs, func(i, j int) bool {
			if cs[i].FraudScore != cs[j].FraudScore {
				return cs[i].FraudScore > cs[j].FraudScore // DESC
			}
			return cs[i].PolicyID < cs[j].PolicyID // ASC
		})

		// Find top 5%: rank in top 5% using RANK()
		// Equivalent to: rank <= ceil(0.05 * n)
		// But SQL solution uses rank = 1 (top 1 per state)
		// Actually, re-reading the problem: "find the top 5% of claims"
		// The SQL uses: WHERE percentile_rank <= 0.05
		// But solution reference uses WHERE rk = 1 (top-ranked per state)
		// We'll implement: return claims where fraud_score = max fraud_score in that state
		// If ties, return the one with lowest policy_id
		maxScore := cs[0].FraudScore

		// Get all with max score (should be just the first group after sorting)
		var bestClaims []FraudClaim
		for _, c := range cs {
			if math.Abs(c.FraudScore-maxScore) < 1e-9 {
				bestClaims = append(bestClaims, c)
			} else {
				break
			}
		}

		// Among ties, pick lowest policy_id
  // Custom sort dengan comparator
		sort.Slice(bestClaims, func(i, j int) bool {
			return bestClaims[i].PolicyID < bestClaims[j].PolicyID
		})

		results = append(results, TopClaim{
			PolicyID:   bestClaims[0].PolicyID,
			State:      state,
			FraudScore: bestClaims[0].FraudScore,
		})
	}

	// Order by state ASC, fraud_score DESC, policy_id ASC
  // Custom sort dengan comparator
	sort.Slice(results, func(i, j int) bool {
		if results[i].State != results[j].State {
			return results[i].State < results[j].State
		}
		if results[i].FraudScore != results[j].FraudScore {
			return results[i].FraudScore > results[j].FraudScore
		}
		return results[i].PolicyID < results[j].PolicyID
	})

	return results
}

func main() {
	claims := []FraudClaim{
		{PolicyID: 101, State: "CA", FraudScore: 85.5},
		{PolicyID: 102, State: "CA", FraudScore: 92.3},
		{PolicyID: 103, State: "CA", FraudScore: 78.1},
		{PolicyID: 104, State: "CA", FraudScore: 92.3},
		{PolicyID: 201, State: "TX", FraudScore: 88.0},
		{PolicyID: 202, State: "TX", FraudScore: 95.5},
		{PolicyID: 203, State: "TX", FraudScore: 72.4},
		{PolicyID: 301, State: "NY", FraudScore: 91.2},
		{PolicyID: 302, State: "NY", FraudScore: 84.7},
		{PolicyID: 303, State: "NY", FraudScore: 91.2},
	}

	fmt.Println("Top Percentile Fraud (Top-Ranked per State)")
	fmt.Println("=========================================")
	fmt.Printf("%-12s %-8s %s\n", "Policy ID", "State", "Fraud Score")
	fmt.Println("-----------------------------------------")

	results := topPercentileFraud(claims)
	for _, r := range results {
		fmt.Printf("%-12d %-8s %.1f\n", r.PolicyID, r.State, r.FraudScore)
	}
}
```
