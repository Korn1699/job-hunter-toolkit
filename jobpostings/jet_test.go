package jobpostings

import (
	"context"
	"fmt"
	"testing"
)

func TestGetJetJobPostings(t *testing.T) {
	jobPostings, err := GetJetJobPostings(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	for jobPosting := range jobPostings {
		fmt.Println("title:", jobPosting.Title, "location:", jobPosting.Location)
	}
}