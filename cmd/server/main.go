package main

import (
	"log"
	"os"

	"github.com/lcsol/lessonService/cmd/helper"
	"github.com/lcsol/lessonService/pkg/config"
	"github.com/lcsol/lessonService/pkg/http/rest/handlers"
	"github.com/lcsol/lessonService/pkg/http/rest/router"
	repo "github.com/lcsol/lessonService/pkg/repository"
	getting "github.com/lcsol/lessonService/pkg/services/getLesson"
)

func main() {
	var (
		conf                                 = config.Get()
		infoLog       *log.Logger            = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
		errLog        *log.Logger            = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
		client                               = helper.Connect(infoLog, errLog, conf.MongoURL)
		lessonRepo    repo.LessonRepository  = repo.NewLessonCollection(client.Database(conf.Database).Collection(conf.LessonCollection))
		getting       getting.Service        = getting.NewService(lessonRepo)
		lessonHandler handlers.LessonHandler = handlers.NewLessonHandler(infoLog, errLog, getting)
		r                                    = router.Routes(lessonHandler)
	)

	helper.Serve(errLog, infoLog, r.router, conf.GetServerURL())
}
