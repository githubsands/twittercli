package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	twitterhttp "github.com/githubsands/twittercli/http"
	"github.com/githubsands/twittercli/util"
	"github.com/gorilla/mux"
)

const (
	defaultLogger = "logs"
	defaultConfig = "config.yml"
)

type config struct {
	authToken       string `yaml:"authtoken"`
	marshalResponse bool   `yaml:"marshalresponse"`
}

func main() {
	var (
		l = util.StartLogger(defaultLogger)
	)

	err := twittercli(l)
	if err != 0 {
		log.Fatal(err)
		os.Exit(1)
	}

	os.Exit(0)
}

func twittercli(l *log.Logger) int {
	var (
		cfg, err = util.GetConfig(defaultConfig)
	)

	if err != nil {
		fmt.Printf("%v\n", err)
		return 1
	}

	tc, err := twitterhttp.NewTwitterClient(cfg)
	if err != nil {
		fmt.Printf("%v\n", err)
		return 1
	}

	req := &http.Request{}

	res, err := tc.GetHTTPResponseBodyBytes(req)
	if err != nil {
		return 1
	}

	fmt.Printf("%v\n", res)

	if cfg.Server() {
		CreateServer().ListenAndServe()
	}

	return 0
}

func CreateServer() *http.Server {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}
