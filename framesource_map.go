package main

import (
	"github.com/amarburg/go-movieset"
	"sync"
)

//TODO.  is Printf on error the best way to get error information out?
type FrameSourceMap struct {
	sync.RWMutex
	nextId   int
	internal map[int]movieset.FrameSource
}

func (rm *FrameSourceMap) Load(key int) (movieset.FrameSource, bool) {
	rm.RLock()
	defer rm.RUnlock()

	value, ok := rm.internal[key]
	return value, ok
}

func (rm *FrameSourceMap) Delete(key int) {
	rm.Lock()
	defer rm.Unlock()
	delete(rm.internal, key)
}

func (rm *FrameSourceMap) Add(value movieset.FrameSource) int {
	rm.Lock()
	defer rm.Unlock()
	id := rm.nextId
	rm.internal[id] = value
	rm.nextId++
	return id
}
