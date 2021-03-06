package main

import (
	"sync"
	"time"
)

type SpamGuard struct {
	duration time.Duration
	posts    map[string]time.Time
	mutex    *sync.Mutex
}

func NewSpamGuard(duration string) *SpamGuard {
	d, err := time.ParseDuration(duration)
	if err != nil {
		panic(err)
	}
	return &SpamGuard{
		duration: d,
		posts:    make(map[string]time.Time),
		mutex:    &sync.Mutex{},
	}
}

func (sg *SpamGuard) CanPost(id string) bool {
	result := true
	now := time.Now()
	sg.mutex.Lock()
	expires, found := sg.posts[id]
	if found {
		if expires.After(now) {
			// Blocked
			result = false
		}
	} else {
		// Add to block
		sg.posts[id] = now.Add(sg.duration)
	}
	sg.clean(now)
	sg.mutex.Unlock()
	return result
}

func (sg *SpamGuard) clean(now time.Time) {
	for key, expires := range sg.posts {
		if expires.Before(now) {
			delete(sg.posts, key)
		}
	}
}
