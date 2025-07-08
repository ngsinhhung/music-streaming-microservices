package types

type Configs struct {
	Server ServerConfigs `json:"server"`
}

type ServerConfigs struct {
	Port string `json:"port"`
	Mode string `json:"mode"`
}
