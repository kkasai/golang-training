package popcount_test

import (
	"fmt"
	"github.com/golang-training/ch09/ex02/popcount"
	"sync"
	"testing"
)

func TestPopCount(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count := popcount.PopCount(0x1234567890ABCDEF)
			if count != 32 {
				panic(fmt.Sprintf("count is %d, want 20\n", count))
			}
		}()
	}

	wg.Wait()
}
