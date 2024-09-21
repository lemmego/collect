package collect

import (
	"errors"
)

type SliceCollection[T any] struct {
	items []T
}

type MapCollection[K comparable, V any] struct {
	items map[K]V
}

func NewSlice[T any](xs []T) *SliceCollection[T] {
	return &SliceCollection[T]{items: xs}
}

func NewMap[K comparable, V any](xs map[K]V) *MapCollection[K, V] {
	return &MapCollection[K, V]{items: xs}
}

func (sc *SliceCollection[T]) Get() []T {
	return sc.Items()
}

func (sc *SliceCollection[T]) All() []T {
	return sc.Items()
}

func (sc *SliceCollection[T]) Items() []T {
	return sc.items
}

func (mc *MapCollection[K, V]) Items() map[K]V {
	return mc.items
}

func (mc *MapCollection[K, V]) Get() map[K]V {
	return mc.Items()
}

func (mc *MapCollection[K, V]) All() map[K]V {
	return mc.Items()
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
	keys := make([]K, 0, len(mc.items))
	for k := range mc.items {
		keys = append(keys, k)
	}

	return NewSlice(keys)
}

func (mc *MapCollection[K, V]) Values() *SliceCollection[V] {
	values := make([]V, 0, len(mc.items))
	for _, v := range mc.items {
		values = append(values, v)
	}

	return NewSlice(values)
}

func (sc *SliceCollection[T]) Each(f func(T, int)) *SliceCollection[T] {
	Each(sc.Items(), f)
	return sc
}

func (mc *MapCollection[K, V]) Each(f func(V, K)) *MapCollection[K, V] {
	for k, v := range mc.items {
		f(v, k)
	}
	return mc
}

func (sc *SliceCollection[T]) Filter(f func(T, int) bool) *SliceCollection[T] {
	sc.items = Filter(sc.Items(), f)
	return sc
}

func (mc *MapCollection[K, V]) Filter(f func(V, K) bool) *MapCollection[K, V] {
	filtered := make(map[K]V, len(mc.items))
	for k, v := range mc.items {
		if f(v, k) {
			filtered[k] = v
		}
	}

	mc.items = filtered

	return mc
}

func (sc *SliceCollection[T]) Map(f func(T, int) T) *SliceCollection[T] {
	sc.items = Map(sc.Items(), f)
	return sc
}

func (mc *MapCollection[K, V]) Map(f func(V, K) V) *MapCollection[K, V] {
	mapped := make(map[K]V, len(mc.items))
	for k, v := range mc.items {
		mapped[k] = f(v, k)
	}

	mc.items = mapped

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
	for _, v := range mc.items {
		if f(v) {
			return v, nil
		}
	}
	var zero V
	return zero, errors.New("not found")
}

func (sc *SliceCollection[T]) Reduce(f func(T, T, int) T, initial T) T {
	return Reduce(sc.Items(), f, initial)
}

func (mc *MapCollection[K, V]) Reduce(f func(V, V, K) V, initial V) V {
	for k, x := range mc.items {
		initial = f(initial, x, k)
	}

	return initial
}

func (sc *SliceCollection[T]) FindIndex(f func(T) bool) int {
	return FindIndex(sc.Items(), f)
}

func (sc *SliceCollection[T]) FindLast(f func(T) bool) (T, bool) {
	return FindLast(sc.Items(), f)
}

func (sc *SliceCollection[T]) FindLastIndex(f func(T) bool) int {
	return FindLastIndex(sc.Items(), f)
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

func Reduce[T any](xs []T, f func(T, T, int) T, initial T) T {
	for k, x := range xs {
		initial = f(initial, x, k)
	}
	return initial
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
