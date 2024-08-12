package sdk

import "testing"

func TestFormatCVE(t *testing.T) {
	t.Run("cve is formatted correctly", func(t *testing.T) {
		tests := []struct {
			name     string
			cve      string
			expected string
		}{
			{
				name:     "cve is already formatted correctly",
				cve:      "CVE-2024-38077",
				expected: "CVE-2024-38077",
			},
			{
				name:     "cve is formatted correctly with lowercase",
				cve:      "cve-2024-38077",
				expected: "CVE-2024-38077",
			},
			{
				name:     "cve is formatted correctly with mixed case",
				cve:      "cVe-2024-38077",
				expected: "CVE-2024-38077",
			},
			{
				name:     "cve is formatted correctly without prefix",
				cve:      "2024-38077",
				expected: "CVE-2024-38077",
			},
			{
				name:     "cve is formatted correctly if underscores are used instead of dashes",
				cve:      "CVE_2024_38077",
				expected: "CVE-2024-38077",
			},
			{
				name:     "cve is formatted correctly if special characters are used",
				cve:      "CVE-2024-38077!",
				expected: "CVE-2024-38077",
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				formattedCve := FormatCVE(test.cve)
				if formattedCve != test.expected {
					t.Errorf("expected %s, got %s", test.expected, formattedCve)
				}
			})
		}

	})
}
