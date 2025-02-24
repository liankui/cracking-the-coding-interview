package cracking_the_coding_interview

import "container/heap"

// TODO: 未完成
type CacheItem struct {
	key   int
	val   int
	freq  int
	index int
}
type LFUCache struct {
	cap     int
	cache   map[int]int
	freqMap map[int]int
	minHeap *MinHeap
}

func NewLFUCache(capacity int) *LFUCache {
	return &LFUCache{
		cap:     capacity,
		cache:   make(map[int]int),
		freqMap: make(map[int]int),
		minHeap: &MinHeap{},
	}
}

type MinHeap []*CacheItem

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].freq < h[j].freq }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*CacheItem))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (l *LFUCache) Get(key int) int {
	if item, ok := l.cache[key]; ok {
		l.incrementFreq(key)
		return item
	}
	return -1
}

func (l *LFUCache) incrementFreq(key int) {
	for _, it := range *l.minHeap {
		if it.key == key {
			it.freq++
			heap.Fix(l.minHeap, it.index)
			break
		}
	}
}

func (l *LFUCache) Put(key int, value int) {
	// if l.cap == 0 {
	// 	return
	// }
	//
	// if item, ok := l.cache[key]; ok {
	// 	l.cache[key] = value
	// }
}
