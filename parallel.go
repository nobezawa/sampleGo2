package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	wait := new(sync.WaitGroup)
	urls := []string{
		"http://example.com",
		"http://example.net",
		"http://example.org",
	}

	for _, url := range urls {
		wait.Add(1)
		// 取得処理をゴルーチンで実行
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}

			defer res.Body.Close()
			fmt.Println(url, res.Status)
			wait.Done()
		}(url)
	}
	wait.Wait()
	// main()が終わらないように待ち合わせる
	//time.Sleep(time.Second)

}
