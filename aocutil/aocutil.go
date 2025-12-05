package aocutil

import "golang.org/x/exp/constraints"

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
	ret = []U{} //make([]U, len(arr))
	for _, val := range arr {
		ret = append(ret, mapper(val))
	}
	return
}

type Vector2[T constraints.Ordered] struct {
	X T
	Y T
}

func (self Vector2[T]) Compare(other Vector2[T]) int {
	if self.X != other.X {
		if self.X < other.X {
			return -1
		}
		return 1
	}
	if self.Y < other.Y {
		return -1
	} else if self.Y == other.Y {
		return 0
	}
	return 1
}

func Unique[T comparable](s []T) []T {
	inResult := make(map[T]bool)
	var result []T
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}
