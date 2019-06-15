package jobpostings

import (
	"context"
	"fmt"
	"testing"
)

func TestGetSourceressJobPostings(t *testing.T) {
	jobPostings, err := GetSourceressJobPostings(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	for jobPosting := range jobPostings {
		fmt.Println("title:", jobPosting.Title, "location:", jobPosting.Location)
	}
}