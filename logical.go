package refined

type And[P1, P2 Predicate] struct {
	_P1 *P1
	_P2 *P2
}
type Or[P1, P2 Predicate] struct {
	_P1 *P1
	_P2 *P2
}
type Not[P1 Predicate] struct {
	_P1 *P1
}
type Xor[P1, P2 Predicate] struct {
	_P1 *P1
	_P2 *P2
}

func (a And[P1, P2]) applyPredicate(v any) (bool, error) {
	var p1 P1
	v1, err1 := p1.applyPredicate(v)
	if err1 != nil {
		return false, err1
	}

	var p2 P2
	v2, err2 := p2.applyPredicate(v)
	if err2 != nil {
		return false, err2
	}

	return v1 && v2, nil
}
func (o Or[P1, P2]) applyPredicate(v any) (bool, error) {
	var p1 P1
	v1, err1 := p1.applyPredicate(v)
	if err1 != nil {
		return false, err1
	}

	var p2 P2
	v2, err2 := p2.applyPredicate(v)
	if err2 != nil {
		return false, err2
	}

	return v1 || v2, nil
}
func (n Not[P1]) applyPredicate(v any) (bool, error) {
	var p1 P1
	res, err := p1.applyPredicate(v)
	return !res, err
}
func (x Xor[P1, P2]) applyPredicate(v any) (bool, error) {
	var p1 P1
	v1, err1 := p1.applyPredicate(v)
	if err1 != nil {
		return false, err1
	}

	var p2 P2
	v2, err2 := p2.applyPredicate(v)
	if err2 != nil {
		return false, err2
	}

	return v1 != v2, nil
}
