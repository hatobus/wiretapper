package wiretapper

import "sync"

type WireTapper struct {
	mu sync.RWMutex
	callCounter uint64
}

func Initialize() *WireTapper {
	wt := &WireTapper{}
	wt.Clear()
	return wt
}

func (w *WireTapper) Clear() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.callCounter = 0
	return
}

func (w *WireTapper) GetCounter() uint64 {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.callCounter
}

func (w *WireTapper) Countup() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.callCounter++
	return
}
