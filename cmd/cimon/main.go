package main

import (
	"flag"
	cci "github.com/bawigga/circle-feed"
	"log"
	"os"
	"time"
)

var (
	CIRCLECI_API_TOKEN string
	POLL_INTERVAL      int
	RUN_ONCE           bool
)

func init() {
	flag.IntVar(&POLL_INTERVAL, "poll-interval", 30, "the number of second to wait between checks")
	flag.BoolVar(&RUN_ONCE, "run-once", false, "only run the check once")

}

func setApiKey() {
	CIRCLECI_API_TOKEN = os.Getenv("CIRCLECI_API_TOKEN")
	if CIRCLECI_API_TOKEN == "" {
		log.Fatal("Environment variable CIRCLECI_API_TOKEN not set\n")
	}
}

func main() {

	flag.Parse()
	circle_url := "https://circleci.com/cc.xml?circle-token=" + CIRCLECI_API_TOKEN
	buildFeed := cci.Feed{
		Url: circle_url,
	}

	for {
		buildFeed.Poll()

		if RUN_ONCE == true {
			os.Exit(0)
		}

		time.Sleep(time.Duration(POLL_INTERVAL) * time.Second)
	}
}
