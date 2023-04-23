package time

import "testing"

func TestParseTime(t *testing.T) {
	tests := []struct {
		desc     string
		input    string
		expected string
	}{
		{
			desc:     "success - time parsed successfuly",
			input:    "11111",
			expected: "0 day(s), 3 hour(s), 5 minute(s)",
		},
		{
			desc:     "failure - input wrong format",
			input:    "a",
			expected: "Unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			parsedTime := ParseUptime(tt.input)
			if parsedTime != tt.expected {
				t.Fatalf("received %s but expected %s", parsedTime, tt.expected)
			}
		})
	}
}
