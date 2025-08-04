package types

type Configs struct {
	Server   ServerConfigs  `json:"server"`
	Postgres PostgresConfig `json:"postgres"`
	Redis    RedisConfig    `json:"redis"`
	Nats     NatsConfig     `json:"nats"`
}

type RedisConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Database int    `json:"database"`
}

type ServerConfigs struct {
	Port string `json:"port"`
	Mode string `json:"mode"`
}

type PostgresConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
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

type PayloadAccessToken struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
}

type PayloadRefreshToken struct {
	UserID int64  `json:"user_id"`
	JTI    string `json:"jti"`
}
