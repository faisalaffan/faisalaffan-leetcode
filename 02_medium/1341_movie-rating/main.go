package main

// LeetCode #1341: Movie Rating
// https://leetcode.com/problems/movie-rating/
// Difficulty: Medium

import "fmt"

func main() {
	movies := []struct {
		id    int
		title string
	}{
		{1, "Avengers"},
		{2, "Frozen 2"},
		{3, "Joker"},
	}
	users := []struct {
		id   int
		name string
	}{
		{1, "Daniel"},
		{2, "Monica"},
		{3, "Maria"},
	}
	ratings := []struct {
		userID    int
		movieID   int
		rating    int
		createdAt string
	}{
		{1, 1, 3, "2020-01-12"},
		{1, 2, 4, "2020-02-11"},
		{1, 3, 2, "2020-02-12"},
		{2, 1, 5, "2020-02-17"},
		{2, 2, 2, "2020-02-01"},
		{2, 3, 5, "2020-03-01"},
		{3, 1, 3, "2020-02-22"},
		{3, 2, 4, "2020-02-25"},
	}

	result := movieRating(ratings, users, movies)
	fmt.Println(result)
}

// Time: O(n) for counting
// Space: O(n) for maps
func movieRating(ratings []struct {
	userID    int
	movieID   int
	rating    int
	createdAt string
}, users []struct {
	id   int
	name string
}, movies []struct {
	id    int
	title string
}) string {
	// Count ratings per user
	userRatings := make(map[int]int)
	for _, r := range ratings {
		userRatings[r.userID]++
	}

	maxRatings := 0
	bestUser := ""
	for _, u := range users {
		cnt := userRatings[u.id]
		if cnt > maxRatings || (cnt == maxRatings && (bestUser == "" || u.name < bestUser)) {
			maxRatings = cnt
			bestUser = u.name
		}
	}

	// Average rating in Feb 2020
	movieScores := make(map[int]struct{ sum, count int })
	for _, r := range ratings {
		if r.createdAt >= "2020-02-01" && r.createdAt <= "2020-02-29" {
			s := movieScores[r.movieID]
			s.sum += r.rating
			s.count++
			movieScores[r.movieID] = s
		}
	}

	bestAvg := 0.0
	bestMovie := ""
	for _, m := range movies {
		if s, ok := movieScores[m.id]; ok && s.count > 0 {
			avg := float64(s.sum) / float64(s.count)
			if avg > bestAvg || (avg == bestAvg && (bestMovie == "" || m.title < bestMovie)) {
				bestAvg = avg
				bestMovie = m.title
			}
		}
	}

	return bestUser + " " + bestMovie
}
