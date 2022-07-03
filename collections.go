package refined

import (
	"container/list"
	"reflect"
)

type Empty = SizeEqualTo[N0[int]]
type NonEmpty = SizeGreaterThan[N0[int]]

type SizeLessThan[N TNumber[int]] struct{}
type SizeGreaterThan[N TNumber[int]] struct{}
type SizeEqualTo[N TNumber[int]] struct{}

func (n SizeLessThan[N]) applyPredicate(a any) (bool, error) {
	var size N
	return applySizePredicateOn(a, func(len int) bool { return len < size.Of() }, n)
}
func (n SizeGreaterThan[N]) applyPredicate(a any) (bool, error) {
	var size N
	return applySizePredicateOn(a, func(len int) bool { return len > size.Of() }, n)
}
func (n SizeEqualTo[N]) applyPredicate(a any) (bool, error) {
	var size N
	return applySizePredicateOn(a, func(len int) bool { return len == size.Of() }, n)
}

func applySizePredicateOn[P Predicate](a any, fn func(int) bool, p P) (bool, error) {
	switch a := a.(type) {
	case list.List:
		return fn(a.Len()), nil
	case string:
		return fn(len(a)), nil
	default:
		v := reflect.ValueOf(a)
		switch v.Kind() {
		case reflect.Array:
			return fn(v.Len()), nil
		case reflect.Map:
			return fn(v.Len()), nil
		case reflect.Slice:
			return fn(v.Len()), nil
		}
		return false, valueTypeError(a, p)
	}
}
