package jobpostings

import (
	"context"
	"fmt"
	"testing"
)

func TestGetLookerJobPostings(t *testing.T) {
	jobPostings, err := GetLookerJobPostings(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	for jobPosting := range jobPostings {
		fmt.Println("title:", jobPosting.Title, "location:", jobPosting.Location)
	}
}