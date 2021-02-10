module handlers

go 1.15

replace lessonService/models => ../models

require (
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/gorilla/mux v1.8.0
	github.com/leodido/go-urn v1.2.1 // indirect
	go.mongodb.org/mongo-driver v1.4.6 // indirect
	lessonService/models v0.0.0-00010101000000-000000000000
)
