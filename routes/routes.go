// Router for the app
package routes

import (
	"github.com/gorilla/mux"
	"github.com/YashaswiNayak99/gorilla-gorm-test/services"
)


func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/posts", services.GetAllPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", services.GetPost).Methods("GET")
	router.HandleFunc("/posts", services.CreatePost).Methods("POST")
	router.HandleFunc("/posts/{id}", services.UpdatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", services.DeletePost).Methods("DELETE")

	return router
}