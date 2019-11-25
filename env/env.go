package env

import (
	"log"
	"os"
)

// ReadEnv reads a non-empty environment variable or exits.
func ReadEnv(key string) string {
	value, present := os.LookupEnv(key)
	if !present {
		log.Fatalf("could not find required environment variable %s", key)
	}
	if value == "" {
		log.Fatalf("found empty value for required environment variable %s", key)
	}

	return value
}

// ReadEnv reads a non-empty environment variable or returns the default string.
func ReadEnvOr(key string, orElse string) string {
	value, present := os.LookupEnv(key)
	if !present {
		log.Printf("could not find environemnt variable %s. Using default %s", key, orElse)
		return orElse
	}
	if value == "" {
		log.Printf("found empty value for environemnt variable %s. Using default %s", key, orElse)
		return orElse
	}

	return value
}
