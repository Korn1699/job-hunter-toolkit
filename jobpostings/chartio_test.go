package jobpostings

import (
	"context"
	"fmt"
	"testing"
)

func TestGetChartioJobPostings(t *testing.T) {
	jobPostings, err := GetChartioJobPostings(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	for jobPosting := range jobPostings {
		fmt.Println("title:", jobPosting.Title, "location:", jobPosting.Location)
	}
}