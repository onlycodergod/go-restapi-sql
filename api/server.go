package api

import (
	"fmt"
	"rest/api/router"
	"rest/migrate"

	"rest/config"
	"log"
	"net/http"
)

func Run() {
	config.Load()
	migrate.Load()
	fmt.Printf("running... at port %d", config.PORT)
	listen(config.PORT)
}

func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router.LoadCORS(r)))
}