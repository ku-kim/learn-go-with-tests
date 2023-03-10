package concurrency

import (
	"sync"
)

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	var wg sync.WaitGroup
	//var mu sync.Mutex // mutex

	results := make(map[string]bool)

	wg.Add(len(urls))
	for _, url := range urls {
		go func(u string) {
			defer wg.Done()
			//mu.Lock()
			results[u] = wc(u)
			//mu.Unlock()
		}(url)
	}

	wg.Wait()
	return results
}
