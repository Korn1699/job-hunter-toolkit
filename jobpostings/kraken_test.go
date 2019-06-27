package jobpostings

import (
	"context"
	"fmt"
	"testing"
)

func TestGetKrakenJobPostings(t *testing.T) {
	jobPostings, err := GetKrakenJobPostings(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	for jobPosting := range jobPostings {
		fmt.Println("title:", jobPosting.Title, "location:", jobPosting.Location)
	}
}