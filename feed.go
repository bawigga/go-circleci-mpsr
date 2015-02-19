package circleci_mpsr

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Feed represents a MPSR capable feed
type Feed struct {
	Url      string
	Projects []Project
}

// Poll fetches the build feed
func (f *Feed) Poll() ([]Project, error) {
	xmlData, err := f.fetchFeedXML()
	if err != nil {
		log.Fatal(err)
	}
	return f.parseXml(xmlData)
}

// fetchFeedXML fetches the Url for the feed and returns the bytes
func (f *Feed) fetchFeedXML() ([]byte, error) {
	log.Println("Fetching: ", f.Url)
	res, err := http.Get(f.Url)
	if err != nil {
		log.Fatal(err)
	}

	// close out the body
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}

// parseXml returns the projects xml from circleci.com with the provided api token
func (f *Feed) parseXml(xmlData []byte) ([]Project, error) {

	p := projectsXmlRoot{}

	err := xml.Unmarshal(xmlData, &p)
	if err != nil {
		log.Panic(err)
	}

	f.Projects = p.Projects

	projectCount := len(f.Projects)
	if projectCount == 0 {
		fmt.Println("No projects found")
	} else {
		fmt.Printf("Found %v project(s).\n", projectCount)
	}

	// fmt.Printf(": %v", projects.ProjectList)

	return f.Projects, err
}
