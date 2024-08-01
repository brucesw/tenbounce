package util

func Contains[T comparable](s []T, want T) bool {
	for _, have := range s {
		if have == want {
			return true
		}
	}

	return false
}
