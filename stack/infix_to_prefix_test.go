package stack_test

import (
	"testing"

	s "github.com/rohanshrestha09/dsa-go/stack"
)

func TestInfixToPrefix(t *testing.T) {
	expected := "abcde*f-/g*-*h^"

	result := s.InfixToPostfix("(a*(b-c/(d*e-f)*g))^h")

	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}
