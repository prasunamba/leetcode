package module1

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("expected %d ,got %d", expected, result)
	}
}

// Always ends with _test.go
// Test function must start with Test

// Test function must accept *testing.T as a parameter
