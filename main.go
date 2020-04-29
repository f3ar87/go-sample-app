package main

import (
	"net/http"

	"github.com/f3ar87/go-sample-app/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}

/*
	Different way to run the application
		go run main.go
		go run . --> check if there is a main package in the directory
		go run github.com/f3ar87/go-sample-app --> fully qualified

	Per buildare l'applicazione vale quanto sopra ma posso usare build al posto di run

	Per eseguire ./my-app
*/
