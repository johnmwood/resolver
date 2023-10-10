package message

import (
	"bytes"
	"encoding/binary"
	"strings"
)

// QNAME
//
//	 a domain name represented as a sequence of labels, where
//		each label consists of a length octet followed by that
//		number of octets.  The domain name terminates with the
//		zero length octet for the null label of the root.  Note
//		that this field may be an odd number of octets; no
//		padding is used.
//
// QTYPE
//
//	 a two octet code which specifies the type of the query.
//	 The values for this field include all codes valid for a
//		TYPE field, together with some more general codes which
//		can match more than one type of RR.
//
// QCLASS
//
//	a two octet code that specifies the class of the query.
//	For example, the QCLASS field is IN for the Internet.
type Question struct {
	QName  []byte
	QType  uint16
	QClass uint16
}

func (q Question) byteEncode() []byte {
	buf := new(bytes.Buffer)

	binary.Write(buf, binary.BigEndian, q.QName)
	binary.Write(buf, binary.BigEndian, q.QType)
	binary.Write(buf, binary.BigEndian, q.QClass)

	return buf.Bytes()
}

// performs dns name compression
// e.g. dns.google.com -> 3dns6google3com0
func byteEncodeHostname(hostname string) []byte {
	encodedName := []byte{}
	labels := strings.Split(hostname, ".")

	for _, label := range labels {
		encodedName = append(encodedName, byte(len(label)))
		for _, r := range label {
			encodedName = append(encodedName, byte(r))
		}
	}
	encodedName = append(encodedName, byte(0))

	return encodedName
}
