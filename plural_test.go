package plural

import (
	"fmt"
	"testing"
)

func TestPluralFormatIntEnglish(t *testing.T) {
	plurals := Plurals{Case{0, "nothing"}, Case{1, "%v thing"}, Case{2, "%v things"}}

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
		{float32(0.1), "0.1 thing"},
		{float32(2.1), "2.1 things"},
		{float64(3), "3 things"},
		{float64(3.00001), "3.00001 things"},
		{ip(0), "nothing"},
		{ip(1), "1 thing"},
		{ip(2), "2 things"},
		{ip(3), "3 things"},
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

func ip(v int) *int {
	return &v
}

func ExamplePlurals() {
	// Plurals{} holds a sequence of cardinal cases where the first match is used, otherwise the last one is used.
	// The last case will typically include a "%v" placeholder for the number.
	// carPlurals and weightPlurals provide English formatted cases for some number of cars and their weight.
	var carPlurals = Plurals{Case{0, "no cars"}, Case{1, "%v car"}, Case{2, "%v cars"}}
	var weightPlurals = Plurals{Case{0, "weigh nothing"}, Case{1, "weighs %g tonne"}, Case{2, "weigh %1.1f tonnes"}}

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
