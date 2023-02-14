package main

import (
	"context"
	"encoding/json"
	"net/http"
	"log"
	"os"
	"os/signal"
	"time"
	"github.com/gorilla/mux"
)

type Time struct {
	CurrentTime string `json:"current_time"`
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	result := map[string]string{"status": "OK"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func getTime(w http.ResponseWriter, r *http.Request) {
	currentTime := []Time{
		{ CurrentTime: http.TimeFormat },
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(currentTime)
}


func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/health", healthCheck)
	router.HandleFunc("/getTime", getTime)

	//log.Fatal(http.ListenAndServe(":2000", router))
	srv := &http.Server{
		Handler:      router,
		Addr:         ":443",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		log.Println("http service running on", srv.Addr)
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}

	}()
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Println("received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(tc)
}
