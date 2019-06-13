package jobpostings

import (
	"context"
	"fmt"
	"testing"
)

func TestGetRollPayJobPostings(t *testing.T) {
	jobPostings, err := GetRollPayJobPostings(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	for jobPosting := range jobPostings {
		fmt.Println("title:", jobPosting.Title, "location:", jobPosting.Location)
	}
}
