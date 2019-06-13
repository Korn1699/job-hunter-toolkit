package jobpostings

import (
	"context"
	"fmt"
	"testing"
)

func TestGetMessageBirdJobPostings(t *testing.T) {
	jobPostings, err := GetMessageBirdJobPostings(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	for jobPosting := range jobPostings {
		fmt.Println("title:", jobPosting.Title, "location:", jobPosting.Location)
	}
}