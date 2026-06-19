package main

// LeetCode #2991: Top Three Wineries (SQL simulation)
// https://leetcode.com/problems/top-three-wineries/
// Difficulty: Hard [Paid]

import (
	"fmt"
	"sort"
)

type Winery struct {
	Country string
	Winery  string
	Points  int
}

func topThreeWineries(data []Winery) []string {
	type key struct{ country, winery string }
	totals := make(map[key]int)
	for _, w := range data {
		k := key{w.Country, w.Winery}
		totals[k] += w.Points
	}
	byCountry := make(map[string][]struct {
		name   string
		points int
	})
	for k, pts := range totals {
		byCountry[k.country] = append(byCountry[k.country], struct {
			name   string
			points int
		}{k.winery, pts})
	}
	countries := make([]string, 0, len(byCountry))
	for c := range byCountry {
		countries = append(countries, c)
	}
	sort.Strings(countries)
	var result []string
	for _, c := range countries {
		list := byCountry[c]
		sort.Slice(list, func(i, j int) bool {
			if list[i].points != list[j].points {
				return list[i].points > list[j].points
			}
			return list[i].name < list[j].name
		})
		top := list[0]
		row := fmt.Sprintf("%s|%s (%d)", c, top.name, top.points)
		if len(list) >= 2 {
			row += fmt.Sprintf("|%s (%d)", list[1].name, list[1].points)
		} else {
			row += "|No second winery"
		}
		if len(list) >= 3 {
			row += fmt.Sprintf("|%s (%d)", list[2].name, list[2].points)
		} else {
			row += "|No third winery"
		}
		result = append(result, row)
	}
	return result
}

func main() {
	data := []Winery{
		{"USA", "RoyalVines", 47},
		{"USA", "RoyalVines", 39},
		{"USA", "SunsetCellars", 85},
		{"USA", "HarmonyHill", 100},
		{"France", "Bordeaux", 95},
		{"France", "Bordeaux", 3},
		{"France", "Loire", 88},
	}
	for _, r := range topThreeWineries(data) {
		fmt.Println(r)
	}
}
