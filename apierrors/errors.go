package apierrors



func IsOneOf(err error, expects ...error) bool {
	for _, e := range expects {
		if e == err {
			return true
		}
	}
	return false
}