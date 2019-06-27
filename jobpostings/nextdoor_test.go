package jobpostings

import (
	"context"
	"fmt"
	"testing"
)

func TestGetNextdoorJobPostings(t *testing.T) {
	jobPostings, err := GetNextdoorJobPostings(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	for jobPosting := range jobPostings {
		fmt.Println("title:", jobPosting.Title, "location:", jobPosting.Location)
	}
}