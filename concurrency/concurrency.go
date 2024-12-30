package myconcurrency

type Checker func(string) bool
type CheckResults struct {
	url string
	res bool
}

func CheckWebSites(c Checker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChan := make(chan CheckResults)

	for _, v := range urls {
		go func(v string) {
			resultsChan <- CheckResults{url: v, res: c(v)}
		}(v)
	}

	for i := 0; i < len(urls); i++ {
		channelRes := <-resultsChan
		results[channelRes.url] = channelRes.res
	}
	return results
}
