package message

type Message struct {
	/*
		- header
		- questions
		- answer
		- authority
		- additional
	*/
}

/*
	defined in RFC 1035 section 4.1.1
    reference: https://datatracker.ietf.org/doc/html/rfc1035#section-4.1.1

*/
type Header struct{}

type Questions struct{}

type Answer struct{}

type Authority struct{}

type Additional struct{}
