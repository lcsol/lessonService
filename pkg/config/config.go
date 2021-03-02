package config

import (
	"flag"
	"fmt"
)

type Config struct {
	serverAddr       string
	serverPort       int
	MongoURL         string
	Database         string
	LessonCollection string
	ModelCollection  string
}

func Get() *Config {
	conf := &Config{}

	flag.StringVar(&conf.serverAddr, "serverAddr", "localhost", "server host")
	flag.IntVar(&conf.serverPort, "serverPort", 8080, "server port")
	// flag.StringVar(&conf.MongoURL, "mongoURL", "mongodb://127.0.0.1:27017", "mongodb url")
	flag.StringVar(&conf.MongoURL, "mongoURL", "mongodb+srv://chao:vedUDwrvzqcz4BjD@dev-pl-0.mqwfb.mongodb.net/test", "mongodb url")
	// flag.StringVar(&conf.Database, "database", "Lessons", "datbase")
	flag.StringVar(&conf.Database, "database", "test", "datbase")
	flag.StringVar(&conf.LessonCollection, "lessonCollection", "lesson", "database collecion of lesson")
	flag.StringVar(&conf.ModelCollection, "modelCollection", "model", "database collecion of model")

	flag.Parse()

	return conf
}

func (c *Config) GetServerURL() string {
	return fmt.Sprintf("%s:%d", c.serverAddr, c.serverPort)
}
