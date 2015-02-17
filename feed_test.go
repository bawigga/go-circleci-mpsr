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
