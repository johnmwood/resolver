package message

import (
	"testing"
)

func TestHexEncodedDNS(t *testing.T) {
	expected := "00160100000100000000000003646e7306676f6f676c6503636f6d0000010001"
	message := NewMessage(
		WithDefaultHeader(),
		WithHeaderID(22),
		WithQuestion("dns.google.com", 1, 1),
	)

	got := message.hexEncodeToString()
	if got != expected {
		t.Errorf("hex encoded:\ngot:      %+v\nexpected: %+v\n", got, expected)
	}
}
