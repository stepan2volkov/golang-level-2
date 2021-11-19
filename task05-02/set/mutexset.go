package set

import "sync"

type MutexSet struct {
	sync.Mutex
	values map[float64]bool
}

func (s *MutexSet) Add(i float64) {
	s.Lock()
	defer s.Unlock()
	s.values[i] = true
}

func (s *MutexSet) Has(i float64) bool {
	s.Lock()
	defer s.Unlock()
	return s.values[i]
}

func NewMutexSet() Set {
	return &MutexSet{values: make(map[float64]bool)}
}
