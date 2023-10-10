package message

import (
	"testing"
)

/*
split:

32 bytes
0016 0100 0001 0000 0000 0000

my dns name:
3364 6e73 3667 6f6f 676c 6533 636f 6d30

0364 6e73 0667 6f6f 676c 6503 636f 6d00 0001 0001

*/

func TestHexEncodedDNS(t *testing.T) {
	expected := "00160100000100000000000003646e7306676f6f676c6503636f6d0000010001"
	message := Message{
		header: Header{
			ID:                  22,
			QueryOrResponse:     false,
			OpCode:              OpQuery, // Standard Query
			AuthoritativeAnswer: false,
			Truncation:          false,
			RecursionDesired:    true,
			RecursionAvailable:  false,
			ResponseCode:        NoError, // No error
			QDCount:             1,
			ANCount:             0,
			NSCount:             0,
			ARCount:             0,
		},
		questions: []Question{{
			QName:  byteEncodeHostname("dns.google.com"),
			QType:  1,
			QClass: 1,
		}},
	}

	got := message.hexEncodeToString()
	if got != expected {
		t.Errorf("hex encoded:\ngot:      %+v\nexpected: %+v\n", got, expected)
	}
}
