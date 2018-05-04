package main

import (
	"github.com/amarburg/go-movieset"
	"sync"
)

//TODO.  is Printf on error the best way to get error information out?
type SequentialMap struct {
	sync.RWMutex
	nextId   int
	internal map[int]movieset.Sequential
}

func (rm *SequentialMap) Load(key int) (movieset.Sequential, bool) {
	rm.RLock()
	defer rm.RUnlock()

	value, ok := rm.internal[key]
	return value, ok
}

func (rm *SequentialMap) Delete(key int) {
	rm.Lock()
	defer rm.Unlock()
	delete(rm.internal, key)
}

func (rm *SequentialMap) Add(value movieset.Sequential) int {
	rm.Lock()
	defer rm.Unlock()
	id := rm.nextId
	rm.internal[id] = value
	rm.nextId++
	return id
}
