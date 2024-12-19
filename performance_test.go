package time

import (
	"testing"
)

func TestPerformance(t *testing.T) {
	ConcurrentPerformance(100, func() {
		a := 1
		b := 2
		c := a + b
		_ = c
		for i := 0; i < 1000; i++ {
			for j := 2; j < 1000; j++ {
				if i%j == 0 {
					break
				}
			}
		}
	})
}
