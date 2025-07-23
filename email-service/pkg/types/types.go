package types

type Configs struct {
	Server ServerConfigs `json:"server"`
	SMTP   SMTPConfig    `json:"smtp"`
	Redis  RedisConfig   `json:"redis"`
	Nats   NatsConfig    `json:"nats"`
}

type ServerConfigs struct {
	Port string `json:"port"`
}

type SMTPConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	From     string `json:"from"`
	Password string `json:"password"`
}

type RedisConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Database int    `json:"database"`
}

type NatsConfig struct {
	Host    string          `json:"host"`
	Port    string          `json:"port"`
	Streams []StreamConfigs `json:"streams"`
}

type StreamConfigs struct {
	Name     string   `json:"name"`
	Subjects []string `json:"subjects"`
}
