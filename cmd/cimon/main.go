package main

import (
	"flag"
	"fmt"
	cci "github.com/bawigga/circle-feed"
	"os"
	"time"
)

var (
	CIRCLECI_API_TOKEN string
	POLL_INTERVAL      int
	RUN_ONCE           bool
)

func init() {
	// Set the CIRCLECI_API_TOKEN
	CIRCLECI_API_TOKEN = os.Getenv("CIRCLECI_API_TOKEN")
	if CIRCLECI_API_TOKEN == "" {
		fmt.Printf("Environment variable CIRCLECI_API_TOKEN not set\n")
		os.Exit(1)
	}

	flag.IntVar(&POLL_INTERVAL, "poll-interval", 30, "the number of second to wait between checks")
	flag.BoolVar(&RUN_ONCE, "run-once", false, "only run the check once")
}

func main() {

	flag.Parse()
	circle_url := "https://circleci.com/cc.xml?circle-token=" + CIRCLECI_API_TOKEN
	buildFeed := cci.Feed{
		url: circle_url,
	}

	for {
		go buildFeed.poll()

		if RUN_ONCE == true {
			os.Exit(0)
		}

		time.Sleep(time.Duration(POLL_INTERVAL) * time.Second)
	}
}
