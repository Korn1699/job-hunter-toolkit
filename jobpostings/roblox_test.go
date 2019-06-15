package jobpostings

import (
	"context"
	"fmt"
	"testing"
)

func TestGetRobloxJobPostings(t *testing.T) {
	jobPostings, err := GetRobloxJobPostings(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	for jobPosting := range jobPostings {
		fmt.Println("title:", jobPosting.Title, "location:", jobPosting.Location)
	}
}
