package jobpostings

import (
	"context"
	"fmt"
	"testing"
)

func TestGetMediumJobPostings(t *testing.T) {
	jobPostings, err := GetMediumJobPostings(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	for jobPosting := range jobPostings {
		fmt.Println("title:", jobPosting.Title, "location:", jobPosting.Location)
	}
}