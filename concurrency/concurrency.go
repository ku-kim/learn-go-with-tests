package concurrency

import (
	"fmt"
	"sync"
)

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	var wg sync.WaitGroup
	//var mu sync.Mutex // mutex

	results := make(map[string]bool)

	wg.Add(len(urls))
	for _, url := range urls {
		go func() {
			fmt.Println(url) // for문의 url이 재사용되기 때문에 끝에 값에만 goroutine이 돌고 있는 문제 발생
			defer wg.Done()
			//mu.Lock()
			results[url] = wc(url)
			//mu.Unlock()
		}()
	}

	wg.Wait()
	return results
}
