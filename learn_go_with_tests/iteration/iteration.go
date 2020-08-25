package iteration

func Repeat(c string, num int) string {
	var repeated string
	for i := 0; i < num; i++ {
		repeated = repeated + c
	}
	return repeated
}
