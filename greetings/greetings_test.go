package greetings

import (
	"regexp"
	"testing"
)

// TESTING FUNCTIONS USING TESTING AND REGEXP

func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := HelloRandom("Gladys")
	if !want.MatchString(msg) || err != nil {
		// t.Error('Hello("Gladys") = %q, %v, want match for %#q, nil', msg, err, want)
		t.Errorf("Hello(%q) = %q, %v; want match for %q, nil", name, msg, err, want.String())

	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := HelloRandom("")
	if msg != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
