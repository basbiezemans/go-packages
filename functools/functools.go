// Package functools provides generic higher-order functions.
package functools

// Any, applied to a predicate and a slice, determines whether any element of
// the slice satisfies the predicate.
func Any[A any](fn func(A) bool, xs []A) bool {
	for _, x := range xs {
		if fn(x) {
			return true
		}
	}
	return false
}

// All, applied to a predicate and a slice, determines whether all elements of
// the slice satisfy the predicate.
func All[A any](fn func(A) bool, xs []A) bool {
	for _, x := range xs {
		if !fn(x) {
			return false
		}
	}
	return true
}

// Reduce, applied to a reducer function and a slice, reduces the slice to a
// single value. It is also known as a left fold, because it "folds" a data
// structure from left to right.
func Reduce[A, B any](fn func(B, A) B, xs []A, initValue B) B {
	acc := initValue
	for _, x := range xs {
		acc = fn(acc, x)
	}
	return acc
}

// Apply (a.k.a. "map"), applies a unary function to each element of a slice.
func Apply[A, B any](fn func(A) B, xs []A) []B {
	var ys = make([]B, len(xs))
	for i, x := range xs {
		ys[i] = fn(x)
	}
	return ys
}

// Filter, applied to a predicate and a slice, filters the slice of those
// elements that satisfy the predicate.
func Filter[A any](fn func(A) bool, xs []A) []A {
	var ys = make([]A, 0, len(xs)/2)
	for _, x := range xs {
		if fn(x) {
			ys = append(ys, x)
		}
	}
	return ys
}
