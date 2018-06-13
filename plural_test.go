package plural

import (
	"fmt"
	"testing"
)

func TestPluralFormatIntEnglish(t *testing.T) {
	p012 := Plurals{Case{0, "nothing"}, Case{1, "%v thing"}, Case{2, "%v things"}}

	cases := []struct {
		n      interface{}
		expect string
	}{
		{0, "nothing"},
		{1, "1 thing"},
		{2, "2 things"},
		{3, "3 things"},
		{400, "400 things"},
		{int8(0), "nothing"},
		{int16(1), "1 thing"},
		{int32(2), "2 things"},
		{int64(3), "3 things"},
		{uint8(0), "nothing"},
		{uint16(1), "1 thing"},
		{uint32(2), "2 things"},
		{uint64(3), "3 things"},
		{float32(2), "2 things"},
		{float32(0), "nothing"},
		{float32(0.1), "0.1 thing"},
		{float32(2.1), "2.1 things"},
		{float64(3), "3 things"},
		{float64(3.00001), "3.00001 things"},
		{fp32(3), "3 things"},
		{fp64(3), "3 things"},
		{ip(0), "nothing"},
		{ip(1), "1 thing"},
		{ip(2), "2 things"},
		{ip(3), "3 things"},
		{ip8(3), "3 things"},
		{ip16(3), "3 things"},
		{ip32(3), "3 things"},
		{ip64(3), "3 things"},
		{uip(3), "3 things"},
		{uip8(3), "3 things"},
		{uip16(3), "3 things"},
		{uip32(3), "3 things"},
		{uip64(3), "3 things"},
	}
	for _, c := range cases {
		s, err := p012.Format(c.n)
		if err != nil {
			t.Errorf("Format(%d) => %v, want %s", c.n, err, c.expect)
		} else if s != c.expect {
			t.Errorf("Format(%d) == %s, want %s", c.n, s, c.expect)
		}
	}
}

func TestSimplePlurals(t *testing.T) {
	p012 := ByOrdinal("nothing", "%v thing", "%v things")

	cases := []struct {
		n      interface{}
		expect string
	}{
		{0, "nothing"},
		{1, "1 thing"},
		{2, "2 things"},
		{3, "3 things"},
		{400, "400 things"},
	}
	for _, c := range cases {
		s, err := p012.Format(c.n)
		if err != nil {
			t.Errorf("Format(%d) => %v, want %s", c.n, err, c.expect)
		} else if s != c.expect {
			t.Errorf("Format(%d) == %s, want %s", c.n, s, c.expect)
		}
	}
}

func TestWithoutPlaceholders(t *testing.T) {
	plurals := ByOrdinal("nothing", "one", "some", "many")

	cases := []struct {
		n      interface{}
		expect string
	}{
		{0, "nothing"},
		{1, "one"},
		{2, "some"},
		{3, "many"},
		{400, "many"},
		{4.1, "many"},
	}
	for _, c := range cases {
		s, err := plurals.Format(c.n)
		if err != nil {
			t.Errorf("Format(%d) => %v, want %s", c.n, err, c.expect)
		} else if s != c.expect {
			t.Errorf("Format(%d) == %s, want %s", c.n, s, c.expect)
		}
	}
}

func TestErrorCase(t *testing.T) {
	plurals := Plurals{Case{0, "nothing"}, Case{1, "%v thing"}, Case{2, "%v things"}}

	cases := []struct {
		n      interface{}
		expect string
	}{
		{"foo", "Unexpected type string for foo"},
		{nil, `Unexpected nil value for Plurals({0 -> "nothing"}, {1 -> "%v thing"}, {2 -> "%v things"})`},
	}
	for _, c := range cases {
		_, err := plurals.Format(c.n)
		if err == nil {
			t.Errorf("Format(%#v) no error, want %s", c.n, c.expect)
		} else if err.Error() != c.expect {
			t.Errorf("Format(%v) == %s, want %s", c.n, err.Error(), c.expect)
		}
	}
}

func ip(v int) *int {
	return &v
}

func ip8(v int8) *int8 {
	return &v
}

func ip16(v int16) *int16 {
	return &v
}

func ip32(v int32) *int32 {
	return &v
}

func ip64(v int64) *int64 {
	return &v
}

func uip(v uint) *uint {
	return &v
}

func uip8(v uint8) *uint8 {
	return &v
}

func uip16(v uint16) *uint16 {
	return &v
}

func uip32(v uint32) *uint32 {
	return &v
}

func uip64(v uint64) *uint64 {
	return &v
}

func fp32(v float32) *float32 {
	return &v
}

func fp64(v float64) *float64 {
	return &v
}

func ExamplePlurals() {
	// Plurals{} holds a sequence of cardinal cases where the first match is used, otherwise the last one is used.
	// The last case will typically include a "%v" placeholder for the number.
	// carPlurals and weightPlurals provide English formatted cases for some number of cars and their weight.
	var carPlurals = Plurals{
		Case{0, "no cars weigh"},
		Case{1, "%v car weighs"},
		Case{2, "%v cars weigh"},
	}
	var weightPlurals = Plurals{
		Case{0, "nothing"},
		Case{1, "%1.1f tonne"},
		Case{2, "%1.1f tonnes"},
	}

	for d := 0; d < 4; d++ {
		s, _ := carPlurals.Format(d)
		w, _ := weightPlurals.Format(float32(d) * 0.6)
		fmt.Printf("%s %s\n", s, w)
	}

	// Output: no cars weigh nothing
	// 1 car weighs 0.6 tonne
	// 2 cars weigh 1.2 tonnes
	// 3 cars weigh 1.8 tonnes
}

func ExampleByOrdinal() {
	// ByOrdinal(...) builds simple common kinds of plurals using small ordinals (0, 1, 2, 3 etc).
	// Notice that the counting starts from zero.
	var carPlurals = ByOrdinal("no cars weigh", "%v car weighs", "%v cars weigh")

	// Note %g, %f etc should be chosen appropriately; both are used here for illustration
	var weightPlurals = ByOrdinal("nothing", "%g tonne", "%1.1f tonnes")

	for d := 0; d < 5; d++ {
		s, _ := carPlurals.Format(d)
		w, _ := weightPlurals.Format(float32(d) * 0.5)
		fmt.Printf("%s %s\n", s, w)
	}

	// Output: no cars weigh nothing
	// 1 car weighs 0.5 tonne
	// 2 cars weigh 1 tonne
	// 3 cars weigh 1.5 tonnes
	// 4 cars weigh 2.0 tonnes
}
