package message

import (
	"encoding/binary"
	"math/rand"
)

/*
defined in RFC 1035 section 4.1.1
reference: https://datatracker.ietf.org/doc/html/rfc1035#section-4.1.1

	 ID

		 A 16 bit identifier assigned by the program that
		 generates any kind of query.  This identifier is copied
			the corresponding reply and can be used by the requester
			to match up replies to outstanding queries.

	 QR

	 A one bit field that specifies whether this message is a
		query (0), or a response (1).

	 OPCODE

	 A four bit field that specifies kind of query in this
		message.  This value is set by the originator of a query
		and copied into the response.  The values are:

	 0               a standard query (QUERY)

	 1               an inverse query (IQUERY)

	 2               a server status request (STATUS)

	 3-15            reserved for future use

	 AA

	 Authoritative Answer - this bit is valid in responses,
		and specifies that the responding name server is an
		authority for the domain name in question section.

		Note that the contents of the answer section may have
		multiple owner names because of aliases.  The AA bit
		corresponds to the name which matches the query name, or
		the first owner name in the answer section.

	 TC

	 TrunCation - specifies that this message was truncated
		due to length greater than that permitted on the
		transmission channel.

	 RD

	 Recursion Desired - this bit may be set in a query and
		is copied into the response.  If RD is set, it directs
		the name server to pursue the query recursively.
		Recursive query support is optional.

	 RA

	 Recursion Available - this be is set or cleared in a
		response, and denotes whether recursive query support is
		available in the name server.

	 Z

	 Reserved for future use.  Must be zero in all queries
		and responses.

	 RCODE

	 Response code - this 4 bit field is set as part of
		responses.  The values have the following
		interpretation:

		0               No error condition

		1               Format error - The name server was
		                unable to interpret the query.

		2               Server failure - The name server was
		                unable to process this query due to a
		                problem with the name server.

		3               Name Error - Meaningful only for
		                responses from an authoritative name
		                server, this code signifies that the
		                domain name referenced in the query does
		                not exist.

		4               Not Implemented - The name server does
		                not support the requested kind of query.

		5               Refused - The name server refuses to
		                perform the specified operation for
		                policy reasons.  For example, a name
		                server may not wish to provide the
		                information to the particular requester,
		                or a name server may not wish to perform
		                a particular operation (e.g., zone
		                transfer) for particular data.

		6-15            Reserved for future use.

	 QDCOUNT

	 an unsigned 16 bit integer specifying the number of
		entries in the question section.

	 ANCOUNT

	 an unsigned 16 bit integer specifying the number of
		resource records in the answer section.

	 NSCOUNT

	 an unsigned 16 bit integer specifying the number of name
		server resource records in the authority records
		section.

	 ARCOUNT

	 an unsigned 16 bit integer specifying the number of
		resource records in the additional records section.
*/
type Header struct {
	ID                  uint16
	QueryOrResponse     bool // query (0) response (1)
	OpCode              uint16
	AuthoritativeAnswer bool
	Truncation          bool
	RecursionDesired    bool
	RecursionAvailable  bool
	// Z
	ResponseCode uint16

	QDCount uint16
	ANCount uint16
	NSCount uint16
	ARCount uint16
}

// op code values
const (
	OpQuery = iota
	OpInverseQuery
	OpStatus
)

// response code values
const (
	NoError = iota
	FormatError
	ServerFailure
	NameError
	NotImplemented
	Refused
)

func NewHeader() Header {
	return Header{
		ID: generateUniqueID(),
	}
}

func (h Header) byteEncode() []byte {
	buf := make([]byte, 12)

	binary.BigEndian.PutUint16(buf[0:2], h.ID)

	flags := uint16(0)
	if h.QueryOrResponse {
		flags |= 1 << 15
	}
	flags |= h.OpCode << 11
	if h.AuthoritativeAnswer {
		flags |= 1 << 10
	}
	if h.Truncation {
		flags |= 1 << 9
	}
	if h.RecursionDesired {
		flags |= 1 << 8
	}
	if h.RecursionAvailable {
		flags |= 1 << 7
	}
	flags |= h.ResponseCode << 4

	binary.BigEndian.PutUint16(buf[2:4], flags)
	binary.BigEndian.PutUint16(buf[4:6], h.QDCount)
	binary.BigEndian.PutUint16(buf[6:8], h.ANCount)
	binary.BigEndian.PutUint16(buf[8:10], h.NSCount)
	binary.BigEndian.PutUint16(buf[10:], h.ARCount)

	return buf
}

func generateUniqueID() uint16 {
	return uint16(rand.Intn(1 << 16))
}
