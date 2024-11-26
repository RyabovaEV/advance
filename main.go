package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	//arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	arr := make([]int, 100000000)
	for i := 0; i < 100000000; i++ {
		arr[i] = i + 1
	}
	numGorutines := 3
	partSize := len(arr) / numGorutines
	ch := make(chan int, numGorutines)

	for i := 0; i < numGorutines; i++ {
		start := partSize * i
		end := start + partSize
		go summArray(arr[start:end], ch)
	}
	var summ int

	for i := 0; i < numGorutines; i++ {
		summ += <-ch
	}

	fmt.Printf("Сумма: %d\n", summ)
	fmt.Println(time.Since(t))
}

func summArray(arr []int, summCh chan int) {
	summPart := 0
	for _, value := range arr {
		summPart += value
	}
	summCh <- summPart
}
