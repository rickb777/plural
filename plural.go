package plural

import (
	"fmt"
	"strings"
)

// Case is the inner element of this API and describes one case. When the number to be described
// matches the number here, the corresponding format string will be used. If the format string
// includes '%', then fmt.Sprintf will be used. Otherwise the format string will be returned verbatim.
type Case struct {
	Number int
	Format string
}

// Plurals provides a list of plural cases in the order they will be searched.
// For plurals of continuous ranges (e.g. weight), the cases must be in ascending number order.
// For plurals of discrete ranges (i.e. integers), the cases can be in any order you require,
// but will conventionally be in ascending number order.
// If no match is found, the last case will be used.
type Plurals []Case

// Format searches through the plural cases for the first match. If none is found, the last
// case is used. The value passed in can be any number type, or pointer to a number type, except
// complex numbers are not supported. The value will be converted to an int in order to
// find the first case that matches.
// The only possible error arises if value has a type that is not numeric.
// It panics if 'plurals' is empty.
func (plurals Plurals) Format(value interface{}) (string, error) {
	switch x := value.(type) {
	case int:
		return plurals.FormatInt(x), nil
	case int8:
		return plurals.FormatInt(int(x)), nil
	case int16:
		return plurals.FormatInt(int(x)), nil
	case int32:
		return plurals.FormatInt(int(x)), nil
	case int64:
		return plurals.FormatInt(int(x)), nil
	case uint8:
		return plurals.FormatInt(int(x)), nil
	case uint16:
		return plurals.FormatInt(int(x)), nil
	case uint32:
		return plurals.FormatInt(int(x)), nil
	case uint64:
		return plurals.FormatInt(int(x)), nil
	case float32:
		return plurals.FormatFloat(x), nil
	case float64:
		return plurals.FormatFloat(float32(x)), nil

	case *int:
		return plurals.FormatInt(*x), nil
	case *int8:
		return plurals.FormatInt(int(*x)), nil
	case *int16:
		return plurals.FormatInt(int(*x)), nil
	case *int32:
		return plurals.FormatInt(int(*x)), nil
	case *int64:
		return plurals.FormatInt(int(*x)), nil
	case *uint:
		return plurals.FormatInt(int(*x)), nil
	case *uint8:
		return plurals.FormatInt(int(*x)), nil
	case *uint16:
		return plurals.FormatInt(int(*x)), nil
	case *uint32:
		return plurals.FormatInt(int(*x)), nil
	case *uint64:
		return plurals.FormatInt(int(*x)), nil
	case *float32:
		return plurals.FormatFloat(*x), nil
	case *float64:
		return plurals.FormatFloat(float32(*x)), nil

	default:
		return "", fmt.Errorf("Unexpected type %T for %v", x, value)
	}
}

// FormatInt expresses an int in plural form. It panics if 'plurals' is empty.
func (plurals Plurals) FormatInt(value int) string {
	for _, p := range plurals {
		if value == p.Number {
			if strings.IndexByte(p.Format, '%') < 0 {
				return p.Format
			}
			return fmt.Sprintf(p.Format, value)
		}
	}
	p := plurals[len(plurals)-1]
	return fmt.Sprintf(p.Format, value)
}

// FormatFloat expresses a float32 in plural form. It panics if 'plurals' is empty.
func (plurals Plurals) FormatFloat(value float32) string {
	for _, p := range plurals {
		if value <= float32(p.Number) {
			if strings.IndexByte(p.Format, '%') < 0 {
				return p.Format
			}
			return fmt.Sprintf(p.Format, value)
		}
	}
	p := plurals[len(plurals)-1]
	return fmt.Sprintf(p.Format, value)
}
