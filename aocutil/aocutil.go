package aocutil

func AtOrElse[T any](s []T, i int, fallback T) T {
	if i >= 0 && i < len(s) {
		return s[i]
	}
	return fallback
}

func Filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func Map[T any, U any](arr []T, mapper func(T) U) (ret []U) {
	ret = make([]U, len(arr))
	for _, val := range arr {
		ret = append(ret, mapper(val))
	}
	return
}
