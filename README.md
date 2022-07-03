# Go Refined

**Go Refined** is a minimalistic type-refinement library.
It is a partial port of Frank Tomas's [Scala refined library](https://github.com/fthomas/refined) and Nikita Volkov's [Haskell refined library](https://hackage.haskell.org/package/refined-0.6.3/docs/Refined.html). It is used to express constraints on type by using type-level predicates.

## Examples
```go
package example
import r "github.com/valentinHenry/refined"

pi, err := r.Refine[int, r.Positive](2)
// pi.Value() == 2, err == nil

pi, err = r.Refine[int, r.Positive](0)
// err == NonComplianceError
// err.Error() == "value 0 does not comply with the predicate refined.Positive"

// Because go does not have type-level numbers yet, it is necessary to implement the TNumber interface 
type N15[N Number] struct{}
func (n N15) Of() N { return 15 }
pj, err = r.Refine[float32, r.LessThan[float32, N15]](12.345)
// pj.Value() == 12.345

pk, err = r.Refine[float32, r.LessThan[uint32, N15]](42)
// err == NonComplianceError
// err.Error() == "value 0 does not comply with the predicate refined.Positive"

str, err := r.Refine[string, r.And[r.NonEmpty, r.Trimmed]]("Hello")
// str.Value() == "Hello"

str, err = r.Refine[string, refined.And[r.NonEmpty, r.Trimmed]]("   hello")
// err == NonComplianceError
// err.Error() == "value 0 does not comply with the predicate ..."

str, err = r.Refine[string, refined.And[r.NonEmpty, r.Trimmed]]("")
// err == NonComplianceError
// err.Error() == "value 0 does not comply with the predicate ..."

func OnSmallSlices(s r.Refined[[]int, r.SizeLessThan[r.N4[int]]]) {
  for _, i2 := range s.Value() {
    fmt.Println(i2)
  }
}

items := []int{1, 2, 3}
ritems, err := r.Refine[[]int, r.SizeLessThan[r.N4[int]]](items)
// ritems.Value() == []int{1, 2, 3}
OnSmallSlices(ritems) // Compiles

items = []int{1, 2, 3, 4}
ritems, err := r.Refine[[]int, r.SizeLessThan[r.N5[int]]](items)
// ritems.Value() == []int{1, 2, 3}
OnSmallSlices(ritems) // Does not compile (bad type)
```

## Provided Predicates
[`logical`](https://github.com/valentinHenry/refined/blob/master/logical.go)
- `And`: the predicate1 AND predicate2
- `Or`: the predicate1 OR predicate`
- `Not`: negate a predicate
- `Xor`: predicate1 XOR predicate2 (p1 && !p2) || (!p1 && p2)

[`string`](https://github.com/valentinHenry/refined/blob/master/strings.go)
- `EndsWith`: checks if a string ends with the suffix S
- `StartsWith`: checks if a string starts with the prefix S
- `Regex`: checks if a string is a valid regex
- `MatchesRegex`: checks if a string matched the regex R
- `ValidInt`: checks if a string is an int ~> int64
- `ValidFloat`: checks if a string is a float32 and/or a float64
- `Trimmed`: checks if a string does not have leading and trailing whitespaces
- `Empty`: checks if a string is empty
- `NonEmpty`: checks if a string is non-empty

[`int` | `int8` | `int32` | `int64` | `float32` | `float64` | `uint` | `uint8` | `uint16` | `uint32` | `uint64`](https://github.com/valentinHenry/refined/blob/master/numbers.go)
- `Positive`: checks if a number is >= 1
- `Negative`: checks if a number is <= -1 (always false for unsigned)
- `NonNegative`: checks if a number is >=0 (always true for unsigned)
- `NonPositive`: checks if a number is <=0 
- `LessThan`: checks if a number is < NB
- `GreaterThan`: checks if a number is > NB
- `From`: checks if a number is >= NB
- `To`: checks if a number is <= NB
- `FromTo`: checks if a number is between `[LNB, RNB]` (LNB <= x <= RNB)
- `Between`: checks if a number is between `]LNB, RNB[` (LNB < x < RNB)
- `FromUntil`: checks if a number is between `[LNB, RNB[` (LNB <= x < RNB)
- `AfterTo`: checks if a number is between `]LNB, RNB]` (LNB < x <= RNB)

[`array`|`slices`|`map`|`List`](https://github.com/valentinHenry/refined/blob/master/collections.go)
- `Empty`: checks if a collection is empty
- `NonEmpty`: checks if a collection is not empty
- `SizeLessThan`: checks if the size of the collection is less(strict) than N
- `SizeGreaterThan`: checks if the size of the collection is greater(strict) than N
- `SizeEqualTo`: checks if the size of the collection is equal to N