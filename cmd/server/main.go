package main

import (
	"github.com/buurzx/d-log/internal/server"
	"log"
)

func main()  {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}


