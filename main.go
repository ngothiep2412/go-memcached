package main

import (
	"encoding/json"
	"fmt"
	"log"
	"main/cache"
	"main/database"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	db, err := database.NewMySQLDB()

	if err != nil {
		log.Fatalf("Connection Failed %s", err)
	}

	defer db.Close()

	client, err := cache.NewMemCached()

	if err != nil {
		log.Fatalf("Could not initialize Memcached client %s", err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/users/{mssv}", func(w http.ResponseWriter, r *http.Request) {
		mssv := mux.Vars(r)["mssv"]

		val, err := client.GetUser(mssv)

		if err == nil {
			fmt.Print("test")
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(&val)
			return
		}

		res, err := db.FindBYMSSV(mssv)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(err.Error())
		}

		_ = client.SetUser(res)

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(&res)
	})

	fmt.Println("Starting server :8080")

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
