package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func publishToChan (resultChannel chan result, url string, wc WebsiteChecker){
	resultChannel <- result{url, wc(url)}
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go publishToChan(resultChannel, url, wc)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
