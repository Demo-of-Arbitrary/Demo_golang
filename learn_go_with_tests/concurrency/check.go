package concurrency

type Checker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(c Checker, websites []string) map[string]bool {
	ret := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range websites {
		go func(s string) {
			resultChannel <- result{s, c(s)}
		}(url)
	}

	for i := 0; i < len(websites); i++ {
		result := <-resultChannel
		ret[result.string] = result.bool
	}
	return ret
}
