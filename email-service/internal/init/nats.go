package init

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"music-streaming-microservices/email-service/global"
	"os"
)

func InitNats() {
	connString := fmt.Sprintf("nats://%s:%s", global.Configs.Nats.Host, global.Configs.Nats.Port)
	nc, err := nats.Connect(connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to Nats Jetstream: %v\n", err)
		os.Exit(1)
	}

	global.NatsConn = nc

	js, err := nc.JetStream()
	if err != nil {
		log.Fatalf("Error creating JetStream context: %v", err)
	}

	global.NatsJetStream = js

	for _, streamConfig := range global.Configs.Nats.Streams {
		stream, err := js.StreamInfo(streamConfig.Name)
		if err != nil {
			log.Fatalf("Error when checking Jetstreams: %v", err)
		}
		log.Printf("Found subjects::: %s \n", stream.Config.Subjects)
	}

	log.Println("Nats JetStream initialized successfully")

	return
}
