package config

import "os"

// KeyConfig :
type KeyConfig struct {
	DBHost      string
	SecretOrKey string
}

func loadConfig() KeyConfig {
	dbHost := os.Getenv("DBHOST")
	secretKey := os.Getenv("SECRETORKEY")

	if dbHost == "" {
		dbHost = "mongodb://localhost:27017"
	}

	if secretKey == "" {
		secretKey = "secretOrKey"
	}

	keyServer := KeyConfig{
		dbHost,
		secretKey,
	}

	return keyServer
}

// Key :
var Key KeyConfig = loadConfig()
