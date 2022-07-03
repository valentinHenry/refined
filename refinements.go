package refined

import (
	"fmt"
)

type Refined[A any, P Predicate] interface {
	Value() A
	_PedicateType() *P // always nil (used to forbid type casting)
}

type refined[A any, P Predicate] struct{ value A }

func (r refined[A, P]) Value() A          { return r.value }
func (r refined[A, P]) _PedicateType() *P { return nil }

type Predicate interface {
	applyPredicate(any) (bool, error)
}

func (rv refined[A, R]) String() string {
	return fmt.Sprint(rv.value)
}

func Refine[T any, P Predicate](v T) (Refined[T, P], error) {
	var predicate P

	res, err := predicate.applyPredicate(v)
	if err != nil {
		return nil, err
	}

	if res {
		return refined[T, P]{v}, nil
	}

	return nil, NonComplianceError[T, P]{v, predicate}
}

func RefineUnsafe[T any, P Predicate](v T) Refined[T, P] {
	res, err := Refine[T, P](v)
	if err != nil {
		panic(err)
	}
	return res
}
