package echotron

import "sync"

type smap[K, V any] sync.Map

func (s *smap[K, V]) load(key K) (V, bool) {
	val, ok := (*sync.Map)(s).Load(key)
	return val.(V), ok
}

func (s *smap[K, V]) store(key K, val V) {
	(*sync.Map)(s).Store(key, val)
}

func (s *smap[K, V]) delete(key K) {
	(*sync.Map)(s).Delete(key)
}
