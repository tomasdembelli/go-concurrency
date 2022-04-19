package context_switch

import (
	"sync"
	"testing"
)

func BenchmarkName(b *testing.B) {
	var begin = make(chan struct{}, 3)
	var comChan = make(chan struct{})
	var token struct{}
	var wg sync.WaitGroup
	wg.Add(2)
	sender := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			comChan <- token
		}
	}
	receiver := func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			<-comChan
		}
	}

	go sender()
	go receiver()
	b.StartTimer()
	close(begin)
	wg.Wait()
}
