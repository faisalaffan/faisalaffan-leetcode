# 3278 — Find Candidates For Data Scientist Position Ii

## Deskripsi

**Soal:** [3278. Find Candidates For Data Scientist Position Ii](https://leetcode.com/problems/find-candidates-for-data-scientist-position-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(c * p) Space: O(c * p)  
**Kompleksitas Ruang:** O(c * p)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3278: Find Candidates for Data Scientist Position II
// https://leetcode.com/problems/find-candidates-for-data-scientist-position-ii/
// Difficulty: Medium
// Time: O(c * p) Space: O(c * p)

import (
	"fmt"
	"sort"
)

func main() {
	candidates := []Candidate{
		{101, "Python", 5}, {101, "Tableau", 3}, {101, "PostgreSQL", 4}, {101, "TensorFlow", 2},
		{102, "Python", 4}, {102, "Tableau", 5}, {102, "PostgreSQL", 4}, {102, "R", 4},
		{103, "Python", 3}, {103, "Tableau", 5}, {103, "PostgreSQL", 5}, {103, "Spark", 4},
	}
	projects := []Project{
		{501, "Python", 4}, {501, "Tableau", 3}, {501, "PostgreSQL", 5},
		{502, "Python", 3}, {502, "Tableau", 4}, {502, "R", 2},
	}
	fmt.Println(topCandidates(candidates, projects))
	// Expected: [{501 101 105} {502 102 130}]
}

type Candidate struct {
	ID         int
	Skill      string
	Proficiency int
}

type Project struct {
	ID         int
	Skill      string
	Importance int
}

type ProjectResult struct {
	ProjectID   int
	CandidateID int
	Score       int
}

func topCandidates(candidates []Candidate, projects []Project) []ProjectResult {
	// Build candidate skill map: candidateID -> skill -> proficiency
  // Membuat map untuk pencarian O(1): key → value
	candSkills := make(map[int]map[string]int)
	for _, c := range candidates {
		if candSkills[c.ID] == nil {
			candSkills[c.ID] = make(map[string]int)
		}
		candSkills[c.ID][c.Skill] = c.Proficiency
	}

	// Build project skill map: projectID -> []{skill, importance}
  // Membuat map untuk pencarian O(1): key → value
	projSkills := make(map[int][]struct {
		skill      string
		importance int
	})
  // Membuat map untuk pencarian O(1): key → value
	projReqCount := make(map[int]int)
	for _, p := range projects {
		projSkills[p.ID] = append(projSkills[p.ID], struct {
			skill      string
			importance int
		}{p.Skill, p.Importance})
		projReqCount[p.ID]++
	}

	type result struct {
		projectID   int
		candidateID int
		score       int
	}
	var allResults []result

	for projID, reqs := range projSkills {
		for candID, skills := range candSkills {
			score := 100
			matched := 0
			qualified := true
			for _, req := range reqs {
				prof, ok := skills[req.skill]
				if !ok {
					qualified = false
					break
				}
				if prof > req.importance {
					score += 10
				} else if prof < req.importance {
					score -= 5
				}
				matched++
			}
			if qualified && matched == projReqCount[projID] {
				allResults = append(allResults, result{projID, candID, score})
			}
		}
	}

	// Sort by project, then by score desc, then candidate id asc
	sort.Slice(allResults, func(i, j int) bool {
		if allResults[i].projectID != allResults[j].projectID {
			return allResults[i].projectID < allResults[j].projectID
		}
		if allResults[i].score != allResults[j].score {
			return allResults[i].score > allResults[j].score
		}
		return allResults[i].candidateID < allResults[j].candidateID
	})

	// Pick top candidate per project
	var out []ProjectResult
  // Membuat map untuk pencarian O(1): key → value
	seen := make(map[int]bool)
	for _, r := range allResults {
		if !seen[r.projectID] {
			seen[r.projectID] = true
			out = append(out, ProjectResult{r.projectID, r.candidateID, r.score})
		}
	}
	return out
}
```
