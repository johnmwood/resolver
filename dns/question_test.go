package message

import (
	"encoding/hex"
	"testing"
)

func TestEncodeHostnameHexValue(t *testing.T) {
	testCases := map[string]string{
		"dns.google.com":        "03646e7306676f6f676c6503636f6d00",
		"subdomain.example.com": "09737562646f6d61696e076578616d706c6503636f6d00",
		"www.openai.org":        "03777777066f70656e6169036f726700",
	}

	for hostname, expected := range testCases {
		res := hex.EncodeToString(byteEncodeHostname(hostname))
		if res != expected {
			t.Errorf("error encoding name: %s\ngot:      %s\nexpected: %s\n\n", hostname, res, expected)
		}
	}
}
