package concurrency

type Checker func(string) bool

func CheckWebsites(c Checker, websites []string) map[string]bool {
	ret := make(map[string]bool)
	for _, url := range websites {
		ret[url] = c(url)
	}
	return ret
}
