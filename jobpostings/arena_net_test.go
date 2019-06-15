package jobpostings

import (
	"context"
	"fmt"
	"testing"
)

func TestGetArenaNetJobPostings(t *testing.T) {
	jobPostings, err := GetArenaNetJobPostings(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	for jobPosting := range jobPostings {
		fmt.Println("title:", jobPosting.Title, "location:", jobPosting.Location)
	}
}
