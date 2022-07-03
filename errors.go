package refined

import (
	"fmt"
	"reflect"
)

type NonComplianceError[A any, P Predicate] struct {
	value     A
	predicate P
}

func (nc NonComplianceError[A, P]) Error() string {
	return fmt.Sprint("value ", nc.value, " does not comply with the predicate ", reflect.TypeOf(nc.predicate))
}

type IllegalValueTypeError[P Predicate] struct {
	value     any
	predicate P
}

func (ivte IllegalValueTypeError[P]) Error() string {
	return fmt.Sprint("value ", ivte.value, " is not of the same type as predicate ", reflect.TypeOf(ivte.predicate))
}
