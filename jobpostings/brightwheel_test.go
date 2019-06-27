package jobpostings

import (
	"context"
	"fmt"
	"testing"
)

func TestGetBrightwheelJobPostings(t *testing.T) {
	jobPostings, err := GetBrightwheelJobPostings(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	for jobPosting := range jobPostings {
		fmt.Println("title:", jobPosting.Title, "location:", jobPosting.Location)
	}
}