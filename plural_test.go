package plural

import "testing"

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
