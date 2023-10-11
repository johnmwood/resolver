package message

type MessageOption func(*Message)

func WithDefaultHeader() MessageOption {
	return func(m *Message) {
		m.header = Header{
			ID:                  generateUniqueID(),
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
		}
	}
}

func WithHeaderID(id uint16) MessageOption {
	return func(m *Message) {
		m.header.ID = id
	}
}

func WithQuestion(hostname string, qtype, qclass uint16) MessageOption {
	return func(m *Message) {
		m.questions = append(m.questions, Question{
			QName:  byteEncodeHostname(hostname),
			QType:  qtype,
			QClass: qclass,
		})
	}
}
