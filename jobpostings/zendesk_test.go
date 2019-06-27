package jobpostings

import (
	"context"
	"fmt"
	"testing"
)

func TestGetZendeskJobPostings(t *testing.T) {
	jobPostings, err := GetZendeskJobPostings(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	for jobPosting := range jobPostings {
		fmt.Println("title:", jobPosting.Title, "location:", jobPosting.Location)
	}
}