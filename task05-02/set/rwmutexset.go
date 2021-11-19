package set

import "sync"

type RWMutexSet struct {
	sync.RWMutex
	values map[float64]bool
}

func (s *RWMutexSet) Add(i float64) {
	s.Lock()
	defer s.Unlock()
	s.values[i] = true
}

func (s *RWMutexSet) Has(i float64) bool {
	s.RLock()
	defer s.RUnlock()
	return s.values[i]
}

func NewRWMutexSet() Set {
	return &RWMutexSet{values: make(map[float64]bool)}
}
