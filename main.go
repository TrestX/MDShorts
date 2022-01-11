package main

import (
	"MdShorts/pkg/router"
	"log"
	"net/http"

	"github.com/aekam27/trestCommon"
	"github.com/rs/cors"
)

// setupGlobalMiddleware will setup CORS
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.AllowAll().Handler
	return handleCORS(handler)
}

// our main function
func main() {

	trestCommon.LoadConfig()
	router := router.NewRouter()
	log.Fatal(http.ListenAndServe(":6019", setupGlobalMiddleware(router)))
}

////md-shorts-backend.doceree.com
