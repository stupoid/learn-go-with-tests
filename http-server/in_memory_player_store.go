package main

import "sync"

type InMemoryPlayerStore struct {
	store map[string]int
	mu    sync.RWMutex
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}, sync.RWMutex{}}
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.mu.Lock()
	defer i.mu.Unlock()
	return i.store[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetLeague() (league []Player) {
	for name, value := range i.store {
		league = append(league, Player{name, value})
	}
	return
}
