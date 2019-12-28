package concurrency

type WebsiteChecker func(string) bool
type websiteCheckResult struct {
	url     string
	success bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	checkedWebsites := make(chan websiteCheckResult)

	for _, url := range urls {
		go func(u string) {
			checkedWebsites <- websiteCheckResult{url: u, success: wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-checkedWebsites
		results[result.url] = result.success
	}

	return results
}
