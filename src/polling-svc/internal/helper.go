package internal

import "os"

func getMongoConnStr() string {
	value, exists := os.LookupEnv("MONGO_URI")
	if !exists {
		value = "mongodb://localhost:27017"
	}
	return value
}
