package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/davecgh/go-spew/spew"
	twitterhttp "github.com/githubsands/twittercli/http"
)

const (
	defaultConfig = "config.yml"
	defaultLogger = "logs"
)

type config struct {
	authToken       string `yaml:"authtoken"`
	marshalResponse bool   `yaml:"marshalresponse"`
}

func main() {
	var (
		l = startLogger()
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
		cfg, err = getConfig()
	)

	if err != nil {
		fmt.Printf("%v\n", err)
		return 1
	}

	tc, err := twitterhttp.NewTwitterClient(cfg.marshalResponse)
	if err != nil {
		fmt.Printf("%v\n", err)
		return 1
	}

	req, err := tc.NewTwitterRequest()
	if err != nil {
		fmt.Printf("%v\n", err)
		return 1
	}

	res, err := tc.GetHTTPResponseBodyBytes(req)
	if err != nil {
		return 1
	}

	fmt.Printf("%v\n", res)
	return 0
}

func startLogger() *log.Logger {
	var (
		loc, _ = os.Getwd()
	)

	w, err := os.Create(defaultLogger)
	if err != nil {
		panic(err)
	}
	defer w.Close()

	return log.New(w, loc, 3)
}

func getConfig() (*config, error) {
	var (
		configPath = defaultConfig
		err        error
	)

	var cfg *config
	fmt.Printf("%v\n", configPath)
	buf, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(buf, &cfg)
	if err != nil {
		log.Fatalf("Failed to unmarshal: %v, err")
	}

	spew.Dump(cfg)
	return cfg, nil
}
