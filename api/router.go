package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type Router struct {
	Port string
}

func (routerInfo Router) Init() error {
	//new router from chi
	router := chi.NewRouter()

	//http server
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + routerInfo.Port,
	}

	log.Printf("Starting port on %v", routerInfo.Port)

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

