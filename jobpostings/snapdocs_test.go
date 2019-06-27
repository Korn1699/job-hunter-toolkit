package jobpostings

import (
	"context"
	"fmt"
	"testing"
)

func TestGetSnapdocsJobPostings(t *testing.T) {
	jobPostings, err := GetSnapdocsJobPostings(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	for jobPosting := range jobPostings {
		fmt.Println("title:", jobPosting.Title, "location:", jobPosting.Location)
	}
}