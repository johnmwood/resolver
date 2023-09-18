package message

type Message struct {
	header     Header
	questions  []Question
	answer     Answer
	authority  Authority
	additional Additional
}

type Answer struct{}

type Authority struct{}

type Additional struct{}
