package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	t := time.Now()
	for i := 0; i < 10; i++ {
		go getResp()
	}
	time.Sleep(time.Second)
	fmt.Println(time.Since(t))
}

func getResp() {
	resp, err := http.Get("HTTP://google.com")
	if err != nil {
		fmt.Printf("Ошибка %s\n", err.Error())
	}
	fmt.Printf("Код: %d\n", resp.StatusCode)
}
