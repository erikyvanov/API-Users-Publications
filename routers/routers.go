package routers

import (
	"log"
	"net/http"
	"os"

	"github.com/erikyvanov/API-Users-Posts/handlers"
	"github.com/erikyvanov/API-Users-Posts/middlewares"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Routers configura el puerto y los routers
func Routers() {
	router := mux.NewRouter()

	//Crear rutas
	router.HandleFunc("/user", middlewares.CheckDB(handlers.Register)).Methods("POST")
	router.HandleFunc("/login", middlewares.CheckDB(handlers.Login)).Methods("POST")
	router.HandleFunc("/view-profile", middlewares.CheckDB(middlewares.ValidateJWT(handlers.ViewProfile))).Methods("GET")
	router.HandleFunc("/modify-profile", middlewares.CheckDB(middlewares.ValidateJWT(handlers.ModifyProfile))).Methods("PUT")

	router.HandleFunc("/post", middlewares.CheckDB(middlewares.ValidateJWT(handlers.NewPost))).Methods("POST")
	router.HandleFunc("/read-posts", middlewares.CheckDB(middlewares.ValidateJWT(handlers.ReadPosts))).Methods("GET")
	router.HandleFunc("/post", middlewares.CheckDB(middlewares.ValidateJWT(handlers.DeletePost))).Methods("DELETE")

	//Creamos el PORT
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
