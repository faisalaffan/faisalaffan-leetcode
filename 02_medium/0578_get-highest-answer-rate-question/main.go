package main

// LeetCode #578: Get Highest Answer Rate Question
// https://leetcode.com/problems/get-highest-answer-rate-question/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	// survey_log: {question_id, action}
	// action: "show", "answer", "skip"
	surveyLog := [][]interface{}{
		{1, "show"},
		{1, "answer"},
		{2, "show"},
		{2, "skip"},
		{3, "show"},
		{3, "answer"},
	}
	fmt.Println(MostAnsweredQuestion(surveyLog))
}

func MostAnsweredQuestion(surveyLog [][]interface{}) int {
	shows := make(map[int]int)
	answers := make(map[int]int)

	for _, entry := range surveyLog {
		qID := entry[0].(int)
		action := entry[1].(string)
		shows[qID]++
		if action == "answer" {
			answers[qID]++
		}
	}

	bestQ := -1
	bestRate := -1.0

	for qID := range shows {
		rate := float64(answers[qID]) / float64(shows[qID])
		if rate > bestRate {
			bestRate = rate
			bestQ = qID
		}
	}

	return bestQ
}
