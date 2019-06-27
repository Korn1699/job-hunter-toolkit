package jobpostings

import (
	"context"
	"fmt"
	"testing"
)

func TestGetGizmodoJobPostings(t *testing.T) {
	jobPostings, err := GetGizmodoJobPostings(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	for jobPosting := range jobPostings {
		fmt.Println("title:", jobPosting.Title, "location:", jobPosting.Location)
	}
}