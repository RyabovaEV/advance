package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	t := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			getResp()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(time.Since(t))
}

func getResp() {
	resp, err := http.Get("HTTP://google.com")
	if err != nil {
		fmt.Printf("Ошибка %s\n", err.Error())
	}
	fmt.Printf("Код: %d\n", resp.StatusCode)
}
