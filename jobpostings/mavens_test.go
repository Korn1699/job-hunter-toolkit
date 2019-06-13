package jobpostings

import (
	"context"
	"fmt"
	"testing"
)

func TestGetMavensJobPostings(t *testing.T) {
	jobPostings, err := GetMavensJobPostings(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	for jobPosting := range jobPostings {
		fmt.Println("title:", jobPosting.Title, "location:", jobPosting.Location)
	}
}
