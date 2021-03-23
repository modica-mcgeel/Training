package messages

import (
	"testing"
)

func TestGreet(t *testing.T) {
	t.Parallel()
	got := Greet("Gopher")
	expect := "Hello, Gopher\n"

	if got != expect {
		t.Errorf("Did not get expected result.  Wanted %q, got: %q\n", expect, got)
	}
}

func TestGreetTable(t *testing.T) {
	t.Parallel()
	scenarios := []struct {
		input  string
		expect string
	}{
		{input: "Gopher", expect: "Hello, Gopher\n"},
		{input: "", expect: "Hello, \n"},
	}
	for _, s := range scenarios {
		got := Greet(s.input)
		if got != s.expect {
			//fmt.Errorf("Names do not match")
			t.Errorf("Did not get expect result for input '%v'.  Expected %q, got %q", s.input, got, s.expect)

		}
	}
}

func TestDepart(t *testing.T) {
	got := depart("Gopher")
	expect := "Goodbye, Gopher\n"

	if got != expect {
		t.Errorf("Did not get expected result.  Wanted %q, got: %q\n", expect, got)

	}
}

func TestFailureTypes(t *testing.T) {
	t.SkipNow()
	t.Error("Error signals a failed test, but doesn't stop the rest of the test from executing")
	t.Fatal("Fatal will fail the test and stop it from executing")
	t.Error("This will never print because it is preceeded with a Fatal error")

}
