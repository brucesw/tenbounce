package util

func Contains[T comparable](s []T, want T) bool {
	for _, have := range s {
		if have == want {
			return true
		}
	}

	return false
}

func Map[T_in comparable, T_out comparable](arr []T_in, fun func(T_in) T_out) []T_out {
	var ret = make([]T_out, len(arr))

	for i := range arr {
		ret[i] = fun(arr[i])
	}

	return ret
}
