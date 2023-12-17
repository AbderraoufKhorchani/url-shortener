package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AbderraoufKhorchani/url-shortener/internal/handlers"
	"github.com/AbderraoufKhorchani/url-shortener/web"
)

const webPort = "8080"

// change values to connect to your postgres database
const (
	DBHost     = "localhost"
	DBPort     = "5432"
	DBUser     = "postgres"
	DBPassword = "password"
	DBName     = "urls"
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable timezone=UTC connect_timeout=5",
		DBHost, DBPort, DBUser, DBName, DBPassword)
	conn, err := handlers.ConnectToDB(dsn)
	if err != nil {
		log.Panic("Can't connect to Postgres!")
	}

	err = handlers.New(conn)
	if err != nil {
		log.Panic("Postgres not responding!")
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: web.Routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
