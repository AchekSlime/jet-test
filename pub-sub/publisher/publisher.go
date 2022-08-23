package publisher

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

type Publisher struct {
	NcConn   *nats.Conn
	js       nats.JetStreamContext
	stream   string
	subjects []string
}

func NewPublisher(natsConnection *nats.Conn, jetStream nats.JetStreamContext) *Publisher {
	return &Publisher{

		NcConn: natsConnection,
		js:     jetStream,
	}
}

func (pub *Publisher) CreateStream(streamName string) {
	pub.stream = streamName
	pub.subjects = append(pub.subjects, streamName+".*")

	if _, err := pub.js.AddStream(&nats.StreamConfig{
		Name:     pub.stream,
		Subjects: pub.subjects,
	}); err != nil {
		log.Fatalln("error creating stream: ", err)
	}
}

func (pub *Publisher) CleanUp(streamName string) {
	if err := pub.js.DeleteStream(streamName); err != nil {
		log.Fatalln("error deleting stream: ", err)
	}
}

func (pub *Publisher) StartPublishing(msgCount int, secondsOnWait int) {
	log.Printf("Publishing on %s", pub.subjects[0])
	for i := 0; i < msgCount; i++ {
		msg := fmt.Sprintf("%s-[%d]- %s", "lol", i, "lol")
		log.Println(fmt.Sprintf("[%d] -> msg: %s", i, msg))

		if _, err := pub.js.Publish(pub.subjects[0], []byte(msg)); err != nil {
			log.Fatalln("error publishing: ", err)
		}
		time.Sleep(time.Duration(secondsOnWait) * time.Second)
	}
}
