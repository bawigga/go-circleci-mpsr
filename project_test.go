package circleci_mpsr

import (
	"encoding/xml"
	"testing"
)

func TestProjectFields(t *testing.T) {
	p := Project{
		Name:            "bawigga/circleci-feed",
		Activity:        "Sleeping",
		LastBuildStatus: "Success",
		LastBuildLabel:  "525",
		LastBuildTime:   "2015-02-10T08:30:00.800Z",
		NextBuildTime:   "2015-02-10T16:00:00.800Z",
		WebUrl:          "https://circleci.com/gh/bawigga/circleci-feed/tree/master",
	}

	if p.Name != "bawigga/circleci-feed" {
		t.Fatalf("expected: %v, got: %v", "bawigga/circleci-feed", p.Name)
	}

	if p.Activity != "Sleeping" {
		t.Fatalf("expected: %v, got: %v", "Sleeping", p.Activity)
	}

	if p.LastBuildStatus != "Success" {
		t.Fatalf("expected: %v, got: %v", "Success", p.LastBuildStatus)
	}

	if p.LastBuildLabel != "525" {
		t.Fatalf("expected: %v, got: %v", "525", p.LastBuildLabel)
	}

	if p.LastBuildTime != "2015-02-10T08:30:00.800Z" {
		t.Fatalf("expected: %v, got: %v", "2015-02-10T08:30:00.800Z", p.LastBuildTime)
	}

	if p.NextBuildTime != "2015-02-10T16:00:00.800Z" {
		t.Fatalf("expected: %v, got: %v", "2015-02-10T08:30:00.800Z", p.NextBuildTime)
	}

	expectedUrl := "https://circleci.com/gh/bawigga/circleci-feed/tree/master"
	if p.WebUrl != expectedUrl {
		t.Fatalf("expected: %v, got: %v", expectedUrl, p.WebUrl)
	}
}

func TestProjectFromXML(t *testing.T) {
	xmlData := `
		<Project lastBuildTime="2015-02-10T08:30:00.800Z"
				 nextBuildTime="2015-02-10T16:00:00.800Z"
				 lastBuildLabel="525"
				 lastBuildStatus="Success"
				 name="bawigga/circleci-feed"
				 activity="Sleeping"
				 webUrl="https://circleci.com/gh/bawigga/circleci-feed/tree/master"></Project>
	`
	project := Project{}
	err := xml.Unmarshal([]byte(xmlData), &project)
	if err != nil {
		t.Fatalf("%v", err)
	}

	if project.Name != "bawigga/circleci-feed" {
		t.Fatalf("expected: %v, got: %v", "bawigga/circleci-feed", project.Name)
	}

	if project.Activity != "Sleeping" {
		t.Fatalf("expected: %v, got: %v", "Sleeping", project.Activity)
	}

	if project.LastBuildStatus != "Success" {
		t.Fatalf("expected: %v, got: %v", "Success", project.LastBuildStatus)
	}

	if project.LastBuildLabel != "525" {
		t.Fatalf("expected: %v, got: %v", "525", project.LastBuildLabel)
	}

	if project.LastBuildTime != "2015-02-10T08:30:00.800Z" {
		t.Fatalf("expected: %v, got: %v", "2015-02-10T08:30:00.800Z", project.LastBuildTime)
	}

	if project.NextBuildTime != "2015-02-10T16:00:00.800Z" {
		t.Fatalf("expected: %v, got: %v", "2015-02-10T16:00:00.800Z", project.NextBuildTime)
	}

	expectedUrl := "https://circleci.com/gh/bawigga/circleci-feed/tree/master"
	if project.WebUrl != expectedUrl {
		t.Fatalf("expected: %v, got: %v", expectedUrl, project.WebUrl)
	}
}

func TestProjectsFromXML(t *testing.T) {
	xmlData := `
		<?xml version="1.0" encoding="UTF-8"?>
		<Projects>
			<Project lastBuildTime="2015-02-10T14:14:02.800Z" lastBuildLabel="525" lastBuildStatus="Success" name="EvoSure/api" activity="Sleeping" webUrl="https://circleci.com/gh/EvoSure/api/tree/master"></Project>
			<Project lastBuildTime="2015-02-04T23:26:05.791Z" lastBuildLabel="1027" lastBuildStatus="Success" name="EvoSure/exchange" activity="Sleeping" webUrl="https://circleci.com/gh/EvoSure/exchange/tree/master"></Project>
		</Projects>
	`
	v := projects{}
	err := xml.Unmarshal([]byte(xmlData), &v)
	if err != nil {
		t.Fatalf("%v", err)
	}

	projects := v.ProjectList
	projectCount := len(projects)
	if projectCount != 2 {
		t.Fatalf("Expected 2 projects but got %v", projectCount)
	}
}
