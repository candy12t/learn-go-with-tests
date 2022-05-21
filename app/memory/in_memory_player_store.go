package memory

import "github.com/candy12t/app/repository"

type InMemoryPlayerStore struct {
	store map[string]int
}

var _ repository.PlayerStore = &InMemoryPlayerStore{}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}
