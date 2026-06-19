package main

// LeetCode #1912: Design Movie Rental System
// https://leetcode.com/problems/design-movie-rental-system/
// Difficulty: Hard

import (
	"container/heap"
	"fmt"
	"sort"
)

// Entry represents a movie copy: (shop, movie, price)
type Entry struct {
	shop, movie, price int
}

// ReportHeap min-heap sorted by (price, shop, movie)
type ReportHeap []Entry

func (h ReportHeap) Len() int { return len(h) }
func (h ReportHeap) Less(i, j int) bool {
	if h[i].price != h[j].price {
		return h[i].price < h[j].price
	}
	if h[i].shop != h[j].shop {
		return h[i].shop < h[j].shop
	}
	return h[i].movie < h[j].movie
}
func (h ReportHeap) Swap(i, j int)     { h[i], h[j] = h[j], h[i] }
func (h *ReportHeap) Push(x any)       { *h = append(*h, x.(Entry)) }
func (h *ReportHeap) Pop() any         { old := *h; n := len(old); x := old[n-1]; *h = old[:n-1]; return x }

// SearchHeap min-heap sorted by (price, shop)
type SearchHeap []Entry

func (h SearchHeap) Len() int { return len(h) }
func (h SearchHeap) Less(i, j int) bool {
	if h[i].price != h[j].price {
		return h[i].price < h[j].price
	}
	return h[i].shop < h[j].shop
}
func (h SearchHeap) Swap(i, j int)     { h[i], h[j] = h[j], h[i] }
func (h *SearchHeap) Push(x any)       { *h = append(*h, x.(Entry)) }
func (h *SearchHeap) Pop() any         { old := *h; n := len(old); x := old[n-1]; *h = old[:n-1]; return x }

type MovieRentalSystem struct {
	avail      map[int]*SearchHeap // per-movie min-heap of available copies
	rented     *ReportHeap         // min-heap of all currently rented copies
	rentedData map[Entry]bool      // tracks which entries are currently rented, keyed by full Entry
	priceOf    map[[2]int]int      // (shop,movie) -> price, for quick lookup in Drop
}

func Constructor(entries [][]int) MovieRentalSystem {
	mrs := MovieRentalSystem{
		avail:      make(map[int]*SearchHeap),
		rented:     &ReportHeap{},
		rentedData: make(map[Entry]bool),
		priceOf:    make(map[[2]int]int),
	}
	byMovie := make(map[int][]Entry)
	for _, e := range entries {
		shop, movie, price := e[0], e[1], e[2]
		byMovie[movie] = append(byMovie[movie], Entry{shop, movie, price})
		mrs.priceOf[[2]int{shop, movie}] = price
	}
	for movie, list := range byMovie {
		sort.Slice(list, func(i, j int) bool {
			if list[i].price != list[j].price {
				return list[i].price < list[j].price
			}
			return list[i].shop < list[j].shop
		})
		h := &SearchHeap{}
		heap.Init(h)
		for _, e := range list {
			heap.Push(h, e)
		}
		mrs.avail[movie] = h
	}
	return mrs
}

func (mrs *MovieRentalSystem) Search(movie int) [][]int {
	h, ok := mrs.avail[movie]
	if !ok {
		return nil
	}
	var result []Entry
	var temp []Entry
	for h.Len() > 0 && len(result) < 5 {
		e := heap.Pop(h).(Entry)
		if !mrs.rentedData[e] {
			result = append(result, e)
			temp = append(temp, e)
		}
	}
	for _, e := range temp {
		heap.Push(h, e)
	}
	res := make([][]int, len(result))
	for i, e := range result {
		res[i] = []int{e.shop, e.movie, e.price}
	}
	return res
}

func (mrs *MovieRentalSystem) Rent(shop int, movie int) {
	price := mrs.priceOf[[2]int{shop, movie}]
	e := Entry{shop, movie, price}
	mrs.rentedData[e] = true
	heap.Push(mrs.rented, e)
}

func (mrs *MovieRentalSystem) Drop(shop int, movie int) {
	price := mrs.priceOf[[2]int{shop, movie}]
	e := Entry{shop, movie, price}
	delete(mrs.rentedData, e)
}

func (mrs *MovieRentalSystem) Report() [][]int {
	var temp []Entry
	var result []Entry
	for mrs.rented.Len() > 0 && len(result) < 5 {
		e := heap.Pop(mrs.rented).(Entry)
		if mrs.rentedData[e] {
			result = append(result, e)
			temp = append(temp, e)
		}
	}
	for _, e := range temp {
		heap.Push(mrs.rented, e)
	}
	if len(result) > 5 {
		result = result[:5]
	}
	res := make([][]int, len(result))
	for i, e := range result {
		res[i] = []int{e.shop, e.movie, e.price}
	}
	return res
}

func DesignMovieRentalSystem() any {
	entries := [][]int{
		{0, 1, 5}, {0, 2, 6}, {0, 3, 7},
		{1, 1, 4}, {1, 2, 7},
		{2, 1, 5},
	}
	mrs := Constructor(entries)
	// Search(1): [[1,1,4],[0,1,5],[2,1,5]]
	_ = mrs.Search(1)
	mrs.Rent(0, 1)
	mrs.Rent(1, 2)
	// Report: [[0,1,5],[1,2,7]]
	_ = mrs.Report()
	mrs.Drop(1, 2)
	// Search(2): [[0,2,6]]
	_ = mrs.Search(2)
	return nil
}

func main() {
	DesignMovieRentalSystem()
	fmt.Println("Movie Rental System test completed")
}
