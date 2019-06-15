package jobpostings

import (
	"context"
	"fmt"
	"testing"
)

func TestGetUpworkJobPostings(t *testing.T) {
	jobPostings, err := GetUpworkJobPostings(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	for jobPosting := range jobPostings {
		fmt.Println("title:", jobPosting.Title, "location:", jobPosting.Location)
	}
}