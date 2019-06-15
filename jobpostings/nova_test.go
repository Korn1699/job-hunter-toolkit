package jobpostings

import (
	"context"
	"fmt"
	"testing"
)

func TestGetNovaJobPostings(t *testing.T) {
	jobPostings, err := GetNovaJobPostings(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	for jobPosting := range jobPostings {
		fmt.Println("title:", jobPosting.Title, "location:", jobPosting.Location)
	}
}