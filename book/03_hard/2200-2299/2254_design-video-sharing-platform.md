# 2254 — Design Video Sharing Platform

## Deskripsi

**Soal:** [2254. Design Video Sharing Platform](https://leetcode.com/problems/design-video-sharing-platform/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func Constructor() *VideoSharingPlatform`

## Solusi Go

```go
package main

// LeetCode #2254: Design Video Sharing Platform
// https://leetcode.com/problems/design-video-sharing-platform/
// Difficulty: Hard [Paid]
//
// Design a video sharing platform with the following operations:
//   - upload(video): returns a unique video ID (auto-increment).
//   - remove(videoId): marks video as removed.
//   - watch(videoId, startMinute, endMinute): records that a user watched
//     the video from startMinute to endMinute (both inclusive, 0-indexed).
//   - like(videoId): increments like count.
//   - dislike(videoId): increments dislike count.
//   - getLikesAndDislikes(videoId): returns [likes, dislikes].
//   - getViews(videoId): returns total view duration across all watches
//     (sum of (endMinute - startMinute + 1) for each watch).

import (
	"fmt"
)

// Video represents a video with its metadata.
type Video struct {
	ID       int
	Likes    int
	Dislikes int
	Removed  bool
	Views    int // total watched minutes
}

// VideoSharingPlatform manages videos.
type VideoSharingPlatform struct {
	videos    map[int]*Video
	nextID    int
}

// Constructor creates a new platform.
func Constructor() *VideoSharingPlatform {
	return &VideoSharingPlatform{
		videos: make(map[int]*Video),
		nextID: 0,
	}
}

// Upload adds a new video and returns its ID.
func (vsp *VideoSharingPlatform) Upload(video string) int {
	id := vsp.nextID
	vsp.nextID++
	vsp.videos[id] = &Video{
		ID:       id,
		Likes:    0,
		Dislikes: 0,
		Removed:  false,
		Views:    0,
	}
	return id
}

// Remove marks a video as removed.
func (vsp *VideoSharingPlatform) Remove(videoID int) {
	if v, ok := vsp.videos[videoID]; ok {
		v.Removed = true
	}
}

// Watch records a watch from startMinute to endMinute.
func (vsp *VideoSharingPlatform) Watch(videoID, startMinute, endMinute int) {
	if v, ok := vsp.videos[videoID]; ok && !v.Removed {
		duration := endMinute - startMinute + 1
		if duration > 0 {
			v.Views += duration
		}
	}
}

// Like increments the like count for a video.
func (vsp *VideoSharingPlatform) Like(videoID int) {
	if v, ok := vsp.videos[videoID]; ok && !v.Removed {
		v.Likes++
	}
}

// Dislike increments the dislike count for a video.
func (vsp *VideoSharingPlatform) Dislike(videoID int) {
	if v, ok := vsp.videos[videoID]; ok && !v.Removed {
		v.Dislikes++
	}
}

// GetLikesAndDislikes returns [likes, dislikes] for a video.
func (vsp *VideoSharingPlatform) GetLikesAndDislikes(videoID int) []int {
	if v, ok := vsp.videos[videoID]; ok && !v.Removed {
		return []int{v.Likes, v.Dislikes}
	}
	return []int{-1, -1}
}

// GetViews returns total view minutes for a video.
func (vsp *VideoSharingPlatform) GetViews(videoID int) int {
	if v, ok := vsp.videos[videoID]; ok && !v.Removed {
		return v.Views
	}
	return -1
}

func main() {
	vsp := Constructor()

	id1 := vsp.Upload("video1")
	id2 := vsp.Upload("video2")
	fmt.Println("Uploaded IDs:", id1, id2) // 0, 1

	vsp.Watch(id1, 0, 5)
	vsp.Watch(id1, 2, 4)
	vsp.Watch(id2, 0, 10)

	vsp.Like(id1)
	vsp.Like(id1)
	vsp.Dislike(id1)

	fmt.Println("Likes/Dislikes for 0:", vsp.GetLikesAndDislikes(id1)) // [2, 1]
	fmt.Println("Views for 0:", vsp.GetViews(id1))                    // (5-0+1)+(4-2+1) = 6+3 = 9
	fmt.Println("Views for 1:", vsp.GetViews(id2))                    // 11

	vsp.Remove(id1)
	fmt.Println("After removal:", vsp.GetLikesAndDislikes(id1)) // [-1, -1]
}
```
