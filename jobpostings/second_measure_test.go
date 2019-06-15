package jobpostings

import (
	"context"
	"fmt"
	"testing"
)

func TestGetSecondMeasureJobPostings(t *testing.T) {
	jobPostings, err := GetSecondMeasureJobPostings(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	for jobPosting := range jobPostings {
		fmt.Println("title:", jobPosting.Title, "location:", jobPosting.Location)
	}
}
