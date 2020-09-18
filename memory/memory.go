package memory

import "sync"

type Memory struct {
	mu sync.RWMutex
	m map[string]interface{}
}

var NewMemory = func() Memory {
	return Memory{
		m: make(map[string]interface{}),
	}
}

func (m *Memory) Set(key string, val interface{}) {
	m.mu.Lock()
	m.m[key] = val
	m.mu.Unlock()
}

func (m *Memory) Get(Key string) interface{} {
	m.mu.RLock()
	val := m.m[Key]
	m.mu.RUnlock()
	return val
}