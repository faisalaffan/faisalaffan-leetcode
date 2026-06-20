# 1236 — Web Crawler

## Deskripsi

**Soal:** [1236. Web Crawler](https://leetcode.com/problems/web-crawler/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(V + E) where V = #urls, E = #links  
**Kompleksitas Ruang:** O(V)

**Algoritma:** —

**Fungsi Solusi:** `func crawl(startUrl string, parser HtmlParser) []string`

## Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1236: Web Crawler
// https://leetcode.com/problems/web-crawler/
// Difficulty: Medium [Paid]

// Given a startUrl and an HtmlParser, crawl all reachable URLs
// under the same hostname (same domain).

// Time: O(V + E) where V = #urls, E = #links
// Space: O(V)

type HtmlParser interface {
	GetUrls(url string) []string
}

type mockParser struct {
	urls map[string][]string
}

func (m *mockParser) GetUrls(url string) []string {
	return m.urls[url]
}

func crawl(startUrl string, parser HtmlParser) []string {
	// Extract hostname
	getHost := func(url string) string {
		// Skip protocol
		host := ""
  // Loop standar: indeks 0 sampai n-1
		for i := 0; i < len(url)-7; i++ {
			if url[i:i+7] == "http://" {
				url = url[7:]
				break
			}
		}
		for _, c := range url {
			if c == '/' || c == ':' {
				break
			}
			host += string(c)
		}
		return host
	}

	hostname := getHost(startUrl)
  // Membuat map untuk pencarian O(1): key → value
	visited := make(map[string]bool)
	queue := []string{startUrl}
	visited[startUrl] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, next := range parser.GetUrls(cur) {
			if !visited[next] && getHost(next) == hostname {
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}

  // Membuat slice untuk menyimpan hasil
	result := make([]string, 0, len(visited))
	for url := range visited {
		result = append(result, url)
	}
	return result
}

func main() {
	parser := &mockParser{
		urls: map[string][]string{
			"http://example.com":     {"http://example.com/about", "http://other.com"},
			"http://example.com/about": {"http://example.com/contact"},
			"http://example.com/contact": {},
		},
	}
	result := crawl("http://example.com", parser)
	fmt.Printf("%v\n", result)
}
```
