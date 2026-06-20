# 1990 — Count The Number Of Experiments

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countExperiments(experiments []Experiment) []ExperimentResult
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + 9) = O(n)  
**Kompleksitas Ruang:** O(p * n) where p = platforms, n = names

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1990: Count the Number of Experiments
// https://leetcode.com/problems/count-the-number-of-experiments/
// Difficulty: Medium (SQL problem — simulated in Go)
//
// Simulates: Count experiments for ALL 9 (platform, experiment_name)
// combinations (3 platforms x 3 names). Include zeros for missing
// combinations.

import (
	"fmt"
)

// Experiment represents the Experiments database table.
type Experiment struct {
	ExperimentID   int
	Platform       string // 'Android', 'IOS', 'Web'
	ExperimentName string // 'Reading', 'Sports', 'Programming'
}

// ExperimentResult holds the output.
type ExperimentResult struct {
	Platform       string
	ExperimentName string
	NumExperiments int
}

// countExperiments simulates the SQL query.
// Time: O(n + 9) = O(n) | Space: O(p * n) where p = platforms, n = names
// n = number of experiment rows.
func countExperiments(experiments []Experiment) []ExperimentResult {
	platforms := []string{"Android", "IOS", "Web"}
	names := []string{"Reading", "Sports", "Programming"}

	// Count existing experiments.
  // Membuat map (HashMap) — pencarian O(1)
	counts := make(map[string]map[string]int)
	for _, e := range experiments {
		if counts[e.Platform] == nil {
			counts[e.Platform] = make(map[string]int)
		}
		counts[e.Platform][e.ExperimentName]++
	}

	// Build the 9-combination result.
	var results []ExperimentResult
	for _, p := range platforms {
		for _, n := range names {
			cnt := 0
			if counts[p] != nil {
				cnt = counts[p][n]
			}
			results = append(results, ExperimentResult{
				Platform:       p,
				ExperimentName: n,
				NumExperiments: cnt,
			})
		}
	}

	return results
}

func main() {
	// Test data from the problem.
	experiments := []Experiment{
		{ExperimentID: 1, Platform: "Android", ExperimentName: "Reading"},
		{ExperimentID: 2, Platform: "Android", ExperimentName: "Sports"},
		{ExperimentID: 3, Platform: "Android", ExperimentName: "Programming"},
		{ExperimentID: 4, Platform: "IOS", ExperimentName: "Reading"},
		{ExperimentID: 5, Platform: "IOS", ExperimentName: "Sports"},
		{ExperimentID: 6, Platform: "Web", ExperimentName: "Reading"},
	}

	results := countExperiments(experiments)

	fmt.Println("Experiment Counts (platform | experiment_name | num_experiments):")
	for _, r := range results {
		fmt.Printf("%s | %s | %d\n", r.Platform, r.ExperimentName, r.NumExperiments)
	}
	// Expected output (all 9 combinations):
	// Android | Reading | 1
	// Android | Sports | 1
	// Android | Programming | 1
	// IOS | Reading | 1
	// IOS | Sports | 1
	// IOS | Programming | 0
	// Web | Reading | 1
	// Web | Sports | 0
	// Web | Programming | 0
}
```
