package module1

import "testing"

func Test_factorial(t *testing.T) {
	actual := Factorial(5)
	expected := 120
	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}
