package iteration

func Iteration(c string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated = repeated + c
	}
	return repeated
}
