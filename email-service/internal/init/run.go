package init

import "music-streaming-microservices/email-service/pkg/wire"

func Run() {
	// Load configuration
	ConfigLoader()

	// Initialize Redis connection
	InitRedis()

	// Initialize Nats Stream connection
	InitNats()

	// Email Handler
	handle := wire.InitEmailHandler()
	handle.EmailHandler()

}
