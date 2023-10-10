package message

import "encoding/hex"

type Message struct {
	header     Header
	questions  []Question
	answer     Answer
	authority  Authority
	additional Additional
}

func (m Message) byteEncode() (encoded []byte) {
	messageEncoded := append(m.header.byteEncode())
	for _, q := range m.questions {
		messageEncoded = append(messageEncoded, q.byteEncode()...)
	}
	return messageEncoded
}

func (m Message) hexEncodeToString() string {
	return hex.EncodeToString(m.byteEncode())
}

type Answer struct{}

type Authority struct{}

type Additional struct{}
