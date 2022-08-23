package main

import (
	"flag"
	"fmt"
	"jet-test/configs"
	"jet-test/pub-sub/consumer"
	"log"
	"runtime"
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

	sub := consumer.NewConsumer(configs.InitNatsConnections())
	defer sub.NcConn.Close()

	sub.Subscribe(fmt.Sprintf("%s.*", streamName))
	runtime.Goexit()
}
