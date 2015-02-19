package circleci_mpsr

import (
	"testing"
)

func TestFeedFields(t *testing.T) {
	feed := Feed{}
	feed.Url = "example.com"

	if feed.Url != "example.com" {
		t.Fatalf("expected: %v, got: %v", "example.com", feed.Url)
	}
}

func TestFeedParseXml(t *testing.T) {
	xmlData := []byte(`
		<?xml version="1.0" encoding="UTF-8"?>
		<Projects>
			<Project lastBuildTime="2015-02-10T14:14:02.800Z" lastBuildLabel="525" lastBuildStatus="Success" name="EvoSure/api" activity="Sleeping" webUrl="https://circleci.com/gh/EvoSure/api/tree/master"></Project>
			<Project lastBuildTime="2015-02-04T23:26:05.791Z" lastBuildLabel="1027" lastBuildStatus="Success" name="EvoSure/exchange" activity="Sleeping" webUrl="https://circleci.com/gh/EvoSure/exchange/tree/master"></Project>
		</Projects>
	`)

	f := Feed{}
	projects, err := f.parseXml(xmlData)
	if err != nil {
		t.Fatalf("error parsing xml")
	}

	numProjects := len(projects)
	if numProjects != 2 {
		t.Fatalf("expected: 2 got: %v", numProjects)
	}
}
