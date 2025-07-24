package consumer

import (
	"github.com/nats-io/nats.go"
	"log"
	"music-streaming-microservices/email-service/global"
	"time"
)

func Consumer() chan []*nats.Msg {
	sub, err := global.NatsJetStream.PullSubscribe(global.Configs.Nats.Streams[0].Subjects[0], "email-service-consumer", nats.BindStream(global.Configs.Nats.Streams[0].Name))
	if err != nil {
		log.Fatalf("Error %s subscribing to stream: %v", global.Configs.Nats.Streams[0].Name, err)
	}

	msgsCh := make(chan []*nats.Msg, 10)

	go func() {
		defer close(msgsCh)
		log.Printf("Consumer fetching......\n")
		for {
			msgs, err := sub.Fetch(5)
			if err != nil {
				log.Printf("[WARN] Failed to fetch messages from stream %s: %v. Sleeping 30 seconds before retry...\n", global.Configs.Nats.Streams[0].Name, err)
				time.Sleep(30 * time.Second)
				continue
			}
			log.Printf("Received %d messages from stream %s", len(msgs), global.Configs.Nats.Streams[0].Name)
			msgsCh <- msgs
		}
	}()

	return msgsCh

}
