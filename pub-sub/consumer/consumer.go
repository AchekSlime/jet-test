package consumer

import (
	"github.com/nats-io/nats.go"
	"log"
)

type Consumer struct {
	NcConn   *nats.Conn
	js       nats.JetStreamContext
	stream   string
	subjects []string
}

func NewConsumer(natsConnection *nats.Conn, jetStream nats.JetStreamContext) *Consumer {
	return &Consumer{
		NcConn: natsConnection,
		js:     jetStream,
	}
}

func (cons *Consumer) Subscribe(subject string) {
	log.Printf("Consumer start listening %s subject", cons.subjects)
	_, err := cons.js.Subscribe(subject, handler)
	if err != nil {
		log.Fatalln("error subscribing: ", err)
	}
}

func handler(msg *nats.Msg) {
	log.Printf("<- from %s, msg: %s", msg.Subject, string(msg.Data))
}
