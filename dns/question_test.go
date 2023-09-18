package message

import (
	"testing"
)

func TestEncodeName(t *testing.T) {
	testCases := map[string]string{
		"dns.google.com":        "3dns6google3com0",
		"subdomain.example.com": "9subdomain7example3com0",
		"www.openai.org":        "3www6openai3org0",
	}

	for hostname, expected := range testCases {
		res := encodeName(hostname)
		if res != expected {
			t.Errorf("error encoding name. got: %s, expected: %s\n", res, expected)
		}
	}
}
