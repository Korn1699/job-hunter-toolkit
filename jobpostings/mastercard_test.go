package jobpostings

import (
	"context"
	"fmt"
	"testing"
)

func TestGetMastercardJobPostings(t *testing.T) {
	jobPostings, err := GetMastercardJobPostings(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	for jobPosting := range jobPostings {
		fmt.Println("title:", jobPosting.Title, "location:", jobPosting.Location)
	}
}