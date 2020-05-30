package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	e := godotenv.Load()
	if e != nil {
		log.Println(e)
	}

	router := httprouter.New()
	router.GET("/healtzh", Healthz)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Println("listen at port " + port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		panic(fmt.Sprintf("%s: %s", "Failed to listen and server", err))
	}
}

func Healthz(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode("ok")
}
