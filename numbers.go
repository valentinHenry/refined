package refined

type PosInt = Refined[int, Positive]
type NonNegInt = Refined[int, NonNegative]
type NegInt = Refined[int, Negative]
type NonPosInt = Refined[int, NonPositive]

type PosFloat32 = Refined[float32, Positive]
type NonNegFloat32 = Refined[float32, NonNegative]
type NegFloat32 = Refined[float32, Negative]
type NonPosFloat32 = Refined[float32, NonPositive]

type Positive struct{}                                            // x >0
type Negative struct{}                                            // x <0
type NonNegative struct{}                                         // x >= 0
type NonPositive struct{}                                         // x <= 0
type LessThan[N Number, NB TNumber[N]] struct{}                   // x < NB
type GreaterThan[N Number, NB TNumber[N]] struct{}                // x > NB
type From[N Number, NB TNumber[N]] struct{}                       // x >= NB
type To[N Number, NB TNumber[N]] struct{}                         // x <= N
type FromTo[N Number, LNB TNumber[N], RNB TNumber[N]] struct{}    // LNB <= x <= RNB
type Between[N Number, LNB TNumber[N], RNB TNumber[N]] struct{}   // LNB < x < RNB
type FromUntil[N Number, LNB TNumber[N], RNB TNumber[N]] struct{} // LNB <= x < RNB
type AfterTo[N Number, LNB TNumber[N], RNB TNumber[N]] struct{}   // LNB < x <= RNB

func (p Positive) applyPredicate(n any) (bool, error) {
	switch n := n.(type) {
	case int:
		return n > 0, nil
	case int8:
		return n > 0, nil
	case int32:
		return n > 0, nil
	case int64:
		return n > 0, nil
	case float32:
		return n > 0, nil
	case float64:
		return n > 0, nil
	case uint:
		return n > 0, nil
	case uint8:
		return n > 0, nil
	case uint16:
		return n > 0, nil
	case uint32:
		return n > 0, nil
	case uint64:
		return n > 0, nil
	default:
		return false, valueTypeError(n, p)
	}
}
func (p Negative) applyPredicate(n any) (bool, error) {
	switch n := n.(type) {
	case int:
		return n < 0, nil
	case int8:
		return n < 0, nil
	case int32:
		return n < 0, nil
	case int64:
		return n < 0, nil
	case float32:
		return n < 0, nil
	case float64:
		return n < 0, nil
	case uint:
		return n < 0, nil
	case uint8:
		return n < 0, nil
	case uint16:
		return n < 0, nil
	case uint32:
		return n < 0, nil
	case uint64:
		return n < 0, nil
	default:
		return false, valueTypeError(n, p)
	}
}
func (p NonNegative) applyPredicate(n any) (bool, error) {
	switch n := n.(type) {
	case int:
		return n >= 0, nil
	case int8:
		return n >= 0, nil
	case int32:
		return n >= 0, nil
	case int64:
		return n >= 0, nil
	case float32:
		return n >= 0, nil
	case float64:
		return n >= 0, nil
	case uint:
		return n >= 0, nil
	case uint8:
		return true, nil
	case uint16:
		return true, nil
	case uint32:
		return true, nil
	case uint64:
		return n >= 0, nil
	default:
		return false, valueTypeError(n, p)
	}
}
func (p NonPositive) applyPredicate(n any) (bool, error) {
	switch n := n.(type) {
	case int:
		return n <= 0, nil
	case int8:
		return n <= 0, nil
	case int32:
		return n <= 0, nil
	case int64:
		return n <= 0, nil
	case float32:
		return n <= 0, nil
	case float64:
		return n <= 0, nil
	case uint:
		return n <= 0, nil
	case uint8:
		return n <= 0, nil
	case uint16:
		return n <= 0, nil
	case uint32:
		return n <= 0, nil
	case uint64:
		return n <= 0, nil
	default:
		return false, valueTypeError(n, p)
	}
}
func (p LessThan[N, NB]) applyPredicate(n any) (bool, error) {
	var nb NB
	return applyOn(nb.Of(), n, func(l N, r N) bool { return l > r }, p)
}
func (p GreaterThan[N, NB]) applyPredicate(n any) (bool, error) {
	var nb NB
	return applyOn(nb.Of(), n, func(l N, r N) bool { return l < r }, p)
}
func (p From[N, NB]) applyPredicate(n any) (bool, error) {
	var nb NB
	return applyOn(nb.Of(), n, func(l N, r N) bool { return l <= r }, p)
}
func (p To[N, NB]) applyPredicate(n any) (bool, error) {
	var nb NB
	return applyOn(nb.Of(), n, func(l N, r N) bool { return l >= r }, p)
}
func (ft Between[N, LNB, RNB]) applyPredicate(n any) (bool, error) {
	var pred And[GreaterThan[N, LNB], LessThan[N, RNB]]
	return pred.applyPredicate(n)
}
func (ft FromTo[N, LNB, RNB]) applyPredicate(n any) (bool, error) {
	var pred And[From[N, LNB], To[N, RNB]]
	return pred.applyPredicate(n)
}
func (ft FromUntil[N, LNB, RNB]) applyPredicate(n any) (bool, error) {
	var pred And[From[N, LNB], LessThan[N, RNB]]
	return pred.applyPredicate(n)
}
func (ft AfterTo[N, LNB, RNB]) applyPredicate(n any) (bool, error) {
	var pred And[GreaterThan[N, LNB], To[N, RNB]]
	return pred.applyPredicate(n)
}
