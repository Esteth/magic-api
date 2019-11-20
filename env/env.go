package env

import (
	"log"
	"os"
)

// ReadEnv reads a non-empty environment variable or exits.
func ReadEnv(key string) string {
	value, present := os.LookupEnv(key)
	if !present {
		log.Fatalf("could not find environment variable %s", key)
	}
	if value == "" {
		log.Fatalf("found empty value for environment variable %s", key)
	}

	return value
}
