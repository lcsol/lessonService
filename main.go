package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"lessonService/handlers"
	"lessonService/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))
	if err != nil {
		errLog.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		errLog.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	infoLog.Printf("Database connection established")

	// Initialize a new instance of lessonHandler
	lessons := models.NewLessonCollection(client.Database(database).Collection(lessonCollection))
	models := models.NewModelCollection(client.Database(database).Collection(modelCollection))
	lessonHandler := handlers.NewLessonHandler(infoLog, errLog, lessons, models)

	// Initialize a new http.Server
	serverURI := fmt.Sprintf("%s:%d", serverAddr, serverPort)

	srv := &http.Server{
		Addr:         serverURI,
		ErrorLog:     errLog,
		Handler:      lessonHandler.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on %s", serverURI)
	err = srv.ListenAndServe()
	errLog.Fatal(err)
}
