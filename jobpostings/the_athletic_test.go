package jobpostings

import (
	"context"
	"fmt"
	"testing"
)

func TestGetTheAthleticJobPostings(t *testing.T) {
	jobPostings, err := GetTheAthleticJobPostings(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	for jobPosting := range jobPostings {
		fmt.Println("title:", jobPosting.Title, "location:", jobPosting.Location)
	}
}