package model

type DatabaseConfig struct {
	Connections map[string]*DatabaseInfo `json:"connections"`
	RedisList   map[string]*RedisInfo    `json:"redis"`
}

type DatabaseInfo struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Prefix   string `json:"prefix"`
}

type RedisInfo struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database int    `json:"database"`
}
