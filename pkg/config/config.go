package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	serverAddr       string
	serverPort       string
	mongoUsername    string
	mongoPassword    string
	mongoURI         string
	Database         string
	LessonCollection string
	ModelCollection  string
}

// NewConfig returns a Config
func NewConfig() *Config {
	// load environment variables from .env file
	if err := godotenv.Load(filepath.Join("../../", ".env")); err != nil {
		log.Println("failed to load env vars")
		return nil
	}

	conf := &Config{}

	flag.StringVar(&conf.serverAddr, "serverAddr", os.Getenv("SERVER_HOST"), "server host")
	flag.StringVar(&conf.serverPort, "serverPort", os.Getenv("SERVER_PORT"), "server port")
	// flag.StringVar(&conf.MongoURL, "mongoURL", "mongodb://127.0.0.1:27017", "mongodb url")
	flag.StringVar(&conf.mongoUsername, "mongoUsername", os.Getenv("MONGODB_USERNAME"), "mongodb username")
	flag.StringVar(&conf.mongoPassword, "mongoPassword", os.Getenv("MONGODB_PASSWORD"), "mongodb password")
	flag.StringVar(&conf.mongoURI, "mongoURI", os.Getenv("MONGODB_URI"), "mongodb uri")
	// flag.StringVar(&conf.Database, "database", "Lessons", "datbase")
	flag.StringVar(&conf.Database, "database", os.Getenv("DB_NAME"), "datbase")
	flag.StringVar(&conf.LessonCollection, "lessonCollection", os.Getenv("DB_COLLECTION"), "database collecion of lesson")
	flag.StringVar(&conf.ModelCollection, "modelCollection", "model", "database collecion of model")

	flag.Parse()

	return conf
}

// GetServerURL returns the server url
func (c *Config) GetServerURL() string {
	return fmt.Sprintf("%s:%s", c.serverAddr, c.serverPort)
}

// GetDatabaseURL returns the database url
func (c *Config) GetDatabaseURL() string {
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s", c.mongoUsername, c.mongoPassword, c.mongoURI, c.Database)
	fmt.Println(c.mongoUsername)
	return uri
}
