package main

import (
	"flag"
	"jet-test/configs"
	"jet-test/pub-sub/publisher"
	"log"
)

func main() {
	log.Println("Welcome to NATS JetStream Publisher...")

	// Parsing args
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		log.Fatalln("missing stream name")
	}
	streamName := args[0]

	pub := publisher.NewPublisher(configs.InitNatsConnections())
	defer pub.NcConn.Close()

	pub.CreateStream(streamName)
	defer pub.CleanUp(streamName)

	pub.StartPublishing(100, 1)
}
