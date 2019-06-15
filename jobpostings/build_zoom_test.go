package jobpostings

import (
	"context"
	"fmt"
	"testing"
)

func TestGetBuildZoomJobPostings(t *testing.T) {
	jobPostings, err := GetBuildZoomJobPostings(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	for jobPosting := range jobPostings {
		fmt.Println("title:", jobPosting.Title, "location:", jobPosting.Location)
	}
}