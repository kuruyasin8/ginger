package config

import "os"

var (
	Port     = ":80"
	MongoUri = os.Getenv("MONGO_URI")
)
