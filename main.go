package main

import (
	"log"
	"net/http"
	"github.com/YashaswiNayak99/gorilla-gorm-test/utility"
	"github.com/YashaswiNayak99/gorilla-gorm-test/routes"
	"github.com/YashaswiNayak99/gorilla-gorm-test/services"
)

func main() {
	
	var db = utility.GetConnection()
	services.SetDB(db)
	var appRouter = routes.CreateRouter()
	
	log.Println("Listening on Port 8000")
	log.Fatal(http.ListenAndServe(":8000", appRouter))

}
