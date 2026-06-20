# 2199 — Finding The Topic Of Each Post

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func findingTheTopicOfEachPost(posts []Post, keywords []Keyword) map[int][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

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
  // Membuat map (HashMap) — pencarian O(1)
	topicWords := make(map[int]map[string]bool)
	for _, kw := range keywords {
		if topicWords[kw.TopicID] == nil {
			topicWords[kw.TopicID] = make(map[string]bool)
		}
		topicWords[kw.TopicID][strings.ToLower(kw.Word)] = true
	}

  // Membuat map (HashMap) — pencarian O(1)
	result := make(map[int][]int)

	for _, post := range posts {
		content := strings.ToLower(post.Content)
		// extract words from content (split on non-alphabetic)
		words := strings.FieldsFunc(content, func(r rune) bool {
			return !((r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '\'')
		})

  // Membuat map (HashMap) — pencarian O(1)
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
  // Loop linear O(n): iterasi setiap elemen
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
