package jobpostings

import (
	"context"
	"fmt"
	"testing"
)

func TestGetHashiCorpJobPostings(t *testing.T) {
	jobPostings, err := GetHashiCorpJobPostings(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	for jobPosting := range jobPostings {
		fmt.Println("title:", jobPosting.Title, "location:", jobPosting.Location)
	}
}