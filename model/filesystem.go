package model

type FilesystemConfig struct {
	Default string                     `json:"default"`
	Clouds  map[string]*FileSystemInfo `json:"clouds"`
}

type FileSystemInfo struct {
	Driver    string `json:"driver"`
	Domain    string `json:"domain"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Bucket    string `json:"bucket"`
	Transport string `json:"transport"`
}
