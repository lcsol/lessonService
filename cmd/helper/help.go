package helper

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect starts a mongoDB client
func Connect(errLog *log.Logger, mongoURL string) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))
	if err != nil {
		errLog.Fatal(err)
	}

	return client, err
}

// Serve initializes a new http.Server
func Serve(errLog *log.Logger, infoLog *log.Logger, router *mux.Router, serverURI string) {
	srv := &http.Server{
		Addr:         serverURI,
		ErrorLog:     errLog,
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	infoLog.Printf("Starting server on %s", serverURI)
	errLog.Fatal(srv.ListenAndServe())
}
