module main

go 1.15

replace (
	lessonService/handlers => ./handlers
	lessonService/models => ./models
)

require (
	go.mongodb.org/mongo-driver v1.4.6 // indirect
	lessonService/handlers v0.0.0-00010101000000-000000000000 // indirect
)
