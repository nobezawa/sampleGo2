package main

import (
	"fmt"
	"log"
	"net/http"
)

func getStatus(urls []string) <-chan string {
	statusChan := make(chan string)
	for _, url := range urls {
		// 取得処理をゴルーチンで実行
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}

			defer res.Body.Close()
			statusChan <- res.Status
		}(url)
	}

	return statusChan

}

func main() {

	urls := []string{
		"http://example.com",
		"http://example.net",
		"http://example.org",
	}

	statusChan := getStatus(urls)

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-statusChan)
	}
}
