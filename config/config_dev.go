// +build dev

package config

const (
	// DbHost - database host (document)
	DbHost = "10.10.10.60"
	// DbUser - user for mongo
	DbUser = "root"
	// DbPassword - password for mongo
	DbPassword = "pass12345"
	// DbPort - Mongo port
	DbPort = 27017
	// DbName - Database name
	DbName = "stonks"
	// DbCollection - collection of stonk sentiments
	DbCollection = "sentiments"
	// APIPort - port to handle ticker requests
	APIPort = 8080
)
