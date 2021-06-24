package concurrency

import "net/http"

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsite(url string) bool {
	response, err := http.Head(url)
	return err == nil && response.StatusCode == http.StatusOK
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
