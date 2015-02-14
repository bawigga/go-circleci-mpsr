// Copyright 2015 Brian Wigginton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This program watches CircleCI projects endpoint for changes and reports when
// a build is broken of fixed.
package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	CIRCLECI_API_TOKEN string
	POLL_INTERVAL      int
	RUN_ONCE           bool
)

// Projects is a struct representing a list of projects from the circleci
// response.
type Projects struct {
	ProjectList []Project `xml:"Project"`
}

// Project represents a single projects status
type Project struct {
	Name            string `xml:"name,attr"`
	Activity        string `xml:"activity,attr"`
	LastBuildStatus string `xml:"lastBuildStatus,attr"`
	LastBuildLabel  string `xml:"lastBuildLabel,attr"`
	LastBuildTime   string `xml:"lastBuildTime,attr"`
	NextBuildTime   string `xml:"nextBuildTime,attr"`
	WebUrl          string `xml:"webUrl,attr"`
}

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

// getProjectsXML returns the projects xml from circleci.com with the provided
// api token
func getProjectsXML() ([]byte, error) {
	circleAddress := "https://circleci.com/cc.xml?circle-token=" + CIRCLECI_API_TOKEN
	log.Println("Fetching: ", circleAddress)

	res, err := http.Get(circleAddress)
	// check for errors
	if err != nil {
		log.Fatal(err)
	}

	// close out the body
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}

func updateProjects() (Projects, error) {

	data, err := getProjectsXML()
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	projects := Projects{}

	err = xml.Unmarshal(data, &projects)
	if err != nil {
		log.Panic(err)
	}

	projectCount := len(projects.ProjectList)
	if projectCount == 0 {
		fmt.Println("No projects found")
	} else {
		fmt.Printf("Found %v project(s).\n", projectCount)
	}

	// fmt.Printf(": %v", projects.ProjectList)

	return projects, err
}

func main() {

	flag.Parse()

	for {
		updateProjects()
		time.Sleep(time.Duration(POLL_INTERVAL) * time.Second)
		if RUN_ONCE == true {
			os.Exit(0)
		}
	}
}
