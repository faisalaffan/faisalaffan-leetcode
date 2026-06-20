# 2353 — Design A Food Rating System

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diminta untuk mendesain (merancang) sebuah struktur data kustom dengan operasi tertentu (insert, delete, search, update). Tugasmu adalah memilih representasi data yang tepat agar setiap operasi berjalan efisien — biasanya O(1) atau O(log n).

Ini adalah soal yang paling sering muncul di interview sistem desain. Kamu perlu memilih kombinasi struktur data yang tepat (HashMap + Heap + LinkedList) untuk mencapai kompleksitas yang diminta.

**Konsep kunci:** HashMap (O(1) lookup), Heap (priority), Doubly Linked List (O(1) remove), TreeMap (ordered keys).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor3(foods []string, cuisines []string, ratings []int) FoodRatings
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Heap / Priority Queue, Stack

**Kompleksitas Waktu:** O(log n) per operation  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Heap / Priority Queue** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2353: Design a Food Rating System
// https://leetcode.com/problems/design-a-food-rating-system/
// Difficulty: Medium
// Time: O(log n) per operation | Space: O(n)

import (
	"container/heap"
	"fmt"
)

type FoodRatings struct {
	foodToCuisine map[string]string
	foodToRating  map[string]int
	cuisineToHeap map[string]*foodHeap
}

type foodItem struct {
	name   string
	rating int
}

type foodHeap []foodItem

func (h foodHeap) Len() int { return len(h) }
func (h foodHeap) Less(i, j int) bool {
	if h[i].rating != h[j].rating {
		return h[i].rating > h[j].rating
	}
	return h[i].name < h[j].name
}
func (h foodHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *foodHeap) Push(x any)   { *h = append(*h, x.(foodItem)) }
func (h *foodHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func Constructor3(foods []string, cuisines []string, ratings []int) FoodRatings {
	fr := FoodRatings{
		foodToCuisine: make(map[string]string),
		foodToRating:  make(map[string]int),
		cuisineToHeap: make(map[string]*foodHeap),
	}
	for i, f := range foods {
		fr.foodToCuisine[f] = cuisines[i]
		fr.foodToRating[f] = ratings[i]
		if fr.cuisineToHeap[cuisines[i]] == nil {
			fr.cuisineToHeap[cuisines[i]] = &foodHeap{}
		}
  // Masukkan elemen ke priority queue
		heap.Push(fr.cuisineToHeap[cuisines[i]], foodItem{f, ratings[i]})
	}
	return fr
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	this.foodToRating[food] = newRating
	cuisine := this.foodToCuisine[food]
  // Masukkan elemen ke priority queue
	heap.Push(this.cuisineToHeap[cuisine], foodItem{food, newRating})
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	for this.cuisineToHeap[cuisine].Len() > 0 {
		top := (*this.cuisineToHeap[cuisine])[0]
		if this.foodToRating[top.name] == top.rating {
			return top.name
		}
  // Ambil elemen terkecil/terbesar dari heap
		heap.Pop(this.cuisineToHeap[cuisine])
	}
	return ""
}

func main() {
	fr := Constructor3(
		[]string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"},
		[]string{"korean", "japanese", "japanese", "greek", "japanese", "korean"},
		[]int{9, 12, 8, 15, 14, 7},
	)
	fmt.Println(fr.HighestRated("korean"))   // "bulgogi" (wait, kimchi=9, bulgogi=7... kimchi is higher)
	fmt.Println(fr.HighestRated("japanese"))  // "ramen"
	fr.ChangeRating("sushi", 16)
	fmt.Println(fr.HighestRated("japanese"))  // "sushi"
	fr.ChangeRating("ramen", 16)
	fmt.Println(fr.HighestRated("japanese"))  // "ramen" (tie, lexicographically smaller)
}
```
