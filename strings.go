package refined

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type NonEmptyString = Refined[string, NonEmpty]
type NonBlankString = Refined[string, And[NonEmpty, Trimmed]]

type EndsWith[S Substring] struct{}
type StartsWith[S Substring] struct{}
type Regex struct{}
type MatchesRegex[R RegexExpression] struct{}
type ValidInt struct{}
type ValidFloat struct{}
type Trimmed struct{}

type RegexExpression interface {
	Expression() *regexp.Regexp
}
type Substring interface {
	Substring() string
}

func (e EndsWith[S]) applyPredicate(a any) (bool, error) {
	var ss S
	return applyFn2OnString(a, ss.Substring(), strings.HasSuffix, e)
}
func (s StartsWith[S]) applyPredicate(a any) (bool, error) {
	var ss S
	return applyFn2OnString(a, ss.Substring(), strings.HasPrefix, s)
}
func (r Regex) applyPredicate(a any) (bool, error) {
	return applyFn1OnString(
		a,
		func(s string) bool {
			_, err := regexp.Compile(s)
			return err == nil
		},
		r,
	)
}
func (s MatchesRegex[R]) applyPredicate(a any) (bool, error) {
	var regex R
	return applyFn1OnString(a, regex.Expression().MatchString, s)
}
func (v ValidInt) applyPredicate(a any) (bool, error) {
	return applyFn1OnString(
		a,
		func(str string) bool {
			_, err := strconv.ParseInt(str, 10, 64)
			return err == nil
		},
		v,
	)
}
func (v ValidFloat) applyPredicate(a any) (bool, error) {
	return applyFn1OnString(
		a,
		func(str string) bool {
			_, err := strconv.ParseFloat(str, 64)
			return err == nil
		},
		v,
	)
}
func (t Trimmed) applyPredicate(a any) (bool, error) {
	return applyFn1OnString(
		a,
		func(str string) bool {
			if len(str) == 0 {
				return true
			}
			r := []rune(str)
			return !(unicode.IsSpace(r[0]) || unicode.IsSpace(r[len(r)-1]))
		},
		t,
	)
}

func applyFn1OnString(left any, fn func(string) bool, p Predicate) (bool, error) {
	switch a := left.(type) {
	case string:
		return fn(a), nil
	default:
		return false, valueTypeError(a, p)
	}
}
func applyFn2OnString(left any, right string, fn func(string, string) bool, p Predicate) (bool, error) {
	switch a := left.(type) {
	case string:
		return fn(a, right), nil
	default:
		return false, valueTypeError(a, p)
	}
}
