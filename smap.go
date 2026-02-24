package echotron

import "sync"

type smap[K, V any] sync.Map

func (s *smap[K, V]) load(key K) (val V, ok bool) {
	v, ok := (*sync.Map)(s).Load(key)
	if !ok {
		return
	}
	return v.(V), ok
}

func (s *smap[K, V]) store(key K, val V) {
	(*sync.Map)(s).Store(key, val)
}

func (s *smap[K, V]) loadOrStore(key K, val V) (actual V, loaded bool) {
	a, loaded := (*sync.Map)(s).LoadOrStore(key, val)
	return a.(V), loaded
}

func (s *smap[K, V]) delete(key K) {
	(*sync.Map)(s).Delete(key)
}
