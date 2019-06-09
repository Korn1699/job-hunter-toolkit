package jobpostings

import (
	"context"
	"fmt"
	"testing"
)

func TestGetCornellUniversityJobPostings(t *testing.T) {
	jobPostings, err := GetCornellUniversityJobPostings(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	for jobPosting := range jobPostings {
		fmt.Println("title:", jobPosting.Title, "location:", jobPosting.Location)
	}
}