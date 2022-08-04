package config

import "go.mongodb.org/mongo-driver/mongo"

type AppConfig struct {
	Port       string
	Collection *mongo.Collection
}
