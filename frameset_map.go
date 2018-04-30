package main

import (
	"github.com/amarburg/go-frameset/frameset"
	"sync"
)

type FrameSetMap struct {
	sync.RWMutex
	nextId   int
	internal map[int]*frameset.FrameSet
}

func (rm *FrameSetMap) Load(key int) (*frameset.FrameSet, bool) {
	rm.RLock()
	defer rm.RUnlock()

	value, ok := rm.internal[key]
	return value, ok
}

func (rm *FrameSetMap) Delete(key int) {
	rm.Lock()
	defer rm.Unlock()
	delete(rm.internal, key)
}

func (rm *FrameSetMap) Add(value *frameset.FrameSet) int {
	rm.Lock()
	defer rm.Unlock()
	id := rm.nextId
	rm.internal[id] = value
	rm.nextId++
	return id
}
