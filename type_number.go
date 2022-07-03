package refined

type Number interface {
	int | int8 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64
}

type TNumber[N Number] interface {
	Of() N
}

type N0[N Number] struct{}

func (i N0[N]) Of() N { return 0 }

type N1[N Number] struct{}

func (i N1[N]) Of() N { return 1 }

// Create the one you need

func applyOn[N Number, P Predicate](left N, right any, fn func(N, N) bool, p P) (bool, error) {
	switch r := right.(type) {
	case N:
		return fn(left, r), nil
	default:
		return false, valueTypeError[P](right, p)
	}
}

func valueTypeError[P Predicate](value any, p P) error {
	return IllegalValueTypeError[P]{value, p}
}
