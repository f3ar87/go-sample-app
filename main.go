package main

import (
	"net/http"

	"github.com/f3ar87/go-sample-app/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

	/*
		--> In alternativa posso dichiarare direttamente la funzione che gestisca le richieste

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "public"+r.URL.Path)
		})
		http.ListenAndServe(":8000", nil)
	*/

	/*
		--> Oppure posso usare degli handler pre-built di Go
		http.ListenAndServe(":8000", http.FileServer(http.Dir("public")))
	*/

}

/*
	Different way to run the application
		go run main.go
		go run . --> check if there is a main package in the directory
		go run github.com/f3ar87/go-sample-app --> fully qualified

	Per buildare l'applicazione vale quanto sopra ma posso usare build al posto di run

	Per eseguire ./my-app
*/
