package go_libs_lawtech

import "sync"

// safe map
type SynchronizedMap struct {
	rw   *sync.RWMutex
	data map[interface{}]interface{}
}

// initialize
func NewSynchronizedMap() *SynchronizedMap {
	return &SynchronizedMap{
		rw:   new(sync.RWMutex),
		data: make(map[interface{}]interface{}),
	}
}

// save
func (sm *SynchronizedMap) Put(k, v interface{}) {
	sm.rw.Lock()
	defer sm.rw.Unlock()

	sm.data[k] = v
}

// retrieve
func (sm *SynchronizedMap) Get(k interface{}) interface{} {
	sm.rw.Lock()
	defer sm.rw.Unlock()

	return sm.data[k]
}

// delete
func (sm *SynchronizedMap) Delete(k interface{}) {
	sm.rw.Lock()
	defer sm.rw.Unlock()

	delete(sm.data, k)
}

// traverse
func (sm *SynchronizedMap) Each(cb func(interface{}, interface{})) {
	sm.rw.Lock()
	defer sm.rw.Unlock()

	for k, v := range sm.data {
		cb(k, v)
	}
}
