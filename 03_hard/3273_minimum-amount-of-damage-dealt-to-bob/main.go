package main

// LeetCode #3273: Minimum Amount of Damage Dealt to Bob
// https://leetcode.com/problems/minimum-amount-of-damage-dealt-to-bob/
// Difficulty: Hard
//
// Bob has `power` attack. Enemies have `damage[i]` and `health[i]`.
// Each second, Bob attacks one enemy (reducing its health by power),
// and every alive enemy deals its damage to Bob.
// Find the minimum total damage Bob takes by choosing the optimal kill order.
//
// This is a scheduling problem. The optimal order is to sort by
//   time_to_kill[i] / damage[i]
// where time_to_kill[i] = ceil(health[i] / power).
// Equivalently, compare using cross-multiplication:
//   t_i * d_j < t_j * d_i  => i before j

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(minDamage(4, []int{1, 2, 3, 4}, []int{4, 5, 6, 8}))
	// Example 2
	fmt.Println(minDamage(1, []int{1, 1, 1, 1}, []int{1, 1, 1, 1}))
	// Example 3
	fmt.Println(minDamage(10, []int{5, 5, 5}, []int{10, 20, 30}))
	// Example 4: single enemy
	fmt.Println(minDamage(3, []int{7}, []int{10}))
	// Example 5: large health
	fmt.Println(minDamage(2, []int{3, 4}, []int{10, 10}))
}

func minDamage(power int, damage []int, health []int) int64 {
	n := len(damage)
	type enemy struct {
		t int64 // time to kill (ceil(health/power))
		d int64 // damage per second
	}
	enemies := make([]enemy, n)
	for i := range enemies {
		enemies[i].t = int64((health[i] + power - 1) / power)
		enemies[i].d = int64(damage[i])
	}

	// Sort by t/d ratio ascending.
	// Equivalent sort comparator: t_i * d_j < t_j * d_i
	sort.Slice(enemies, func(i, j int) bool {
		return enemies[i].t*enemies[j].d < enemies[j].t*enemies[i].d
	})

	var totalDamage int64
	var elapsed int64
	for _, e := range enemies {
		elapsed += e.t
		totalDamage += elapsed * e.d
	}

	return totalDamage
}
