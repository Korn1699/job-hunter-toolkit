package jobpostings

import (
	"context"
	"fmt"
	"testing"
)

func TestGetApptioJobPostings(t *testing.T) {
	jobPostings, err := GetApptioJobPostings(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	for jobPosting := range jobPostings {
		fmt.Println("title:", jobPosting.Title, "location:", jobPosting.Location)
	}
}
