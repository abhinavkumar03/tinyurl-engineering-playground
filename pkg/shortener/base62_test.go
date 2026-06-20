package shortener

import "testing"

func TestEncode(t *testing.T) {

	tests := []struct {
		input    int64
		expected string
	}{
		{0, "0"},
		{1, "1"},
		{10, "A"},
		{35, "Z"},
		{36, "a"},
		{61, "z"},
		{62, "10"},
	}

	for _, test := range tests {

		result := Encode(test.input)

		if result != test.expected {
			t.Fatalf(
				"expected %s got %s",
				test.expected,
				result,
			)
		}
	}
}
