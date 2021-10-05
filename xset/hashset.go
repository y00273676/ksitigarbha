package xset

import (
	"sync"
)

type threadUnsafeHashSet map[interface{}]struct{}

func newThreadUnsafeHashSet(size int) threadUnsafeHashSet {
	if size > 0 {
		return make(threadUnsafeHashSet, size)
	}
	return make(threadUnsafeHashSet)
}

func (s *threadUnsafeHashSet) Add(o interface{}) bool {
	_, ok := (*s)[o]
	if ok {
		return false
	}

	(*s)[o] = struct{}{}
	return true
}

func (s *threadUnsafeHashSet) Length() int {
	return len(*s)
}

func (s *threadUnsafeHashSet) Clear() {
	*s = newThreadUnsafeHashSet(0)
}

func (s *threadUnsafeHashSet) Clone() Set {
	clone := newThreadUnsafeHashSet(s.Length())
	for o := range *s {
		clone.Add(o)
	}
	return &clone
}

func (s *threadUnsafeHashSet) Contains(o ...interface{}) bool {
	for _, v := range o {
		if _, ok := (*s)[v]; !ok {
			return false
		}
	}
	return true
}

func (s *threadUnsafeHashSet) Remove(o interface{}) {
	delete(*s, o)
}

func (s *threadUnsafeHashSet) ToSlice() []interface{} {
	dist := make([]interface{}, 0, s.Length())
	for o := range *s {
		dist = append(dist, o)
	}

	return dist
}

func (s *threadUnsafeHashSet) Iterate(fn IterateFn) {
	for o := range *s {
		if stop := fn(o); stop {
			break
		}

	}
}

func (s *threadUnsafeHashSet) Intersect(other Set) Set {
	intersection := newThreadUnsafeHashSet(0)

	s.Iterate(func(o interface{}) (stop bool) {
		if other.Contains(o) {
			intersection.Add(o)
		}
		return false
	})

	return &intersection
}

func (s *threadUnsafeHashSet) Union(other Set) Set {
	union := newThreadSafeHashSet(0)
	s.Iterate(func(o interface{}) (stop bool) {
		union.Add(o)
		return false
	})
	other.Iterate(func(o interface{}) (stop bool) {
		union.Add(o)
		return false
	})

	return &union
}

func (s *threadUnsafeHashSet) Difference(other Set) Set {
	diff := newThreadSafeHashSet(0)
	s.Iterate(func(o interface{}) (stop bool) {
		if !other.Contains(o) {
			diff.Add(o)
		}
		return false
	})

	return &diff
}

type threadSafeHashSet struct {
	s threadUnsafeHashSet
	sync.RWMutex
}

func newThreadSafeHashSet(size int) threadSafeHashSet {
	return threadSafeHashSet{
		s: newThreadUnsafeHashSet(size),
	}
}

func (s *threadSafeHashSet) Add(o interface{}) bool {
	s.Lock()
	defer s.Unlock()
	ok := s.s.Add(o)
	return ok
}

func (s *threadSafeHashSet) Length() int {
	s.RLock()
	defer s.RUnlock()
	return s.s.Length()
}

func (s *threadSafeHashSet) Clear() {
	s.Lock()
	defer s.Unlock()
	s.s = newThreadUnsafeHashSet(0)
}

func (s *threadSafeHashSet) Clone() Set {
	s.RLock()
	defer s.RUnlock()
	unsafeClone := s.s.Clone().(*threadUnsafeHashSet)
	return &threadSafeHashSet{
		s: *unsafeClone,
	}
}

func (s *threadSafeHashSet) Contains(o ...interface{}) bool {
	s.RLock()
	defer s.RUnlock()

	return s.s.Contains(o...)
}

func (s *threadSafeHashSet) Remove(o interface{}) {
	s.Lock()
	defer s.Unlock()

	s.s.Remove(o)
}

func (s *threadSafeHashSet) ToSlice() []interface{} {
	s.RLock()
	defer s.RUnlock()

	return s.s.ToSlice()
}

func (s *threadSafeHashSet) Iterate(fn IterateFn) {
	s.RLock()
	defer s.RUnlock()

	s.s.Iterate(fn)
}

func (s *threadSafeHashSet) Intersect(other Set) Set {
	s.RLock()
	defer s.RUnlock()

	safe, is := other.(*threadSafeHashSet)
	if is {
		safe.RLock()
		defer safe.RUnlock()
	}

	unsafeIntersection := s.s.Intersect(other).(*threadUnsafeHashSet)
	return &threadSafeHashSet{
		s: *unsafeIntersection,
	}

}

func (s *threadSafeHashSet) Union(other Set) Set {
	s.RLock()
	defer s.RUnlock()

	safe, is := other.(*threadSafeHashSet)
	if is {
		safe.RLock()
		defer safe.RUnlock()
	}

	return s.s.Union(other)
}

func (s *threadSafeHashSet) Difference(other Set) Set {
	s.RLock()
	defer s.RUnlock()

	safe, is := other.(*threadSafeHashSet)
	if is {
		safe.RLock()
		defer safe.RUnlock()
	}

	return s.s.Difference(other)
}
