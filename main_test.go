package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSerial(t *testing.T) {

	start := time.Now()
	request, _ := makeBatch()
	for i := 0; i < 100; i++ {
		//postRequest()
		postMultiAsset(request)
	}
	elapsed := time.Since(start)
	fmt.Println("Time elapsed ", elapsed)
}

// 每一个post请求一个goroutine显然不合适！！！！！！！
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

//配置任务数和协程数量
var (
	numberOfWorkers = 3
	totalJobs       = 100
)

//使用channel实现线程池功能
func TestPool(t *testing.T) {
	request, _ := makeBatch()
	start := time.Now()
	pool := new(GoroutinePool)
	pool.Init(numberOfWorkers, totalJobs)

	for i := 0; i < totalJobs; i++ {
		pool.AddTask(func() error {
			//return postRequest()
			return postMultiAsset(request)
		})
	}

	isFinished := false

	pool.SetFinishCallback(func() {
		func(isFinished *bool) {
			*isFinished = true
		}(&isFinished)

	})

	pool.Start()

	for !isFinished {
		time.Sleep(time.Millisecond * 100)
	}

	pool.Stop()
	fmt.Println("All is well !")
	elapsed := time.Since(start)
	fmt.Println("Time elapsed ", elapsed)

}
