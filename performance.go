package time

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

func ConcurrentPerformance(maxCount int64, fn func()) {
	const (
		SecPerRound = 5
	)

	for thCount := int64(1); thCount <= maxCount; thCount++ {
		fmt.Printf("Measuring performance with %d goroutines...\n", thCount)

		var totalOps int64
		ctx, cancel := context.WithCancel(context.Background())

		for i := int64(0); i < thCount; i++ {
			go func() {
				for {
					select {
					case <-ctx.Done():
						return
					default:
						fn()
						atomic.AddInt64(&totalOps, 1)
					}
				}
			}()
		}

		time.Sleep(time.Duration(SecPerRound) * time.Second)
		cancel()

		totalOpsFinal := atomic.LoadInt64(&totalOps)
		throughput := float64(totalOpsFinal) / float64(SecPerRound)
		fmt.Printf("Goroutines: %2d, Total Ops: %8d, Throughput: %8.2f ops/s\n", thCount, totalOpsFinal, throughput)
	}
}
