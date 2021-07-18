package main

import "runtime"

func main() {
	n := runtime.NumCPU()
	for i := 0; i < n-1; i++ {
		go func() {
			for {
			}
		}()
	}
	for {
	}
}
