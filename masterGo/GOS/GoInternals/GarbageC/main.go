package main

import (
	"fmt"
	"runtime"
	"time"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("Mem.Alloc: ", mem.Alloc)
	fmt.Println("Mem.TotalAlloc: ", mem.TotalAlloc)
	fmt.Println("Mem.HeapAlloc: ", mem.HeapAlloc)
	fmt.Println("mem.NumGC: ", mem.NumGC)
	fmt.Println("----------")
}

func main() {
	var mem runtime.MemStats
	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed")
		}

	}

	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("Operation failed")
		}
		time.Sleep(5 * time.Second)
	}

	printStats(mem)
}
