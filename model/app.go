package model

type AppConfig struct {
	Name     string       `json:"name"`
	Key      string       `json:"key"`
	Timezone string       `json:"timezone"`
	Version  string       `json:"version"`
	Cookie   *CookieInfo  `json:"cookie"`
	Session  *SessionInfo `json:"session"`
	Notify   *NotifyInfo  `json:"notify"`
	Token    *TokenInfo   `json:"token"`
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
	Issuer     string `json:"issuer"`
	Audience   string `json:"audience"`
	PrivateKey []byte `json:"private_key"`
	ExpireTime int    `json:"expire_time"`
}
