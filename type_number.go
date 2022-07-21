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

type N2[N Number] struct{}

func (i N2[N]) Of() N { return 2 }

type N3[N Number] struct{}

func (i N3[N]) Of() N { return 3 }

type N4[N Number] struct{}

func (i N4[N]) Of() N { return 4 }

type N5[N Number] struct{}

func (i N5[N]) Of() N { return 5 }

type N6[N Number] struct{}

func (i N6[N]) Of() N { return 6 }

type N7[N Number] struct{}

func (i N7[N]) Of() N { return 7 }

type N8[N Number] struct{}

func (i N8[N]) Of() N { return 8 }

type N9[N Number] struct{}

func (i N9[N]) Of() N { return 9 }

type N10[N Number] struct{}

func (i N10[N]) Of() N { return 10 }

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
