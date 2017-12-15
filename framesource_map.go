package main


import (
	"sync"
	"github.com/amarburg/go-frameset/framesource"
)

//TODO.  is Printf on error the best way to get error information out?
type FrameSourceMap struct {
	sync.RWMutex
	nextId				int
	internal			map[int]framesource.FrameSource
}

func (rm *FrameSourceMap) Load(key int) (framesource.FrameSource, bool) {
	rm.RLock()
	defer rm.RUnlock()

	value,ok := rm.internal[key]
	return value,ok
}

func (rm *FrameSourceMap) Delete(key int) {
	rm.Lock()
	defer rm.Unlock()
	delete(rm.internal, key)
}

func (rm *FrameSourceMap) Add(value framesource.FrameSource) int {
	rm.Lock()
	defer rm.Unlock()
	id := rm.nextId
	rm.internal[id] = value
	rm.nextId++
return id
}
