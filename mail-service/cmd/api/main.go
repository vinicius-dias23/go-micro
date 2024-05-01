package main

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct {
	Mailer Mail
}

const webPort = "80"

func main() {
	app := Config{
		Mailer: createMail(),
	}

	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	log.Println("Starting mail service on port", webPort)

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
