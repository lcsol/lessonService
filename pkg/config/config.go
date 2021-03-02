package config

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	serverAddr       string
	serverPort       string
	MongoURL         string
	Database         string
	LessonCollection string
	ModelCollection  string
}

func Get() *Config {
	conf := &Config{}
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s", os.Getenv("MONGODB_USERNAME"), os.Getenv("MONGODB_PASSWORD"), os.Getenv("MONGODB_URI"))

	flag.StringVar(&conf.serverAddr, "serverAddr", os.Getenv("SERVER_HOST"), "server host")
	flag.StringVar(&conf.serverPort, "serverPort", os.Getenv("SERVER_PORT"), "server port")
	// flag.StringVar(&conf.MongoURL, "mongoURL", "mongodb://127.0.0.1:27017", "mongodb url")
	flag.StringVar(&conf.MongoURL, "mongoURL", uri, "mongodb url")
	// flag.StringVar(&conf.Database, "database", "Lessons", "datbase")
	flag.StringVar(&conf.Database, "database", os.Getenv("DB_NAME"), "datbase")
	flag.StringVar(&conf.LessonCollection, "lessonCollection", os.Getenv("DB_COLLECTION"), "database collecion of lesson")
	flag.StringVar(&conf.ModelCollection, "modelCollection", "model", "database collecion of model")

	flag.Parse()

	return conf
}

func (c *Config) GetServerURL() string {
	return fmt.Sprintf("%s:%d", c.serverAddr, c.serverPort)
}
