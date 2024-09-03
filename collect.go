package collect

func Each[T any](xs []T, f func(T, int)) {
	for i, x := range xs {
		f(x, i)
	}
}

func Map[T any, U any](xs []T, f func(T, int) U) []U {
	ys := make([]U, len(xs))
	for i, x := range xs {
		ys[i] = f(x, i)
	}
	return ys
}

func Filter[T any](xs []T, f func(T, int) bool) []T {
	ys := make([]T, 0, len(xs))
	for i, x := range xs {
		if f(x, i) {
			ys = append(ys, x)
		}
	}
	return ys
}

func Find[T any](xs []T, f func(T) bool) (T, bool) {
	for _, x := range xs {
		if f(x) {
			return x, true
		}
	}
	var zero T
	return zero, false
}

func FindIndex[T any](xs []T, f func(T) bool) int {
	for i, x := range xs {
		if f(x) {
			return i
		}
	}
	return -1
}

func FindLast[T any](xs []T, f func(T) bool) (T, bool) {
	for i := len(xs) - 1; i >= 0; i-- {
		if f(xs[i]) {
			return xs[i], true
		}
	}
	var zero T
	return zero, false
}

func FindLastIndex[T any](xs []T, f func(T) bool) int {
	for i := len(xs) - 1; i >= 0; i-- {
		if f(xs[i]) {
			return i
		}
	}
	return -1
}

func Count[T any](xs []T, f func(T, int) bool) int {
	count := 0
	for i, x := range xs {
		if f(x, i) {
			count++
		}
	}
	return count
}

func Some[T any](xs []T, f func(T, int) bool) bool {
	for i, x := range xs {
		if f(x, i) {
			return true
		}
	}
	return false
}

func Every[T any](xs []T, f func(T, int) bool) bool {
	for i, x := range xs {
		if !f(x, i) {
			return false
		}
	}
	return true
}

func None[T any](xs []T, f func(T, int) bool) bool {
	for i, x := range xs {
		if f(x, i) {
			return false
		}
	}
	return true
}

func Concat[T any](xs []T, ys []T) []T {
	zs := make([]T, 0, len(xs)+len(ys))
	zs = append(zs, xs...)
	zs = append(zs, ys...)
	return zs
}

func ConcatMap[T any, U any](xs []T, f func(T) []U) []U {
	ys := make([]U, 0, len(xs))
	for _, x := range xs {
		ys = append(ys, f(x)...)
	}
	return ys
}

func Reverse[T any](xs []T) []T {
	zs := make([]T, len(xs))
	for i, x := range xs {
		zs[len(xs)-i-1] = x
	}
	return zs
}

func Uniq[T comparable](xs []T) []T {
	m := make(map[T]struct{}, len(xs))
	for _, x := range xs {
		m[x] = struct{}{}
	}
	zs := make([]T, 0, len(m))
	for x := range m {
		zs = append(zs, x)
	}
	return zs
}

func UniqBy[T any, U comparable](xs []T, f func(T) U) []T {
	m := make(map[U]T, len(xs))
	for _, x := range xs {
		m[f(x)] = x
	}
	zs := make([]T, 0, len(m))
	for _, x := range m {
		zs = append(zs, x)
	}
	return zs
}
