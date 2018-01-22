package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSerial(t *testing.T) {
	start := time.Now()
	for i := 0; i < 10; i++ {
		postRequest()
	}
	elapsed := time.Since(start)
	fmt.Println("Time elapsed ", elapsed)
}

func TestParallel(t *testing.T) {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			postRequest()
			wg.Done()
		}()
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Time elapsed ", elapsed)

}
