package plural_test

import (
	"fmt"
	"testing"

	"github.com/rickb777/plural/v2"
)

func TestPluralFormatFloatEnglish(t *testing.T) {
	p012 := plural.ByOrdinal("nothing", "%v thing", "%v things")

	cases := []struct {
		n      float64
		expect string
	}{
		{0, "nothing"},
		{0.5, "0.5 thing"},
		{1, "1 thing"},
		{2, "2 things"},
		{3, "3 things"},
		{400, "400 things"},
		{0, "nothing"},
		{1, "1 thing"},
		{2, "2 things"},
		{3, "3 things"},
		{0, "nothing"},
		{1, "1 thing"},
		{2, "2 things"},
		{3, "3 things"},
		{2, "2 things"},
		{0, "nothing"},
		{0.1, "0.1 thing"},
		{2.1, "2.1 things"},
		{3, "3 things"},
		{3.00001, "3.00001 things"},
	}
	for i, c := range cases {
		s := p012.Continuous(c.n)
		if s != c.expect {
			t.Errorf("#%d Format(%g) == %q, want %s", i, c.n, s, c.expect)
		}
	}
}

func TestPluralLengths(t *testing.T) {
	cases := []plural.Cases{
		plural.FromZero("nothing"),
		plural.FromZero("nothing", "one thing"),
		plural.FromZero("nothing", "one thing", "%d things"),
		plural.FromZero("nothing", "one thing", "two things", "%d things"),
	}
	for i, c := range cases {
		s := c.Countable(0)
		if s != "nothing" {
			t.Errorf("#%d Format(0) == %q, want nothing", i, s)
		}
	}
}

func TestByOrdinalFromZero(t *testing.T) {
	p012 := plural.ByOrdinal("nothing", "%d thing", "%d things")

	cases := []struct {
		n      int
		expect string
	}{
		{0, "nothing"},
		{1, "1 thing"},
		{2, "2 things"},
		{3, "3 things"},
		{400, "400 things"},
	}
	for _, c := range cases {
		s := p012.Countable(c.n)
		if s != c.expect {
			t.Errorf("Format(%d) == %s, want %s", c.n, s, c.expect)
		}
	}
}

func TestFromOne(t *testing.T) {
	p12x := plural.FromOne("one thing", "two things", "%d things")

	cases := []struct {
		n      int
		expect string
	}{
		{0, ""},
		{1, "one thing"},
		{2, "two things"},
		{3, "3 things"},
		{400, "400 things"},
	}
	for _, c := range cases {
		s := p12x.Countable(c.n)
		if s != c.expect {
			t.Errorf("Format(%d) == %s, want %s", c.n, s, c.expect)
		}
	}
}

func TestWithoutPlaceholders(t *testing.T) {
	plurals := plural.ByOrdinal("nothing", "one", "some", "many")

	cases := []struct {
		n      float64
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
		s := plurals.Continuous(c.n)
		if s != c.expect {
			t.Errorf("Format(%g) == %s, want %s", c.n, s, c.expect)
		}
	}
}

func ExampleRegular() {
	// Regular caters for the common simple case of a noun that is easy to pluralise, e.g. by appending "s".
	// It uses the Zero and AddS functions, which can be altered if required.

	var catPlurals = plural.Regular("cat")

	for d := 0; d < 4; d++ {
		s := catPlurals.Countable(d)
		fmt.Println(s)
	}

	// Output: no cats
	// 1 cat
	// 2 cats
	// 3 cats
}

func ExampleIrregular() {
	// Irregular is similar to Regular, but differs by catering for the case of nouns that have different
	// forms in the singular and in the plural.

	var lollyPlurals = plural.Irregular("lolly", "lollies")

	for d := 1; d < 4; d++ {
		s := lollyPlurals.Countable(d)
		fmt.Println(s)
	}

	// Output: 1 lolly
	// 2 lollies
	// 3 lollies
}

func ExampleByOrdinal() {
	// ByOrdinal is easy to use for any collection of counting phrases, even with nouns that have irregular
	// plurals. In this case, "lolly" and "lollies" are used.
	//
	// ByOrdinal creates Cases that hold a sequence of cardinal cases where the
	// first matching case is used. Otherwise, if there's no match, the last one is used.

	var lollyPlurals = plural.ByOrdinal("no ice lollies", "1 ice lolly", "%v ice lollies")

	for d := 0; d < 4; d++ {
		s := lollyPlurals.Countable(d)
		fmt.Println(s)
	}

	// Output: no ice lollies
	// 1 ice lolly
	// 2 ice lollies
	// 3 ice lollies
}

func ExampleFromZero() {
	// FromZero creates Cases that hold a sequence of cardinal cases where the
	// first matching case is used. Otherwise, if there's no match, the last one is used.
	//
	// Often, the last case will include a "%d", "%g" or "%v" placeholder for the number,
	// but placeholders are not mandatory in any of the cases.
	//
	// ByOrdinal could be used instead of FromZero (they are aliases);
	// it builds simple common kinds of plurals using small ordinals (0, 1, 2, 3 etc).

	// bikePlurals and weightPlurals provide English formatted cases for some number of bikes and their weight.
	var bikePlurals = plural.FromZero("no bikes weigh", "%d bike weighs", "%d bikes weigh")

	var weightPlurals = plural.FromZero("nothing", "%1.1f tonne", "%1.1f tonnes")

	for d := 0; d < 5; d++ {
		s := bikePlurals.Countable(d)
		w := weightPlurals.Continuous(float64(d) * 0.5)
		fmt.Println(s, w)
	}

	// Output: no bikes weigh nothing
	// 1 bike weighs 0.5 tonne
	// 2 bikes weigh 1.0 tonne
	// 3 bikes weigh 1.5 tonnes
	// 4 bikes weigh 2.0 tonnes
}

func ExampleFromOne() {
	// FromOne creates Cases that hold a sequence of cardinal cases where the
	// first matching case is used. Otherwise, if there's no match, the last one is used.
	//
	// Often, the last case will include a "%d", "%g" or "%v" placeholder for the number,
	// but placeholders are not mandatory in any of the cases.

	// mugPlurals and volumePlurals provide English formatted cases for some number of mugs and their volume.
	var mugPlurals = plural.FromOne("%d mug holds", "%d mugs hold")

	// Note %g, %f etc should be chosen appropriately
	var volumePlurals = plural.FromOne("%g litre", "%g litres")

	for d := 1; d < 6; d++ {
		s := mugPlurals.Countable(d)
		w := volumePlurals.Continuous(float64(d) * 0.25)
		fmt.Println(s, w)
	}

	// Output: 1 mug holds 0.25 litre
	// 2 mugs hold 0.5 litre
	// 3 mugs hold 0.75 litre
	// 4 mugs hold 1 litre
	// 5 mugs hold 1.25 litres
}
