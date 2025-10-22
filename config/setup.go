package config

import "os"

// Simple helpers to read configuration from environment with safe defaults.
func GetEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func GetServerPort() string {
	// return value like ":8080"
	p := GetEnv("SERVER_PORT", "8080")
	return ":" + p
}

func GetMQTTBroker() string { return GetEnv("MQTT_BROKER", "") }
func GetMQTTUser() string   { return GetEnv("MQTT_USER", "") }
func GetMQTTPass() string   { return GetEnv("MQTT_PASS", "") }
func GetMQTTTopic() string  { return GetEnv("MQTT_TOPIC", "lowbackpain/sensor/data") }

func GetJWTSecret() string { return GetEnv("JWT_SECRET", "dev-secret") }
