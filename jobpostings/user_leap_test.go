package jobpostings

import (
	"context"
	"fmt"
	"testing"
)

func TestGetUserLeapJobPostings(t *testing.T) {
	jobPostings, err := GetUserLeapJobPostings(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	for jobPosting := range jobPostings {
		fmt.Println("title:", jobPosting.Title, "location:", jobPosting.Location)
	}
}