package configs

import (
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

const (
	JsUrl = "http://localhost:4222"
)

func InitNatsConnections() (*nats.Conn, nats.JetStreamContext) {
	nc, err := nats.Connect(JsUrl, nats.MaxReconnects(5), nats.ReconnectWait(20*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	return nc, js
}
