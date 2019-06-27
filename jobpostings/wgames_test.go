package jobpostings

import (
	"context"
	"fmt"
	"testing"
)

func TestGetWGamesJobPostings(t *testing.T) {
	jobPostings, err := GetWGamesJobPostings(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	for jobPosting := range jobPostings {
		fmt.Println("title:", jobPosting.Title, "location:", jobPosting.Location)
	}
}