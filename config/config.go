package config

type Config struct {
	StoragePath string
	ListenAddr  string
	// 其他配置项
}

func LoadConfig() (*Config, error) {
	// 从文件或环境变量加载配置
}