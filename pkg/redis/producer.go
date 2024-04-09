package redis

import (
	"context"
	"github.com/mcmacedo/redis-performance-poc/pkg/ports"
	"github.com/redis/go-redis/v9"
	"log"
)

type Producer struct {
	id         string
	client     *redis.Client
	streamName string
}

func (rp *Producer) SendMessage(ctx context.Context, message ports.Message) (string, error) {
	result, err := rp.client.XAdd(ctx, &redis.XAddArgs{
		Stream: rp.streamName,
		Values: []interface{}{"log", message.Data},
	}).Result()

	if err != nil {
		log.Println("Erro ao enviar mensagem:", err)
		return "", err
	}

	log.Printf("Producer id: %s Mensagem enviada com ID: %s", rp.id, result)
	return result, nil
}
