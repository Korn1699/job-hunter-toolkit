package jobpostings

import (
	"context"
	"fmt"
	"testing"
)

func TestGetTripAdvisorJobPostings(t *testing.T) {
	jobPostings, err := GetTripAdvisorJobPostings(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	for jobPosting := range jobPostings {
		fmt.Println("title:", jobPosting.Title, "location:", jobPosting.Location)
	}
}