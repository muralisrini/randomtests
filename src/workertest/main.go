package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
	"workertest/workerpool"
)

var N,M int

var verbose bool

func log(s string) {
	if verbose {
		fmt.Printf("%s", s)
	}
}

func workWithPool(async bool, wp workerpool.WorkerPool) {
	workSlice := make([]func(), 0, N)
	results := make(chan int)
	for i := 0; i < N;  i++ {
		iLcl := i
		workSlice = append(workSlice, func() {
					results <- M*iLcl
				})
	}

	if async {
		wp.AssignWorkAsync(workSlice...)
	} else {
		wp.AssignWorkSync(workSlice...)
	}
	
	log("Pool result: ")
	for i := 0; i < len(workSlice); i++ {
		res := <-results
		log(fmt.Sprintf("%d ", res))
	}
	log("\n")
}

func workGoroutines() {
	var wg sync.WaitGroup
	wg.Add(N)
	results := make([]int, N)
	for i := 0; i < N;  i++ {
		iLcl := i
		go func() {
			results[iLcl] = M*iLcl
			wg.Done()
		} ()
	}
	wg.Wait()
	log("Goroutines result: ")
	for i := 0; i < N; i++ {
		log(fmt.Sprintf("%d ", results[i]))
	}
	log("\n")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//if N is large this setting to true will output a ton of stuff
	verbose = false

	//change this to get different results
	N = 10000
	M = 10

	//don't count initialization for computing elapsed time
	wp, _ := workerpool.NewGenericWorkerPool(N)

	start := time.Now()
	workWithPool(true, wp)
	elapsed := time.Since(start)
	fmt.Printf("Elapsed time using async pool %s\n", elapsed)

	start = time.Now()
	workWithPool(false, wp)
	elapsed = time.Since(start)
	fmt.Printf("Elapsed time using sync pool %s\n", elapsed)

	start = time.Now()
	workGoroutines()
	elapsed = time.Since(start)
	fmt.Printf("Elapsed time using goroutine %s\n", elapsed)
}
