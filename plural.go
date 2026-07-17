package plural

import (
	"fmt"
	"strings"
)

// Cases provides a list of plural cases in the order they will be searched.
// The cases must be listed in ascending number order.
// They are continuous (i.e. without gaps) and monotonic (i.e. always rising in ordinal).
type Cases struct {
	indices []int
	labels  string
}

// Zero is a pluggable function that prepends "no " and appends "s" to a noun.
// This is used by [Regular].
// Replace this with whatever your own language needs.
var Zero = func(noun string) string { return "no " + noun + "s" }

// AddS is a pluggable function that appends "s" to a noun. This is used by [Regular].
// Replace this with whatever your own language needs.
var AddS = func(noun string) string { return noun + "s" }

// Regular returns plurals for regular nouns typical in English such as words like "cat", "tree",
// for which the plural form is simply by adding letter 's'.
//
// It uses [FromZero], passing the noun with [Zero] for the zero case, the noun itself for the unit case,
// and with [AddS] for all counts above one. This allows the many easy regular nouns to be pluralised
// with little effort.
//
// For example
//
//	plural.Regular("thing")
//
// can then be used via [Cases.Countable] or [Cases.Continuous] to produce descriptive formatted
// values.
//
// Irregular nouns such as "caddy" should not use this function (unless "caddys" really is the plural
// form you want). Instead, use [FromZero] or [FromOne].
//
// Also, Regular is too simplistic for phrases instead of nouns; for these, [FromZero] or [FromOne]
// will be more appropriate.
func Regular(noun string) Cases {
	return FromZero(Zero(noun), "%v "+noun, "%v "+AddS(noun))
}

// ByOrdinal is an alias for FromZero.
func ByOrdinal(zeroth string, rest ...string) Cases {
	return FromZero(zeroth, rest...)
}

// FromZero constructs a simple set of cases using small ordinals (0, 1, 2, 3 etc), which is a
// common requirement. The last case will be used for all subsequent numbers. Any cases may
// include a fmt.Sprintf number placeholder, Usually, this will be %v, which will support the countable
// and continuous cases effectively.
//
// For example
//
//	plural.FromZero("nothing", "%v caddy", "%v caddies")
//	plural.FromZero("none", "one", "two", "many")
//
// can then be used via [Cases.Countable] or [Cases.Continuous] to produce descriptive formatted
// values.
func FromZero(zeroth string, rest ...string) Cases {
	return newPlurals(false, zeroth, rest...)
}

// FromOne constructs a simple set of cases using small positive numbers (1, 2, 3 etc), which is a
// common requirement. The last case will be used for all subsequent numbers. Any cases may
// include a fmt.Sprintf number placeholder, Usually, this will be %v, which will support the countable
// and continuous cases effectively.
//
// For example
//
//	plural.FromOne("%v cat", "%v cats")
//	plural.FromOne("one", "two", "many")
//
// can then be used via [Cases.Countable] or [Cases.Continuous] to produce descriptive formatted
// values.
//
// Note the behaviour of formatting when the count is zero. As a consequence of [Cases.Countable]
// evaluating its cases in ascending order, FromOne(...).Countable(0) will return a blank string.
// If this might arise, use [FromZero] instead.
func FromOne(first string, rest ...string) Cases {
	return newPlurals(true, first, rest...)
}

func newPlurals(addZero bool, first string, rest ...string) Cases {
	p := Cases{
		indices: make([]int, 0, len(rest)),
	}
	var buf strings.Builder
	buf.WriteString(first)
	i := len(first)
	if addZero {
		p.indices = append(p.indices, 0)
	}
	p.indices = append(p.indices, i)

	switch len(rest) {
	case 0:
	case 1:
		buf.WriteString(rest[0])
	default:
		for _, c := range rest[:len(rest)-1] {
			i += len(c)
			p.indices = append(p.indices, i)
			buf.WriteString(c)
		}
		buf.WriteString(rest[len(rest)-1])
	}

	p.labels = buf.String()
	return p
}

//-------------------------------------------------------------------------------------------------

// Countable expresses a countable number in plural form.
// If no match is found, the last case will be used.
func (plurals Cases) Countable(value int) string {
	prev := 0

	for i, c := range plurals.indices {
		if value <= i {
			message := plurals.labels[prev:c]
			return render(message, value)
		}
		prev = c
	}

	prev = plurals.indices[len(plurals.indices)-1]
	message := plurals.labels[prev:]
	return render(message, value)
}

//-------------------------------------------------------------------------------------------------

// Continuous expresses a continuous number in plural form.
// If no match is found, the last case will be used.
// If the case placeholders contain "%d", this will be replaced by "%g" for floating point formatting.
func (plurals Cases) Continuous(value float64) string {
	prev := 0

	for i, c := range plurals.indices {
		if value <= float64(i) {
			message := plurals.labels[prev:c]
			return render(message, value)
		}
		prev = c
	}

	prev = plurals.indices[len(plurals.indices)-1]
	message := plurals.labels[prev:]
	message = strings.ReplaceAll(message, "%d", "%g")
	return render(message, value)
}

//-------------------------------------------------------------------------------------------------

// render renders a specific case with a given value.
func render[N int | float64](format string, value N) string {
	if strings.IndexByte(format, '%') < 0 {
		return format
	}
	return fmt.Sprintf(format, value)
}
