package main

import (
	"fmt"
	"sync"
)

type safeCounter struct {
	counts map[string]int
	mux *sync.Mutex

}
func (sc safeCounter) inc(key string) {
	sc.mux.Lock()
	defer sc.mux.Unlock()
	sc.slowIncrement(key)
}
func (sc safeCounter) val(key string) int {
	sc.mux.Lock()
	defer sc.mux.Unlock()
	return sc.counts[key]
}
func (sc safeCounter) slowIncrement(key string) {
	sc.counts[key]++	
}

func main() {

	fmt.Println("hi")
}