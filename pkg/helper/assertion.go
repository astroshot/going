package helper

// Assert panics when val is false
func Assert(val bool, msg string) {
	if !val {
		panic(msg)
	}
}
