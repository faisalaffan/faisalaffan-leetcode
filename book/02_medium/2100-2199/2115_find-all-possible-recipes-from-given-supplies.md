# 2115 — Find All Possible Recipes From Given Supplies

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** O(n + m + s)  
**Kompleksitas Ruang:** O(n + m + s)

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2115: Find All Possible Recipes from Given Supplies
// https://leetcode.com/problems/find-all-possible-recipes-from-given-supplies/
// Difficulty: Medium
// Time: O(n + m + s) | Space: O(n + m + s)

import "fmt"

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
  // Membuat map (HashMap) — pencarian O(1)
	supplySet := make(map[string]bool)
	for _, s := range supplies {
		supplySet[s] = true
	}

  // Membuat map (HashMap) — pencarian O(1)
	recipeIdx := make(map[string]int)
	for i, r := range recipes {
		recipeIdx[r] = i
	}

	// indegree for recipes (how many ingredients still needed)
  // Alokasi slice integer
	indegree := make([]int, len(recipes))
	// For each recipe ingredient, which recipes need it
  // Membuat map (HashMap) — pencarian O(1)
	graph := make(map[string][]int)
	for i, ing := range ingredients {
		for _, ig := range ing {
			if !supplySet[ig] {
				indegree[i]++
				graph[ig] = append(graph[ig], i)
			}
		}
	}

	queue := []int{}
	for i, d := range indegree {
		if d == 0 {
			queue = append(queue, i)
		}
	}

	result := []string{}
	for len(queue) > 0 {
		r := queue[0]
		queue = queue[1:]
		result = append(result, recipes[r])

		// This recipe is now a supply for other recipes
		for _, next := range graph[recipes[r]] {
			indegree[next]--
			if indegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findAllRecipes(
		[]string{"bread"},
		[][]string{{"yeast", "flour"}},
		[]string{"yeast", "flour", "corn"},
	))
	// Expected: ["bread"]

	// Test case 2
	fmt.Println("Test 2:", findAllRecipes(
		[]string{"bread", "sandwich"},
		[][]string{{"yeast", "flour"}, {"bread", "meat"}},
		[]string{"yeast", "flour", "meat"},
	))
	// Expected: ["bread", "sandwich"]

	// Test case 3
	fmt.Println("Test 3:", findAllRecipes(
		[]string{"bread", "sandwich", "burger"},
		[][]string{{"yeast", "flour"}, {"bread", "meat"}, {"sandwich", "meat", "bread"}},
		[]string{"yeast", "flour", "meat"},
	))
	// Expected: ["bread", "sandwich", "burger"]
}
```
