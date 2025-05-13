package module1

import "testing"

func TestEvenorOdd(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{0, true},
		{1, false},
		{3, false},
		{4, true},
		{7, false},
	}
	for _, value := range tests {
		result := EvenorOdd(value.input)
		if result != value.expected {
			t.Errorf("expected %v got %v", value.expected, result)
		}
	}
}
