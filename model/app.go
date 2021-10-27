package model

type AppConfig struct {
	Name        string           `json:"name"`
	Key         string           `json:"key"`
	Timezone    string           `json:"timezone"`
	Version     string           `json:"version"`
	Cookie      *CookieInfo      `json:"cookie"`
	Session     *SessionInfo     `json:"session"`
	Notify      *NotifyInfo      `json:"notify"`
	Token       *TokenInfo       `json:"token"`
	DbPrefix    string           `json:"db_prefix"`
	DbIncrement map[string]int64 `json:"db_increment"`
}

type CookieInfo struct {
	Domain string `json:"domain"`
	MaxAge int    `json:"max_age"`
	Path   string `json:"path"`
}

type SessionInfo struct {
	Driver      string     `json:"driver"`
	CookieName  string     `json:"cookie_name"`
	MaxLifeTime int64      `json:"max_life_time"`
	Provider    *RedisInfo `json:"provider"`
}

type NotifyInfo struct {
	Driver      string     `json:"driver"`
	CookieName  string     `json:"cookie_name"`
	MaxLifeTime int64      `json:"max_life_time"`
	Provider    *RedisInfo `json:"provider"`
}

type TokenInfo struct {
	AccessTokenExp    int64  `json:"access_token_exp"`
	RefreshTokenExp   int64  `json:"refresh_token_exp"`
	IsGenerateRefresh bool   `json:"is_generate_refresh"`
	RedisAddr         string `json:"redis_addr"`
	RedisDB           int    `json:"redis_db"`
	PrivateKey        string `json:"private_key"`
}
