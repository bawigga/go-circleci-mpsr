package circleci_mpsr

import (
	"testing"
)

func TestFeedFields(t *testing.T) {
	feed := Feed{}
	feed.url = "example.com"

	if feed.url != "example.com" {
		t.Fatalf("expected: %v, got: %v", "example.com", feed.url)
	}
}
