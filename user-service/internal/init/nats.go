package init

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"music-streaming-microservices/user-service/global"
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
		streamInfo, err := js.StreamInfo(streamConfig.Name)
		if err != nil && err != nats.ErrStreamNotFound {
			log.Fatalf("An unexpected error occurred while checking stream info: %v", err)
		}

		if streamInfo == nil {
			// Create the stream if it does not exist
			log.Printf("Creating stream %s for subject %v", streamConfig.Name, streamConfig.Subjects)
			_, err = js.AddStream(&nats.StreamConfig{
				Name:     streamConfig.Name,
				Subjects: streamConfig.Subjects,
			})
			if err != nil {
				log.Fatalf("Error creating stream %s: %v", streamConfig.Name, err)
			}
			log.Printf("Stream %s created successfully", streamConfig.Name)
		}
	}

	log.Println("Nats JetStream initialized successfully")

	return
}
