package circleci_mpsr

import (
	"encoding/xml"
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

	// fixme: code smell
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

// parseXml parses the XML feed and returns a slice of projects
func (f *Feed) parseXml(xmlData []byte) ([]Project, error) {

	var response struct {
		XMLName  xml.Name  `xml:"Projects"`
		Projects []Project `xml:"Project"`
	}

	err := xml.Unmarshal(xmlData, &response)
	if err != nil {
		log.Panic(err)
	}

	return response.Projects, err
}
