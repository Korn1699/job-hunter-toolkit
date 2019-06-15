package jobpostings

import (
	"context"
	"fmt"
	"testing"
)

func TestGetSmarkingJobPostings(t *testing.T) {
	jobPostings, err := GetSmarkingJobPostings(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	for jobPosting := range jobPostings {
		fmt.Println("title:", jobPosting.Title, "location:", jobPosting.Location)
	}
}