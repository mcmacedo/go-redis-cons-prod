package ports

import "context"

type Message struct {
	Data string
}

// ProducerPort define o contrato para o envio de mensagens.
type ProducerPort interface {
	SendMessage(ctx context.Context, message Message) (string, error)
}

// ConsumerPort define o contrato par o recebimento de mensagens.
type ConsumerPort interface {
	ReceiveMessage(ctx context.Context) (Message, error)
}
