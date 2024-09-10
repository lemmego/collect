package collect

import "errors"

type SliceCollection[T any] struct {
	originalItems []T
	items         []T
}

type MapCollection[K comparable, V any] struct {
	originalItems map[K]V
	items         map[K]V
}

func NewSlice[T any](xs []T) *SliceCollection[T] {
	return &SliceCollection[T]{originalItems: xs, items: xs}
}

func NewMap[K comparable, V any](xs map[K]V) *MapCollection[K, V] {
	return &MapCollection[K, V]{originalItems: xs, items: xs}
}

func (sc *SliceCollection[T]) Get() []T {
	return sc.Items()
}

func (sc *SliceCollection[T]) All() []T {
	return sc.originalItems
}

func (sc *SliceCollection[T]) Items() []T {
	if sc.items == nil {
		sc.items = make([]T, len(sc.originalItems))
		copy(sc.items, sc.originalItems)
	}
	return sc.items
}

func (mc *MapCollection[K, V]) Items() map[K]V {
	if mc.items == nil {
		mc.items = make(map[K]V, len(mc.originalItems))
		for k, v := range mc.originalItems {
			mc.items[k] = v
		}
	}
	return mc.items
}

func (mc *MapCollection[K, V]) Get() map[K]V {
	return mc.Items()
}

func (mc *MapCollection[K, V]) All() map[K]V {
	return mc.originalItems
}

func (sc *SliceCollection[T]) Len() int {
	return len(sc.Items())
}

func (mc *MapCollection[K, V]) Len() int {
	return len(mc.Items())
}

func (sc *SliceCollection[T]) At(i int) T {
	return sc.Items()[i]
}

func (mc *MapCollection[K, V]) At(i K) V {
	return mc.Items()[i]
}

func (mc *MapCollection[K, V]) Keys() *SliceCollection[K] {
	keys := make([]K, 0, len(mc.originalItems))
	for k := range mc.originalItems {
		keys = append(keys, k)
	}

	return NewSlice(keys)
}

func (mc *MapCollection[K, V]) Values() *SliceCollection[V] {
	values := make([]V, 0, len(mc.originalItems))
	for _, v := range mc.originalItems {
		values = append(values, v)
	}

	return NewSlice(values)
}

func (sc *SliceCollection[T]) Each(f func(T, int)) *SliceCollection[T] {
	Each(sc.Items(), f)
	return sc
}

func (mc *MapCollection[K, V]) Each(f func(V, K)) *MapCollection[K, V] {
	for k, v := range mc.originalItems {
		f(v, k)
	}
	return mc
}

func (sc *SliceCollection[T]) Filter(f func(T, int) bool) *SliceCollection[T] {
	sc.items = Filter(sc.Items(), f)
	return sc
}

func (mc *MapCollection[K, V]) Filter(f func(V, K) bool) *MapCollection[K, V] {
	mc.items = make(map[K]V, len(mc.originalItems))
	for k, v := range mc.originalItems {
		if f(v, k) {
			mc.items[k] = v
		}
	}
	return mc
}

func (sc *SliceCollection[T]) Map(f func(T, int) T) *SliceCollection[T] {
	sc.items = Map(sc.Items(), f)
	return sc
}

func (mc *MapCollection[K, V]) Map(f func(V, K) V) *MapCollection[K, V] {
	mc.items = make(map[K]V, len(mc.originalItems))
	for k, v := range mc.originalItems {
		mc.items[k] = f(v, k)
	}
	return mc
}

func (sc *SliceCollection[T]) Find(f func(T) bool) (T, error) {
	val, found := Find(sc.Items(), f)
	if found {
		return val, nil
	}

	var zero T
	return zero, errors.New("not found")
}

func (mc *MapCollection[K, V]) Find(f func(V) bool) (V, error) {
	for _, v := range mc.originalItems {
		if f(v) {
			return v, nil
		}
	}
	var zero V
	return zero, errors.New("not found")
}

func (sc *SliceCollection[T]) Reduce(f func(T, T) T, zero T) T {
	return Reduce(sc.Items(), f, zero)
}

// ================== Base Functions ==================

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

func Reduce[T any](xs []T, f func(T, T) T, zero T) T {
	for _, x := range xs {
		zero = f(zero, x)
	}
	return zero
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
