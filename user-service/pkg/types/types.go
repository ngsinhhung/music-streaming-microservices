package types

type Configs struct {
	Server   ServerConfigs  `json:"server"`
	Postgres PostgresConfig `json:"postgres"`
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
