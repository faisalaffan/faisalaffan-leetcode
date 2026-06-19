package main

// LeetCode #2115: Find All Possible Recipes from Given Supplies
// https://leetcode.com/problems/find-all-possible-recipes-from-given-supplies/
// Difficulty: Medium
// Time: O(n + m + s) | Space: O(n + m + s)

import "fmt"

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	supplySet := make(map[string]bool)
	for _, s := range supplies {
		supplySet[s] = true
	}

	recipeIdx := make(map[string]int)
	for i, r := range recipes {
		recipeIdx[r] = i
	}

	// indegree for recipes (how many ingredients still needed)
	indegree := make([]int, len(recipes))
	// For each recipe ingredient, which recipes need it
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
