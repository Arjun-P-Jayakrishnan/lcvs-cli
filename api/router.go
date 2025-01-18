package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type dbRouter struct {
	port string
}

func initServer(routerInfo dbRouter) error {
	//new router from chi
	router := chi.NewRouter()

	//http server
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + routerInfo.port,
	}

	log.Printf("Starting port on %v", routerInfo.port)

	//Listening on the port for inputs
	srvErr := srv.ListenAndServe()

	router.Use(
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"http://*", "https://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"*"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		}))



	return srvErr
}
