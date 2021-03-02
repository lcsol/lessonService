package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/inspiritvr-organization/lesson-service-draft/cmd/helper"
	"github.com/inspiritvr-organization/lesson-service-draft/pkg/config"
	"github.com/inspiritvr-organization/lesson-service-draft/pkg/http/rest/handlers"
	"github.com/inspiritvr-organization/lesson-service-draft/pkg/http/rest/router"
	repo "github.com/inspiritvr-organization/lesson-service-draft/pkg/repository"
	adding "github.com/inspiritvr-organization/lesson-service-draft/pkg/services/addLesson"
	getting "github.com/inspiritvr-organization/lesson-service-draft/pkg/services/getLesson"
)

func main() {
	var (
		conf                                 = config.Get()
		infoLog       *log.Logger            = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
		errLog        *log.Logger            = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
		client, err                          = helper.Connect(errLog, conf.MongoURL)
		lessonRepo    repo.LessonRepository  = repo.NewLessonCollection(client.Database(conf.Database).Collection(conf.LessonCollection))
		getting       getting.Service        = getting.NewService(lessonRepo)
		adding        adding.Service         = adding.NewService(lessonRepo)
		lessonHandler handlers.LessonHandler = handlers.NewLessonHandler(infoLog, errLog, getting, adding)
		router                               = router.Routes(lessonHandler)
	)
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

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	helper.Serve(errLog, infoLog, router, conf.GetServerURL())
}
