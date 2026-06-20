# 2199 — Finding The Topic Of Each Post

## Deskripsi

**Soal:** [2199. Finding The Topic Of Each Post](https://leetcode.com/problems/finding-the-topic-of-each-post/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func findingTheTopicOfEachPost(posts []Post, keywords []Keyword) map[int][]int`

## Solusi Go

```go
package main

// LeetCode #2199: Finding the Topic of Each Post
// https://leetcode.com/problems/finding-the-topic-of-each-post/
// Difficulty: Hard [Paid]
//
// Given tables: Posts (post_id, content) and Keywords (topic_id, word),
// find for each post the topics that appear in the content (case-insensitive).
// A keyword is considered "appearing" if it appears as a standalone word
// (i.e., not part of another word) in the post content.

import (
	"fmt"
	"strings"
)

// Post represents a post with ID and content.
type Post struct {
	ID      int
	Content string
}

// Keyword maps a topic ID to a keyword string.
type Keyword struct {
	TopicID int
	Word    string
}

// findingTheTopicOfEachPost returns a map from post ID to sorted topic IDs
// whose keywords appear as standalone words (case-insensitive) in the post content.
func findingTheTopicOfEachPost(posts []Post, keywords []Keyword) map[int][]int {
	// build topic keyword index: topicID -> set of lowercase keywords
  // Membuat map untuk pencarian O(1): key → value
	topicWords := make(map[int]map[string]bool)
	for _, kw := range keywords {
		if topicWords[kw.TopicID] == nil {
			topicWords[kw.TopicID] = make(map[string]bool)
		}
		topicWords[kw.TopicID][strings.ToLower(kw.Word)] = true
	}

  // Membuat map untuk pencarian O(1): key → value
	result := make(map[int][]int)

	for _, post := range posts {
		content := strings.ToLower(post.Content)
		// extract words from content (split on non-alphabetic)
		words := strings.FieldsFunc(content, func(r rune) bool {
			return !((r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '\'')
		})

  // Membuat map untuk pencarian O(1): key → value
		wordSet := make(map[string]bool)
		for _, w := range words {
			// strip surrounding punctuation if any
			w = strings.Trim(w, ".,!?;:\"'()[]{}")
			if w != "" {
				wordSet[w] = true
			}
		}

		var matchedTopics []int
		for topicID, kws := range topicWords {
			for kw := range kws {
				if wordSet[kw] {
					matchedTopics = append(matchedTopics, topicID)
					break // one match per topic suffices
				}
			}
		}

		// sort matched topics (simple insertion sort for small slices)
  // Loop standar: indeks 0 sampai n-1
		for i := 0; i < len(matchedTopics); i++ {
			for j := i + 1; j < len(matchedTopics); j++ {
				if matchedTopics[j] < matchedTopics[i] {
					matchedTopics[i], matchedTopics[j] = matchedTopics[j], matchedTopics[i]
				}
			}
		}

		result[post.ID] = matchedTopics
	}

	return result
}

func main() {
	posts := []Post{
		{1, "I love apples and bananas"},
		{2, "The new laptop is great"},
		{3, "Cats are better than dogs"},
	}

	keywords := []Keyword{
		{1, "apple"},
		{1, "banana"},
		{2, "laptop"},
		{3, "cat"},
		{3, "dog"},
	}

	result := findingTheTopicOfEachPost(posts, keywords)
	for _, p := range posts {
		fmt.Printf("Post %d: topics %v\n", p.ID, result[p.ID])
	}
}
```
