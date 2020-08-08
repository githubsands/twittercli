package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	twitterhttp "github.com/githubsands/twittercli/http"
	"github.com/githubsands/twittercli/util"
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
	return 0
}
