package config

import "os"

type Config struct {
	App  AppConfig
	HTTP HTTPConfig
}

type AppConfig struct {
	Name string
	Env  string
}

type HTTPConfig struct {
	APIAddr     string
	GatewayAddr string
	RAGAddr     string
	SearchAddr  string
	ReportAddr  string
}

func Load() Config {
	return Config{
		App: AppConfig{
			Name: getEnv("APP_NAME", "agent-platform"),
			Env:  getEnv("APP_ENV", "dev"),
		},
		HTTP: HTTPConfig{
			APIAddr:     getEnv("API_ADDR", ":8080"),
			GatewayAddr: getEnv("GATEWAY_ADDR", ":8081"),
			RAGAddr:     getEnv("RAG_ADDR", ":8082"),
			SearchAddr:  getEnv("SEARCH_ADDR", ":8083"),
			ReportAddr:  getEnv("REPORT_ADDR", ":8084"),
		},
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
