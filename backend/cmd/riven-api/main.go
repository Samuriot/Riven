package main

import (
	"log"
	"os"
)

var default_url string = "ws://localhost:7880"
var default_port string = "8080"

type Config struct {
	APIKey string
	APISecret string
	LiveKitURL string
	Port string
}

func CheckEnvVar(name, defaultVal string) string {
	if val := os.Getenv(name); val != "" {
		return val
	}
	return defaultVal
}

func RequireEnvVar(name string) string {
	val := os.Getenv(name) 
	if val == "" {
		log.Fatalf("missing required env var %s", name)
	}
	return val
}

func LoadEnv() Config {
	return Config {
		APIKey: RequireEnvVar("LIVEKIT_API_KEY"),
		APISecret: RequireEnvVar("LIVEKIT_API_SECRET"),
		LiveKitURL: CheckEnvVar("LIVEKIT_URL", default_url),
		Port: CheckEnvVar("PORT", default_port),
	}
}

func main() {
	cfg := LoadEnv()
	// TODO: set logs to a logfile
	log.Printf("riven-api on :%s, LiveKit at %s", cfg.Port, cfg.LiveKitURL)
}
