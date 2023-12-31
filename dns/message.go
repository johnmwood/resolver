package message

import "encoding/hex"

type Message struct {
	header     Header
	questions  []Question
	answer     Answer
	authority  Authority
	additional Additional
}

func NewMessage(opts ...MessageOption) Message {
	m := Message{}
	for _, opt := range opts {
		opt(&m)
	}
	return m
}

func (m Message) ByteEncode() (encoded []byte) {
	messageEncoded := append(m.header.byteEncode())
	for _, q := range m.questions {
		messageEncoded = append(messageEncoded, q.byteEncode()...)
	}
	return messageEncoded
}

func (m Message) hexEncodeToString() string {
	return hex.EncodeToString(m.ByteEncode())
}

type Answer struct{}

type Authority struct{}

type Additional struct{}
